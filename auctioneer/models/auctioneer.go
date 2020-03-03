package models

const mySQLHost = "127.0.0.1:3306"

var mySQLConnection = fmt.Sprintf("root:@tcp(%s)/greedy_bidder", mySQLHost)

// NewBiddingRound - Creates a new bidding round
func NewBiddingRound(auctionID string) error {
	db, err := sql.Open("mysql", mySQLConnection)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return nil
}

// connect to mysql
// create a table auctions
// row create
