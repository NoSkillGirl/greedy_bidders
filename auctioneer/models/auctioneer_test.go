package models

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func init() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_NAME", "greedy_bidder")
}

func TestDeclareWinner(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var auctionID = "auc1"
	var bidderID = "localhost_test"
	var price = 56.86
	rows := mock.NewRows([]string{""})
	mock.ExpectQuery("insert into auctions (.+) VALUES (.+)").
		WithArgs(auctionID, bidderID, price).
		WillReturnRows(rows)
	err = DeclareWinner(db, auctionID, bidderID, price)
	fmt.Println(err)
}

func TestDeclareWinnerShouldGiveError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var auctionID = "auc1"
	var bidderID = "localhost_test"
	var price = 56.86
	mock.ExpectQuery("insert into auctions (.+) VALUES (.+)").
		WithArgs(auctionID, bidderID, price).
		WillReturnError(fmt.Errorf("some error"))
	err = DeclareWinner(db, auctionID, bidderID, price)
	fmt.Println(err)
}
