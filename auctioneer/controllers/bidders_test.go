// package controllers

// import (
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
// )

// func init() {
// 	os.Setenv("DB_HOST", "localhost")
// 	os.Setenv("DB_PORT", "3306")
// 	os.Setenv("DB_USER", "root")
// 	os.Setenv("DB_NAME", "greedy_bidder")
// }

// func TestRegisterBidder(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	constants.SetConstants()

// 	tests := []TestStruct{
// 		{`{}`, http.StatusBadRequest, "", 0},
// 		{`{"bidder_id":""}`, http.StatusBadRequest, "", 0},
// 		{`{"host":""}`, http.StatusBadRequest, "", 0},
// 		{`{"bidder_id":"","host":""}`, http.StatusBadRequest, "", 0},
// 		{`{"bidder_id":"wrong-bidder-id","host":"testhost1"}`, http.StatusBadRequest, "", 0},
// 		{`{"bidder_id":"6fbe9d36-649e-11ea-a90c-acde48001132","host":"testhost1"}`, http.StatusOK, "", 0},
// 		{`{"bidder_id":"6fbe9d36-649e-11ea-a90c-acde48001132","host":"testhost1"}`, http.StatusOK, "", 0},
// 	}

// 	for i, testCase := range tests {
// 		var reader io.Reader
// 		reader = strings.NewReader(testCase.requestBody) //Convert string to reader
// 		request, err := http.NewRequest("POST", "http://localhost:8080/RegisterBidder", reader)
// 		res, err := http.DefaultClient.Do(request)

// 		if err != nil {
// 			t.Error(err) //Something is wrong while sending request
// 			return
// 		}
// 		body, _ := ioutil.ReadAll(res.Body)

// 		tests[i].responseBody = strings.TrimSpace(string(body))
// 		tests[i].observedStatusCode = res.StatusCode
// 	}

// 	DisplayTestCaseResults("TestRegisterBidder", tests, t)
// }

// func TestGetActiveRegisteredBidders(t *testing.T) {
// 	tests := []TestStruct{
// 		{`{}`, http.StatusOK, "", 0},
// 	}
// 	for i, testCase := range tests {
// 		var reader io.Reader
// 		reader = strings.NewReader(testCase.requestBody) //Convert string to reader
// 		request, err := http.NewRequest("POST", "http://localhost:8080/GetActiveRegisteredBidders", reader)
// 		res, err := http.DefaultClient.Do(request)

// 		if err != nil {
// 			t.Error(err) //Something is wrong while sending request
// 			return
// 		}
// 		body, _ := ioutil.ReadAll(res.Body)

// 		tests[i].responseBody = strings.TrimSpace(string(body))
// 		tests[i].observedStatusCode = res.StatusCode
// 	}

// 	DisplayTestCaseResults("TestGetActiveRegisteredBidders", tests, t)
// }
