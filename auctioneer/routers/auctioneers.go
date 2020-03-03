package routers

import (
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/controllers"
)

// AuctioneersRoutes - All Auctioneer Related Routes
func AuctioneersRoutes() {
	http.HandleFunc("/NewAuction", controllers.NewBiddingRound)
	http.HandleFunc("/ListRegisteredBidders", controllers.ListRegisteredBidders)
}
