package main

import (
	"net/http"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/routers"
)

func main() {
	routers.AuctioneersRoutes()
	http.ListenAndServe(":8082", nil)
}
