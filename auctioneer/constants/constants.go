package constants

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NoSkillGirl/greedy_bidders/auctioneer/log"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

// Database - database config
type Database struct {
	Database DatabaseConf `yaml:database`
}

// DatabaseConf - for reading database config
type DatabaseConf struct {
	UserName string `yaml:user_name`
	Password string `yaml:password`
	Port     string `yaml:"port"`
	Name     string `yaml:name`
	Host     string `yaml:host`
}

// Config - available through out the application
var Config Database

func (c *Database) setValuesFromConfig() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		panic(err)
	}
}

// GetDatabaseConnection - for getting database connection
func (c *Database) GetDatabaseConnection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := c.Database.UserName
	dbPass := c.Database.Password
	dbName := c.Database.Name
	dbHost := c.Database.Host
	dbPort := c.Database.Port
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// One time migrations needed for bootstraping database, will be skipped from the second time.
func (c *Database) migrations() {
	db := c.GetDatabaseConnection()
	_, err := db.Exec(`create table if not exists auctions
	(
		id               varchar(200) not null,
		winner_bidder_id varchar(200) null,
		price            float        null,
		constraint auctions_id_uindex unique (id),
		constraint auctions_pk PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Error.Println("Unable to create auctions table")
		panic(err)
	}

	_, err = db.Exec(`create table if not exists bidders
	(
		id     varchar(200) not null,
		domain varchar(255) null,
		online tinyint(1)   null,
		constraint bidder_id_uindex unique (id),
		constraint bidder_pk primary key (id)
	);`)
	if err != nil {
		log.Error.Println("Unable to create bidders table")
		panic(err)
	}
}

// SetConstants - this method is exposed for the main package to setup
func SetConstants() {
	// Setting up Logger
	log.SetupLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	// Read values from config
	Config.setValuesFromConfig()
	// Run Migrations
	Config.migrations()
}
