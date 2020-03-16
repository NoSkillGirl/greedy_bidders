package models

import (
	"database/sql"
	"fmt"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
	"github.com/NoSkillGirl/greedy_bidders/auctioneer/log"
)

// DeclareWinner - stores in the database about the auction winner and the price
func DeclareWinner(db *sql.DB, auctionID string, bidderID string, price float64) {
	if db == nil {
		db = constants.DbConfig.GetDatabaseConnection()
	}
	addAuctionQuery := `insert into auctions (id, winner_bidder_id, price) VALUES ('%s', '%s', %f)`
	addAuctionQueryString := fmt.Sprintf(addAuctionQuery, auctionID, bidderID, price)
	fmt.Println(addAuctionQueryString)

	insert, err := db.Query(addAuctionQueryString)
	if err != nil {
		log.Error.Println("Error occured while inserting auction winner details in the database", err)
	}

	insert.Close()
}
