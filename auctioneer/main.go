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

	// contextLogger := log.WithFields(log.Fields{
	// 	"common": "this is a common field",
	// 	"other":  "I also should be logged always",
	// })

	// Start http server on port 8080
	http.ListenAndServe(":8080", nil)
}
