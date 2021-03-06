package routers

import (
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/controllers"
)

// BidderRoutes - bidder specific routes
func BidderRoutes() {
	http.HandleFunc("/", controllers.HealthCheck)
	http.HandleFunc("/RegisterBidder", controllers.RegisterBidder)
	http.HandleFunc("/GetActiveRegisteredBidders", controllers.GetActiveRegisteredBidders)
}
