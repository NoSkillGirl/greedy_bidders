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

	"github.com/google/uuid"
)

type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
}

func init() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_NAME", "greedy_bidder")
}

func TestNewBiddingRound(t *testing.T) {

	tests := []TestStruct{
		{`{}`, http.StatusBadRequest, "", 0},
		{`{"auction_id":""}`, http.StatusBadRequest, "", 0},
		{`{"auction_id":"4783249"}`, http.StatusOK, "", 0},
	}

	for i, testCase := range tests {
		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader
		request, err := http.NewRequest("GET", "http://localhost:8080/NewAuction", reader)
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
	DisplayTestCaseResults("TestNewBiddingRound", tests, t)
}

func DisplayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {
	for _, test := range tests {
		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}

func BenchmarkTestNewBiddingRound(t *testing.B) {
	t.StartTimer()
	for j := 0; j < t.N; j++ {
		u, _ := uuid.NewUUID()
		reqBody := fmt.Sprintf(`{"auction_id":"%v"}`, u)
		tests := []TestStruct{
			{reqBody, http.StatusOK, "", 0},
		}
		for _, testCase := range tests {
			var reader io.Reader
			reader = strings.NewReader(testCase.requestBody) //Convert string to reader

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				NewBiddingRound(w, r)
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
