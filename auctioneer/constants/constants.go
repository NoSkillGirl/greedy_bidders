package constants

import (
	"database/sql"
	"os"

	logger "github.com/NoSkillGirl/greedy_bidders/auctioneer/log"
	_ "github.com/go-sql-driver/mysql"
)

// Database - for reading database config
type Database struct {
	UserName string
	Password string
	Port     string
	Name     string
	Host     string
}

// DbConfig - available through out the application
var DbConfig Database

func (db *Database) setValuesFromConfig() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		logger.Log.Debug(`Make sure all these environment variables are set. 
		1. DB_HOST
		2. DB_PORT
		3. DB_USER
		4. DB_PASS
		5. DB_NAME

		`)

		logger.Log.Error("Example: DB_HOST=localhost DB_PORT=3306 DB_USER=pooja DB_PASS=oreo DB_NAME=greedy_bidder go run main.go")
		panic("Environment variables not set")
	}

	db.Host = dbHost
	db.Port = dbPort
	db.UserName = dbUser
	db.Password = dbPass
	db.Name = dbName
}

// GetDatabaseConnection - for getting database connection
func (db *Database) GetDatabaseConnection() (con *sql.DB) {
	dbDriver := "mysql"
	con, err := sql.Open(dbDriver, db.UserName+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.Name)
	if err != nil {
		logger.Log.Error("error in opening db connection")
		panic(err.Error())
	}
	return con
}

// One time migrations needed for bootstraping database, will be skipped from the second time.
func (db *Database) migrations() {
	dbCon := db.GetDatabaseConnection()
	_, err := dbCon.Exec(`create table if not exists auctions
	(
		id               varchar(200) not null,
		winner_bidder_id varchar(200) null,
		price            float        null,
		constraint auctions_id_uindex unique (id),
		constraint auctions_pk PRIMARY KEY (id)
	);`)
	if err != nil {
		logger.Log.Error("Unable to create auctions table")
		panic(err)
	}

	_, err = dbCon.Exec(`create table if not exists bidders
	(
		id     varchar(200) not null,
		domain varchar(255) null,
		online tinyint(1)   null,
		constraint bidder_id_uindex unique (id),
		constraint bidder_pk primary key (id)
	);`)
	if err != nil {
		logger.Log.Error("Unable to create bidders table")
		panic(err)
	}
}

// var Log *logus.Entry

// SetConstants - this method is exposed for the main package to setup
func SetConstants() {
	// Setting up Logger
	// logger.log.SetupLogger()
	// Read values from config
	DbConfig.setValuesFromConfig()
	// Run Migrations
	DbConfig.migrations()
}
