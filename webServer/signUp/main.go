package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vdshakh/intellias-go-courses/signUp/common/config"
	db "github.com/vdshakh/intellias-go-courses/signUp/db/sqlc"
)

var newQueries *db.Queries

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic(err.Error())
	}

	conf := config.NewConfigFromEnv()

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		fmt.Printf("db connection failed: %v ", err)
		os.Exit(1)
	}

	newQueries = db.New(conn)

	http.HandleFunc("/users", SignUpHandler)

	err = http.ListenAndServe(conf.Port, nil)
	if err != nil {
		fmt.Printf("error starting server: %s", err)
		os.Exit(1)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)

	case http.MethodGet:
		GetUser(w, r)

	case http.MethodPut:
		UpdateUser(w, r)

	case http.MethodDelete:
		DeleteUser(w, r)

	default:
		WriteResponse(w, http.StatusMethodNotAllowed, "wrong http method")
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"status": http.StatusText(statusCode),
	}
	if data != nil {
		response["data"] = data
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("json encode failed: %v", err)
	}
}

func ReadData(w http.ResponseWriter, r *http.Request) (UserData, error) {
	var user UserData

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("ioutil.ReadAll failed: %v", err))
	}

	if err = json.Unmarshal(bodyBytes, &user); err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("json.Unmarshal failed: %v", err))

		return user, fmt.Errorf("json.Unmarshal failed: %w", err)
	}

	return user, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req InputData

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("ioutil.ReadAll failed: %v", err))
	}

	if err = json.Unmarshal(bodyBytes, &req); err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("json.Unmarshal failed: %v", err))

		return
	}

	err = req.Validate()
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("validation failed: %v", err))

		return
	}

	_, err = newQueries.GetUserByEmail(r.Context(), req.Email)
	if err == nil {
		WriteResponse(w, http.StatusBadRequest, "user exists")

		return
	}

	user := db.CreateUserParams{
		ID:       uuid.New().String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	err = newQueries.CreateUser(r.Context(), user)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("createUser failed: %v", err))

		return
	}

	WriteResponse(w, http.StatusAccepted, "created user successfully")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := ReadData(w, r)
	if err != nil {
		return
	}

	account, err := newQueries.GetUser(r.Context(), user.ID)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("getUser failed: %v", err))
	}

	WriteResponse(w, http.StatusOK, account)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := ReadData(w, r)
	if err != nil {
		return
	}

	err = validatePassword(user.Password)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("validation failed: %v", err))

		return
	}

	updUser := db.UpdatePasswordParams{
		ID:       user.ID,
		Password: user.Password,
	}

	_, err = newQueries.UpdatePassword(r.Context(), updUser)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("updateUser failed: %v", err))
	}

	WriteResponse(w, http.StatusAccepted, "updated user successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := ReadData(w, r)
	if err != nil {
		return
	}

	err = newQueries.DeleteUser(r.Context(), user.ID)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf("deleteUser failed: %v", err))
	}

	WriteResponse(w, http.StatusAccepted, "deleted user successfully")
}
