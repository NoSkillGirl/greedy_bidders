package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/log"
	"github.com/NoSkillGirl/greedy_bidders/auctioneer/models"
)

type NewBiddingRoundRequest struct {
	AuctionID string `json:"auction_id"`
}

type NewBiddingRoundResponse struct {
	BidderID string  `json:"bidder_id"`
	Price    float64 `json:"price"`
}

// NewBiddingRound - requests all bidders and gets responses within 200ms and declares a winner
func NewBiddingRound(w http.ResponseWriter, r *http.Request) {

	// Req Obj
	var reqJSON NewBiddingRoundRequest

	// Res Obj
	resp := NewBiddingRoundResponse{}
	w.Header().Set("Content-Type", "application/json")

	// Req Decode
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	biddersMap := models.GetActiveRegisteredBidders()

	type PlaceBidRequest struct {
		AuctionID string `json:"auction_id"`
	}

	type PlaceBidResponse struct {
		BidderId string  `json:"bidder_id"`
		Price    float64 `json:"price"`
	}

	BiddersMap := make(map[string]float64)

	var wg sync.WaitGroup
	for _, value := range biddersMap {

		go func(v string, wg *sync.WaitGroup) {

			thisRequest := PlaceBidRequest{
				AuctionID: reqJSON.AuctionID,
			}

			thisRequestInBytes, err := json.Marshal(thisRequest)

			req, err := http.NewRequest("POST", v, bytes.NewBuffer(thisRequestInBytes))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{
				Timeout: 200 * time.Millisecond,
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			} else {
				var placeBidResponseJSON PlaceBidResponse

				// Res Decode
				err = json.NewDecoder(resp.Body).Decode(&placeBidResponseJSON)
				if err != nil {
					log.Error.Println(err)
					json.NewEncoder(w).Encode(resp)
					return
				}

				BiddersMap[placeBidResponseJSON.BidderId] = placeBidResponseJSON.Price
			}
			wg.Done()
		}(value, &wg)

		wg.Add(1)

	}

	wg.Wait()

	winnerBidderID := ""
	var winnderBidderPrice float64 = 0

	for k, v := range BiddersMap {
		if v > winnderBidderPrice {
			winnderBidderPrice = v
			winnerBidderID = k
		}
	}

	resp.BidderID = winnerBidderID
	resp.Price = math.Round(winnderBidderPrice*100) / 100

	models.DeclareWinner(reqJSON.AuctionID, resp.BidderID, resp.Price)

	json.NewEncoder(w).Encode(resp)
}
