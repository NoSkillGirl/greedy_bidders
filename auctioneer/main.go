package main

import (
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
	"github.com/NoSkillGirl/greedy_bidders/auctioneer/routers"
)

func init() {
	constants.SetConstants()
	routers.BidderRoutes()
	routers.AuctioneersRoutes()
}

func main() {
	http.ListenAndServe(":8080", nil)
}
