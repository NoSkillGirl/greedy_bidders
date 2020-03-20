package models

import (
	"database/sql"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
	logger "github.com/NoSkillGirl/greedy_bidders/log"
)

// DeclareWinner - stores in the database about the auction winner and the price
func DeclareWinner(db *sql.DB, auctionID string, bidderID string, price float64) error {
	if db == nil {
		db = constants.DbConfig.GetDatabaseConnection()
	}

	insert, err := db.Query(
		`insert into auctions (id, winner_bidder_id, price) VALUES (?, ?, ?)`,
		auctionID, bidderID, price,
	)

	if err != nil {
		logger.Log.Error("Error occured while inserting auction winner details in the database", err)
		return err
	}
	insert.Close()
	return nil
}
