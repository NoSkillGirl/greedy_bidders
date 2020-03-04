package models

import (
	"fmt"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/constants"
)

func GetActiveRegisteredBidders() map[string]string {
	db := constants.Config.GetDatabaseConnection()

	selectAllUsersQuery := `SELECT id, domain FROM bidders where online = 1;`

	// perform a db.Query select
	rows, err := db.Query(selectAllUsersQuery)

	// if there is an error, handle it
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	biddersMap := make(map[string]string)

	for rows.Next() {

		var bidderID, bidderDomain string

		err = rows.Scan(&bidderID, &bidderDomain)

		if err != nil {
			// handle this error
			panic(err)
		}

		biddersMap[bidderID] = bidderDomain

	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return biddersMap
}

func RegisterBidder(BidderId string, host string) {
	db := constants.Config.GetDatabaseConnection()

	addBidderQuery := `INSERT INTO bidders (id, domain, online) VALUES ('%s', '%s', true)`

	addBidderQueryString := fmt.Sprintf(addBidderQuery, BidderId, host)
	fmt.Println(addBidderQueryString)

	// perform a db.Query insert
	insert, err := db.Query(addBidderQueryString)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
