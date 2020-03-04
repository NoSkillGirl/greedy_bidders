package main

import (
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
	"github.com/NoSkillGirl/greedy_bidders/auctioneer/routers"
)

func init() {
	// Set Constants
	constants.SetConstants()
	// Set Routes
	routers.BidderRoutes()
	routers.AuctioneersRoutes()
}

func main() {
	// Start http server
	http.ListenAndServe(":8080", nil)
}
