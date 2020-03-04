package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/NoSkillGirl/greedy_bidders/bidder/constants"
)

type PlaceBidRequest struct {
	AuctionID string `json:"auction_id"`
}

type PlaceBidResponse struct {
	BidderId string  `json:"bidder_id"`
	Price    float64 `json:"price"`
}

func PlaceBid(w http.ResponseWriter, r *http.Request) {
	// Req Obj
	var reqJSON PlaceBidRequest
	config := constants.C.GetConf()

	// Res Obj
	resp := PlaceBidResponse{}
	w.Header().Set("Content-Type", "application/json")

	// Req Decode
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	price := r1.Float64() * 100

	resp.BidderId = constants.GetBidderId()
	resp.Price = price

	time.Sleep(time.Duration(config.Delay) * time.Millisecond)
	json.NewEncoder(w).Encode(resp)
}
