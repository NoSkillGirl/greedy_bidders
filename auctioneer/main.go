package main

import (
	"fmt"
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
	fmt.Println("Routers Setup Done")
}

func main() {
	// Start http server on port 8080
	http.ListenAndServe(":8080", nil)
}
