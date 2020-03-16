package routers

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/NoSkillGirl/greedy_bidders/bidder/controllers"
)

func BidderRoutes() {
	http.HandleFunc("/", controllers.PlaceBid)
}
