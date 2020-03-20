package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/models"
	logger "github.com/NoSkillGirl/greedy_bidders/log"
)

type RegisterBidderRequest struct {
	BidderID string `json:"bidder_id"`
	Host     string `json:"host"`
}

type RegisterBidderResponse struct {
}

type GetActiveRegisteredBiddersRequest struct {
}

type GetActiveRegisteredBiddersResponse struct {
	BidderIds []string `json:"bidder_ids"`
}

// HealthCheck - health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{}")
}

// RegisterBidder - endpoint for registering a bidder
func RegisterBidder(w http.ResponseWriter, r *http.Request) {
	// Req Obj
	var reqJSON RegisterBidderRequest

	// Res Obj
	resp := RegisterBidderResponse{}
	w.Header().Set("Content-Type", "application/json")

	// Req Decode
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		logger.Log.Error(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if reqJSON.BidderID == "" || reqJSON.Host == "" {
		logger.Log.Error("Bidder ID or/and host is not present in the req")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	validUUIDbool := validUUID(reqJSON.BidderID)
	if validUUIDbool == false {
		logger.Log.Error("Bidder ID present in req is not valid")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = models.RegisterBidder(nil, reqJSON.BidderID, reqJSON.Host)
	if err != nil {
		logger.Log.Error(err)
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

// GetActiveRegisteredBidders - endpoint for getting the list of active bidders
func GetActiveRegisteredBidders(w http.ResponseWriter, r *http.Request) {
	// Req Obj
	var reqJSON GetActiveRegisteredBiddersRequest

	// Res Obj
	resp := GetActiveRegisteredBiddersResponse{}
	w.Header().Set("Content-Type", "application/json")

	// Req Decode
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		logger.Log.Error(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	biddersMap, err := models.GetActiveRegisteredBidders(nil)

	if err != nil {
		logger.Log.Error(err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	bidderIds := make([]string, 0)
	for key, _ := range biddersMap {
		bidderIds = append(bidderIds, key)
	}

	resp.BidderIds = bidderIds
	json.NewEncoder(w).Encode(resp)
}

func validUUID(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}
	// r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	// fmt.Println("r: ", r.MatchString(uuid))
	// return r.MatchString(uuid)
	return true
}
