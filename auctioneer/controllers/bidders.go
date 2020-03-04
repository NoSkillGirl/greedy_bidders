package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/models"
)

type RegisterBidderRequest struct {
	BidderId string `json:"bidder_id"`
	Host     string `json:"host"`
}

type RegisterBidderResponse struct {
}

type GetActiveRegisteredBiddersRequest struct {
}

type GetActiveRegisteredBiddersResponse struct {
	BidderIds []string `json:"bidder_ids"`
}

func RegisterBidder(w http.ResponseWriter, r *http.Request) {
	// Req Obj
	var reqJSON RegisterBidderRequest

	// Res Obj
	resp := RegisterBidderResponse{}
	w.Header().Set("Content-Type", "application/json")

	// Req Decode
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	fmt.Println(r)

	models.RegisterBidder(reqJSON.BidderId, reqJSON.Host)

	json.NewEncoder(w).Encode(resp)
}

func GetActiveRegisteredBidders(w http.ResponseWriter, r *http.Request) {

	// Req Obj
	var reqJSON GetActiveRegisteredBiddersRequest

	// Res Obj
	resp := GetActiveRegisteredBiddersResponse{}
	w.Header().Set("Content-Type", "application/json")

	// Req Decode
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	biddersMap := models.GetActiveRegisteredBidders()

	bidderIds := make([]string, 0)
	for key, _ := range biddersMap {
		bidderIds = append(bidderIds, key)
	}

	resp.BidderIds = bidderIds

	json.NewEncoder(w).Encode(resp)
}
