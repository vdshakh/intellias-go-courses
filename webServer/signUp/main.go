package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vdshakh/intellias-go-courses/signUp/common/config"
	queryrepo "github.com/vdshakh/intellias-go-courses/signUp/db/sqlc"
	"io/ioutil"
	"net/http"
	"os"
)

var newQueries *queryrepo.Queries
var conf *config.Config

func main() {
	err := godotenv.Load("config/.env") //check paymentService repo
	if err != nil {
		panic(err.Error())
	}

	conf = config.NewConfigFromEnv()

	http.HandleFunc(conf.Endpoint, SignUpHandler)

	err = http.ListenAndServe(conf.Port, nil)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := sql.Open(conf.DbDriver, conf.DbSource) // to main(), create struct
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf(`{"db connection failed": "%v"}`, err))
	}

	newQueries = queryrepo.New(conn)

	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)

		return

	case http.MethodGet:
		GetUser(w, r)

		return

	case http.MethodPut:
		UpdateUser(w, r)

		return

	case http.MethodDelete:
		DeleteUser(w, r)

		return

	default:
		WriteResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf(`{"message": "wrong http method"}`))

		return
	}
}

func WriteResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req Data

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"json.Unmarshal failed": "%v"}`, err))

		return
	}

	err := req.Validate()
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"validation failed": "%v"}`, err))

		return
	}

	_, err = newQueries.GetUserByEmail(context.Background(), req.Email)
	if err == nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"message": "user exists"}`))

		return
	}

	user := queryrepo.CreateUserParams{
		ID:       uuid.New().String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	err = newQueries.CreateUser(context.Background(), user)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"createUser failed": "%v"}`, err))
	}

	WriteResponse(w, http.StatusAccepted, fmt.Sprintf(`{"message": "created user successfully}`))

	return
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var uID userID

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &uID); err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"json.Unmarshal failed": "%v"}`, err))

		return
	}

	account, err := newQueries.GetUser(context.Background(), uID.ID)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"getUser failed": "%v"}`, uID.ID))
	}

	WriteResponse(w, http.StatusOK, fmt.Sprintf(`{"user": "%v"}`, account))

	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var uID userID

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &uID); err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"json.Unmarshal failed": "%v"}`, err))

		return
	}

	err := validatePassword(uID.Password)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"validatePassword failed": "%v"}`, err))

		return
	}

	updUser := queryrepo.UpdatePasswordParams{
		ID:       uID.ID,
		Password: uID.Password,
	}

	_, err = newQueries.UpdatePassword(context.Background(), updUser)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"updateUser failed": "%v"}`, err))
	}

	WriteResponse(w, http.StatusAccepted, fmt.Sprintf(`{"message": "updated user successfully}`))

	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var uID userID

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bodyBytes, &uID); err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"json.Unmarshal failed": "%v"}`, err))

		return
	}

	err := newQueries.DeleteUser(context.Background(), uID.ID)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"deleteUser failed": "%v"}`, err))
	}

	WriteResponse(w, http.StatusAccepted, fmt.Sprintf(`{"message": "deleted user successfully}`))

	return
}