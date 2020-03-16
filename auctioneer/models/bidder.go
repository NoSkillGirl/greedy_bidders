package models

import (
	"database/sql"
	"fmt"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
	"github.com/NoSkillGirl/greedy_bidders/auctioneer/log"
)

// GetActiveRegisteredBidders - returns a map of bidder_id and domain
func GetActiveRegisteredBidders(db *sql.DB) map[string]string {
	if db == nil {
		db = constants.DbConfig.GetDatabaseConnection()
	}
	selectBidderQuery := `SELECT id, domain FROM bidders where online = 1;`

	rows, err := db.Query(selectBidderQuery)

	if err != nil {
		log.Error.Println("Error in selecting the bidders from database", err)
	}

	biddersMap := make(map[string]string)

	for rows.Next() {
		var bidderID, bidderDomain string
		err = rows.Scan(&bidderID, &bidderDomain)
		if err != nil {
			log.Error.Println("Error in Scaning from the databsae", err)
		} else {
			biddersMap[bidderID] = bidderDomain
		}
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Error.Println("Error in Scaning from the databsae", err)
	}

	rows.Close()

	return biddersMap
}

// RegisterBidder - for storing bidder information in the database
func RegisterBidder(db *sql.DB, bidderID string, host string) {
	if db == nil {
		db = constants.DbConfig.GetDatabaseConnection()
	}

	addBidderQuery := `INSERT INTO bidders (id, domain, online) VALUES ('%s', '%s', true) on duplicate key update domain = '%s', online = true;`

	addBidderQueryString := fmt.Sprintf(addBidderQuery, bidderID, host, host)
	fmt.Println(addBidderQueryString)

	insert, err := db.Query(addBidderQueryString)

	// if there is an error inserting, handle it
	if err != nil {
		log.Error.Println("Failed to insert bidder id in the database", err)
	}

	insert.Close()
}
