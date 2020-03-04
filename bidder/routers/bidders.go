package routers

import (
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/bidder/controllers"
)

func BidderRoutes() {
	http.HandleFunc("/", controllers.PlaceBid)
}
