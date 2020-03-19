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

func TestGetActiveRegisteredBidders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := mock.NewRows([]string{"id", "title"}).
		AddRow(1, "one").
		AddRow(2, "two")

	mock.ExpectQuery("SELECT id, domain FROM bidders where online = 1").WillReturnRows(rows)
	biddersMap, err := GetActiveRegisteredBidders(db)

	fmt.Println(biddersMap)
}
func TestGetActiveRegisteredBiddersShouldGiveError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, domain FROM bidders where online = 1").WillReturnError(fmt.Errorf("some error"))
	biddersMap, err := GetActiveRegisteredBidders(db)

	fmt.Println(biddersMap, err)
}

func TestRegisterBidder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var bidderID = "24567"
	var host = "localhost_test"
	rows := mock.NewRows([]string{""})
	mock.ExpectQuery("INSERT INTO bidders (.+) VALUES (.+) on duplicate key update domain = (.+), online = true;").
		WithArgs(bidderID, host, host).
		WillReturnRows(rows)
	err = RegisterBidder(db, bidderID, host)
	fmt.Println(err)
}
func TestRegisterBidderShouldGiveError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var bidderID = "24567"
	var host = "localhost_test"
	mock.ExpectQuery("INSERT INTO bidders (.+) VALUES (.+) on duplicate key update domain = (.+), online = true;").
		WithArgs(bidderID, host, host).
		WillReturnError(fmt.Errorf("some error"))
	err = RegisterBidder(db, bidderID, host)
}
