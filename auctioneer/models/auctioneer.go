package models

import (
	"fmt"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
)

func DeclareWinner(auctionID string, bidderID string, price float64) {
	db := constants.Config.GetDatabaseConnection()

	addAuctionQuery := `insert into auctions (id, winner_bidder_id, price) VALUES ('%s', '%s', %f)`

	addAuctionQueryString := fmt.Sprintf(addAuctionQuery, auctionID, bidderID, price)
	fmt.Println(addAuctionQueryString)

	// perform a db.Query insert
	insert, err := db.Query(addAuctionQueryString)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
