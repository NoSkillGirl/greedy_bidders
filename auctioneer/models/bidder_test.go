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

	mock.ExpectBegin()
	mock.ExpectExec("SELECT id, domain FROM bidders where online = 1").WillReturnResult(sqlmock.NewResult(1, 1))

	biddersMap := GetActiveRegisteredBidders(db)

	fmt.Println(biddersMap)
}
