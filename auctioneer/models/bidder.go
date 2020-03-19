package models

import (
	"database/sql"
	"fmt"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
	logger "github.com/NoSkillGirl/greedy_bidders/log"
)

// GetActiveRegisteredBidders - returns a map of bidder_id and domain
func GetActiveRegisteredBidders(db *sql.DB) (map[string]string, error) {
	biddersMap := make(map[string]string)
	if db == nil {
		fmt.Println("Coming inside")
		db = constants.DbConfig.GetDatabaseConnection()
	}

	rows, err := db.Query("SELECT id, domain FROM bidders where online = 1")

	if err != nil {
		fmt.Println(err)
		logger.Log.Error("Error in selecting the bidders from database", err)
		return biddersMap, err
	}

	for rows.Next() {
		var bidderID, bidderDomain string
		err = rows.Scan(&bidderID, &bidderDomain)
		if err != nil {
			logger.Log.Error("Error in Scaning from the databsae", err)
		} else {
			biddersMap[bidderID] = bidderDomain
		}
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		logger.Log.Error("Error in Scaning from the databsae", err)
	}

	rows.Close()

	return biddersMap, err
}

// RegisterBidder - for storing bidder information in the database
func RegisterBidder(db *sql.DB, bidderID string, host string) error {
	if db == nil {
		db = constants.DbConfig.GetDatabaseConnection()
	}

	insert, err := db.Query(
		`INSERT INTO bidders (id, domain, online) VALUES ('$1', '$2', true) on duplicate key update domain = '$3', online = true;`,
		bidderID, host, host,
	)

	// if there is an error inserting, handle it
	if err != nil {
		logger.Log.Errorf(`Failed to insert bidder id in the database %v`, err)
		return err
	}

	insert.Close()
	return nil
}
