package handler

import (
	"GO_PROJECT/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEvenAvg(t *testing.T) {

	data := []float64{5, 10, 15, 20}
	body := model.Operation{Op: "avg", Nums: data}
	reqbody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/even", strings.NewReader(string(reqbody)))
	w := httptest.NewRecorder()

	EvenAvg(w, req)

	response := w.Result()
	respBody, _ := io.ReadAll(response.Body)

	if http.StatusOK != response.StatusCode {
		t.Fatalf("expected statusCode: %v, got: %v", http.StatusOK, response.StatusCode)
	}

	expected := "15"

	if expected != string(respBody) {
		t.Errorf("expected: %v, got: %v", expected, string(reqbody))
	}

}

func TestEvenAvgInvalidMethod(t *testing.T) {
	data := []float64{5, 10, 15, 20}
	body := model.Operation{Op: "avg", Nums: data}
	reqbody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodGet, "/even", strings.NewReader(string(reqbody)))
	w := httptest.NewRecorder()

	EvenAvg(w, req)

	response := w.Result()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %v, got %v", http.StatusMethodNotAllowed, response.StatusCode)
	}
}

func TestEvenAvgEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/even", nil)
	w := httptest.NewRecorder()

	EvenAvg(w, req)

	resp := w.Result()
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %v, got %v", resp.StatusCode, resp.StatusCode)
	}

	err_msg := map[string]string{"error": "failed to read req body"}
	response, _ := json.Marshal(err_msg)

	if string(respBody) != string(response) {
		t.Errorf("expected error message %v; got %v", err_msg, string(respBody))
	}
}

func TestOddAvg(t *testing.T) {
	data := []float64{5, 10, 15, 20}
	body := model.Operation{Op: "avg", Nums: data}
	reqbody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/odd", strings.NewReader(string(reqbody)))
	w := httptest.NewRecorder()

	OddAvg(w, req)

	response := w.Result()
	respBody, _ := io.ReadAll(response.Body)

	if http.StatusOK != response.StatusCode {
		t.Fatalf("expected statusCode: %v, got: %v", http.StatusOK, response.StatusCode)
	}

	expected := "10"

	if expected != string(respBody) {
		t.Errorf("expected: %v, got: %v", expected, string(reqbody))
	}

}

func TestOddAvgInvalidMethod(t *testing.T) {
	data := []float64{5, 10, 15, 20}
	body := model.Operation{Op: "avg", Nums: data}
	reqbody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodGet, "/odd", strings.NewReader(string(reqbody)))
	w := httptest.NewRecorder()

	OddAvg(w, req)

	response := w.Result()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %v, got %v", http.StatusMethodNotAllowed, response.StatusCode)
	}
}

func TestOddAvgEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/odd", nil)
	w := httptest.NewRecorder()

	OddAvg(w, req)

	resp := w.Result()
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %v, got %v", resp.StatusCode, resp.StatusCode)
	}

	err_msg := map[string]string{"error": "failed to read req body"}
	response, _ := json.Marshal(err_msg)

	if string(respBody) != string(response) {
		t.Errorf("expected error message %v; got %v", err_msg, string(respBody))
	}
}

func TestEvenoddAvg(t *testing.T) {
	data := []float64{5, 10, 15, 20}
	body := model.Operation{Op: "avg", Nums: data}
	reqbody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/evenodd", strings.NewReader(string(reqbody)))
	w := httptest.NewRecorder()

	EvenoddAvg(w, req)

	response := w.Result()
	respBody, _ := io.ReadAll(response.Body)

	if http.StatusOK != response.StatusCode {
		t.Fatalf("expected statusCode: %v, got: %v", http.StatusOK, response.StatusCode)
	}

	expected := "12.5"

	if expected != string(respBody) {
		t.Errorf("expected: %v, got: %v", expected, string(reqbody))
	}

}

func TestEvenoddAvgInvalidMethod(t *testing.T) {
	data := []float64{5, 10, 15, 20}
	body := model.Operation{Op: "avg", Nums: data}
	reqbody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodGet, "/evenodd", strings.NewReader(string(reqbody)))
	w := httptest.NewRecorder()

	EvenoddAvg(w, req)

	response := w.Result()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %v, got %v", http.StatusMethodNotAllowed, response.StatusCode)
	}
}

func TestEvenoddAvgEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/evenodd", nil)
	w := httptest.NewRecorder()

	EvenoddAvg(w, req)

	resp := w.Result()
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %v, got %v", resp.StatusCode, resp.StatusCode)
	}

	err_msg := map[string]string{"error": "failed to read req body"}
	response, _ := json.Marshal(err_msg)

	if string(respBody) != string(response) {
		t.Errorf("expected error message %v; got %v", err_msg, string(respBody))
	}
}