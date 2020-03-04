package constants

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

func GetBidderId() string {
	return BidderID
}

func setBidderId() {
	fmt.Println("Setting Bidder Id")
	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("UUID Generation Error")
	}
	BidderID = id.String()
}

type Conf struct {
	Port                  string `yaml:"port"`
	Delay                 int64  `yaml:"delay"`
	AuctioneerRegisterURL string `yaml:"auctioneer_register_url"`
}

func (c *Conf) setValuesFromConfig() {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func (c *Conf) GetConf() *Conf {
	return c
}

var BidderID string
var C Conf

func SetConstants() {
	setBidderId()
	C.setValuesFromConfig()
}
