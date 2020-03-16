// package controllers

// import (
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
// )

// func init() {
// 	os.Setenv("DB_HOST", "localhost")
// 	os.Setenv("DB_PORT", "3306")
// 	os.Setenv("DB_USER", "root")
// 	os.Setenv("DB_NAME", "greedy_bidder")
// }

// type TestStruct struct {
// 	requestBody        string
// 	expectedStatusCode int
// 	responseBody       string
// 	observedStatusCode int
// }

// func TestNewAuction(t *testing.T) {
// 	constants.SetConstants()

// 	tests := []TestStruct{
// 		{`{}`, http.StatusBadRequest, "", 0},
// 		{`{"auction_id":""}`, http.StatusBadRequest, "", 0},
// 		{`{"auction_id":"auct2"}`, http.StatusOK, "", 0},
// 	}

// 	for i, testCase := range tests {
// 		var reader io.Reader
// 		reader = strings.NewReader(testCase.requestBody) //Convert string to reader
// 		request, err := http.NewRequest("POST", "http://localhost:8080/NewAuction", reader)
// 		res, err := http.DefaultClient.Do(request)

// 		if err != nil {
// 			t.Error(err) //Something is wrong while sending request
// 			return
// 		}
// 		body, _ := ioutil.ReadAll(res.Body)

// 		tests[i].responseBody = strings.TrimSpace(string(body))
// 		tests[i].observedStatusCode = res.StatusCode
// 	}

// 	DisplayTestCaseResults("TestNewAuction", tests, t)
// }

// func DisplayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {
// 	for _, test := range tests {
// 		if test.observedStatusCode == test.expectedStatusCode {
// 			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
// 		} else {
// 			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
// 		}
// 	}
// }
