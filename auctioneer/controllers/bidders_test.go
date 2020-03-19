package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func init() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_NAME", "greedy_bidder")
}

func TestRegisterBidder(t *testing.T) {

	tests := []TestStruct{
		{`{}`, http.StatusBadRequest, "", 0},
		{`{"bidder_id":""}`, http.StatusBadRequest, "", 0},
		{`{"host":""}`, http.StatusBadRequest, "", 0},
		{`{"bidder_id":"","host":""}`, http.StatusBadRequest, "", 0},
		{`{"bidder_id":"wrong-bidder-id","host":"testhost1"}`, http.StatusBadRequest, "", 0},
		{`{"bidder_id":"6fbe9d36-649e-11ea-a90c-acde48001132","host":"testhost1"}`, http.StatusOK, "", 0},
		{`{"bidder_id":"6fbe9d36-649e-11ea-a90c-acde48001132","host":"testhost1"}`, http.StatusOK, "", 0},
	}

	for i, testCase := range tests {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader
		request, err := http.NewRequest("POST", "http://localhost:8080/RegisterBidder", reader)
		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
			return
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode
	}
	fmt.Println("=====Result")
	DisplayTestCaseResults("TestRegisterBidder", tests, t)
}

func TestGetActiveRegisteredBidders(t *testing.T) {
	tests := []TestStruct{
		{`{}`, http.StatusOK, "", 0},
	}
	for i, testCase := range tests {
		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader
		request, err := http.NewRequest("GET", "http://localhost:8080/GetActiveRegisteredBidders", reader)
		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
			return
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode
	}
	fmt.Println("=====Result")
	DisplayTestCaseResults("TestGetActiveRegisteredBidders", tests, t)
}

func BenchmarkTestGetActiveRegisteredBidders(t *testing.B) {
	t.StartTimer()
	for j := 0; j < t.N; j++ {
		tests := []TestStruct{
			{"{}", http.StatusOK, "", 0},
		}
		for _, testCase := range tests {
			var reader io.Reader
			reader = strings.NewReader(testCase.requestBody) //Convert string to reader

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				GetActiveRegisteredBidders(w, r)
			}))
			defer ts.Close()

			request, err := http.NewRequest("GET", ts.URL, reader)
			res, err := http.DefaultClient.Do(request)

			if err != nil {
				t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", testCase.requestBody, testCase.expectedStatusCode, testCase.responseBody, testCase.observedStatusCode)
				t.Error(err) //Something is wrong while sending request
				return
			}
			body, _ := ioutil.ReadAll(res.Body)

			testCase.responseBody = strings.TrimSpace(string(body))
			testCase.observedStatusCode = res.StatusCode
			if res.StatusCode != testCase.expectedStatusCode {
				t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", testCase.requestBody, testCase.expectedStatusCode, testCase.responseBody, testCase.observedStatusCode)
			}
		}
	}
}
