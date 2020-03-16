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

	"github.com/NoSkillGirl/greedy_bidders/bidder/constants"
	"github.com/google/uuid"
)

type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
}

func init() {
	os.Setenv("AUCTIONEER_URL", "localhost:8080")
	os.Setenv("PORT", "8081")
	os.Setenv("DELAY", "150")
	os.Setenv("HOST", "localhost")
}

func TestPlaceBid(t *testing.T) {
	constants.SetConstants()

	tests := []TestStruct{
		{`{}`, http.StatusBadRequest, "", 0},
		{`{"auction_id":""}`, http.StatusBadRequest, "", 0},
		{`{"auction_id":"1"}`, http.StatusOK, "", 0},
		{`{"auction_id":"2"}`, http.StatusOK, "", 0},
	}

	for i, testCase := range tests {
		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			PlaceBid(w, r)
		}))
		defer ts.Close()

		// "http://"+constants.Config.Host+":"+constants.Config.Port+"/"
		request, err := http.NewRequest("POST", ts.URL, reader)
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
	displayTestCaseResults("TestPlaceBid", tests, t)
}

func displayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {
	for _, test := range tests {
		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}

// Benchmarking Tests

func BenchmarkTestPlaceBid(t *testing.B) {
	constants.SetConstants()
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
				PlaceBid(w, r)
			}))
			defer ts.Close()

			request, err := http.NewRequest("POST", ts.URL, reader)
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
