package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

const numberForRandom = 1

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/random", GetRandomNumberHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func GetRandomNumberHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusMethodNotAllowed)
		functionResponse := map[string]map[string]interface{}{"data": {"message": "wrong http method"}}

		err := json.NewEncoder(w).Encode(functionResponse)
		if err != nil {
			panic(err.Error())
		}

		return

	case http.MethodGet:
		minParameter := r.URL.Query().Get("min")
		maxParameter := r.URL.Query().Get("max")

		min, err := strconv.Atoi(minParameter)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Can't convert min parameter to int")))

			return
		}

		max, err := strconv.Atoi(maxParameter)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Can't convert max parameter to int")))

			return
		}

		result := rand.Intn(max-min+numberForRandom) + min

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Your random number between %v and %v: %v", min, max, result)))

		return
	}
}
