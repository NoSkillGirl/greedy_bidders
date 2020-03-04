package constants

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Database struct {
	Database DatabaseConf `yaml:database`
}

type DatabaseConf struct {
	UserName string `yaml:user_name`
	Password string `yaml:password`
	Port     string `yaml:"port"`
	Name     string `yaml:name`
	Host     string `yaml:host`
}

var Config Database

func (c *Database) setValuesFromConfig() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

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

func (c *Database) createInitialTables() {
	db := c.GetDatabaseConnection()
	_, err := db.Exec(`create table auctions
	(
		id               varchar(200) not null,
		winner_bidder_id varchar(200) null,
		price            float        null,
		constraint auction_id_uindex
			unique (id)
	);`)
	if err != nil {
		// panic(err)
	}

	_, err = db.Exec(`create table bidders
	(
		id     varchar(200) not null,
		domain varchar(255) null,
		online tinyint(1)   null,
		constraint bidder_id_uindex
			unique (id)
	);`)
	if err != nil {
		// panic(err)
	}

	_, err = db.Exec(`alter table auctions add primary key (id);`)
	if err != nil {
		// panic(err)
	}

	_, err = db.Exec(`alter table bidders add primary key (id);`)
	if err != nil {
		// panic(err)
	}

}

func SetConstants() {
	Config.setValuesFromConfig()
	Config.createInitialTables()

}
