package main

import (
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
	"github.com/NoSkillGirl/greedy_bidders/auctioneer/routers"
)

func main() {
	// routers.AuctioneersRoutes()
	constants.SetConstants()
	routers.BidderRoutes()
	routers.AuctioneersRoutes()
	http.ListenAndServe(":8080", nil)
}
