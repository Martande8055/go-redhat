package handler

import (
	"GO_PROJECT/average"
	"GO_PROJECT/model"
	"encoding/json"
	"fmt"
	"net/http"
)

var operation model.Operation

func sendResponce(w http.ResponseWriter, statusCode int, payload interface{}) {

	// response, _ := json.Marshal(payload)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	// w.Write(response)
	json.NewEncoder(w).Encode(payload)
}

func sendError(w http.ResponseWriter, StatusCode int, err string) {
	err_msg := map[string]string{"error": err}
	sendResponce(w, StatusCode, err_msg)
}

func handler(w http.ResponseWriter, r *http.Request, str string) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&operation)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "failed to read req body")
		return
	}

	if operation.Op == "" {
		sendError(w, http.StatusInternalServerError, "the 'operation' field is required")
		return
	}

	var obj average.Avg
	switch str {
	case "even":
		obj = &average.Even{Op: operation.Op, Values: operation.Nums}

	case "odd":
		obj = &average.Odd{Op: operation.Op, Values: operation.Nums}

	case "evenodd":
		obj = &average.EvenOdd{Op: operation.Op, Values: operation.Nums}
	}

	avg, err := obj.Calcuate()

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// sendResponce(w, http.StatusOK, avg)
	resultString := fmt.Sprintf("%s of the %s numbers :", operation.Op, str)
	resultmap := map[string]interface{}{"msg": resultString, "result": avg}
	sendResponce(w, http.StatusOK, resultmap)
}

func EvenAvg(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "even")
}

func OddAvg(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "odd")
}

func EvenoddAvg(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "evenodd")
}
