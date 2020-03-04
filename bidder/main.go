package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/bidder/constants"
	"github.com/NoSkillGirl/greedy_bidders/bidder/routers"
)

func main() {
	constants.SetConstants()
	config := constants.C.GetConf()
	routers.BidderRoutes()
	informAuctioneerAboutYou(config.AuctioneerRegisterURL)
	http.ListenAndServe(":"+config.Port, nil)
}

func informAuctioneerAboutYou(auctioneerURL string) {
	type RegisterBidderRequest struct {
		BidderId string `json:"bidder_id"`
		Host     string `json:"host"`
	}

	bidderID := constants.GetBidderId()
	host := "http://localhost:" + constants.C.Port

	thisRequest := RegisterBidderRequest{
		BidderId: bidderID,
		Host:     host,
	}

	thisRequestInBytes, err := json.Marshal(thisRequest)

	req, err := http.NewRequest("POST", auctioneerURL, bytes.NewBuffer(thisRequestInBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading response body from smsService", err)
	}

	fmt.Println(string(body))
}
