package constants

import (
	"fmt"
	"strconv"

	// "io/ioutil"
	// "log"
	"os"

	"github.com/google/uuid"
	// "gopkg.in/yaml.v2"
)

// GetBidderID - returns bidderID
func GetBidderID() string {
	return BidderID
}

func setBidderID() {
	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("UUID Generation Error")
	}
	BidderID = id.String()
}

// Conf - to store config
type Conf struct {
	Port                  string `yaml:"port"`
	Delay                 int    `yaml:"delay"`
	AuctioneerRegisterURL string `yaml:"auctioneer_register_url"`
	Host                  string `yaml:"host"`
}

func (c *Conf) setValuesFromConfig() {
	// yamlFile, err := ioutil.ReadFile("config.yaml")
	// if err != nil {
	// 	log.Printf("yamlFile.Get err   #%v ", err)
	// }
	// err = yaml.Unmarshal(yamlFile, c)
	// if err != nil {
	// 	log.Fatalf("Unmarshal: %v", err)
	// }
	// auctioneer_register_url
	c.AuctioneerRegisterURL = "http://" + os.Getenv("AUCTIONEER_URL") + "/RegisterBidder"
	c.Port = os.Getenv("PORT")

	i, _ := strconv.Atoi(os.Getenv("DELAY"))
	c.Delay = i
	c.Host = os.Getenv("HOST")
}

// GetConf - GetConfig
func (c *Conf) GetConf() *Conf {
	return c
}

// BidderID - For this Bidder
var BidderID string

// C - Config
var C Conf

// SetConstants - to be called once before server start
func SetConstants() {
	setBidderID()
	C.setValuesFromConfig()
}
