package constants

import (
	"database/sql"
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
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func SetConstants() {
	Config.setValuesFromConfig()
}
