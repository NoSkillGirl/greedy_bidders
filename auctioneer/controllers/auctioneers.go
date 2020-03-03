package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NoSkillGirl/auctioneer/models"
)

type struct NewBiddingRoundRequest {
	AuctionID string
}

type struct NewBiddingRoundResponse {
	BidderID string
	// TODO: check if float64 is needed?
	Price float64
}

// NewBiddingRound Request
func NewBiddingRound(w http.ResponseWriter, r *http.Request) {

	// Req Obj
	var reqJSON NewBiddingRoundRequest

	// Res Obj
	resp := NewBiddingRoundResponse{}
	w.Header().Set("Content-Type", "application/json")

	// Req Decode
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		// TODO: use logger
		fmt.Println(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// NewBiddingRound model call
	err := models.NewBiddingRound(reqJSON.AuctionID)
	if err != nil {

	}



	// if errOccured == true {
	// 	resp.Status = 500
	// 	resp.Response = ResponseMsg{}
	// 	resp.Error = ErrorMessage{
	// 		Msg: "Internal Server Error",
	// 	}
	// } else {
	// 	resp.Status = 200
	// 	resp.Response = ResponseMsg{
	// 		Msg: "booking details succesfully added",
	// 	}
	// 	resp.Error = ErrorMessage{}

	// 	// Should send SMS to the customer with booking details.
	// 	type SMSRequest struct {
	// 		Mobile  string
	// 		Message string
	// 	}

	// 	smsRequest := SMSRequest{
	// 		Mobile:  "+918904621381",
	// 		Message: "Your Bus booking was successful",
	// 	}

	// 	smsRequestInBytes, err := json.Marshal(smsRequest)

	// 	req, err := http.NewRequest("POST", "http://localhost:8081/SendSMS", bytes.NewBuffer(smsRequestInBytes))
	// 	req.Header.Set("Content-Type", "application/json")

	// 	client := &http.Client{}
	// 	resp, err := client.Do(req)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		fmt.Println("Error in reading response body from smsService", err)
	// 	}

	// 	fmt.Println(string(body))
	// }
	json.NewEncoder(w).Encode(resp)
}

func ListRegisteredBidders(w http.ResponseWriter, r *http.Request){
	

}
