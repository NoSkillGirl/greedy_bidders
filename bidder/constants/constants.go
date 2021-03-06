package constants

import (
	"fmt"
	"os"
	"strconv"

	logger "github.com/NoSkillGirl/greedy_bidders/log"
	"github.com/google/uuid"
)

// GetBidderID - returns bidderID
func GetBidderID() string {
	return BidderID
}

func setBidderID() {
	preSetBidderID := os.Getenv("BIDDER_ID")
	if preSetBidderID != "" {
		BidderID = preSetBidderID
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		logger.Log.Error("UUID Generation Error")
	}
	BidderID = id.String()

	fmt.Printf(`AutoGenerated BidderID: %s`, BidderID)
}

// Conf - to store config
type Conf struct {
	Port                  string `yaml:"port"`
	Delay                 int    `yaml:"delay"`
	AuctioneerRegisterURL string `yaml:"auctioneer_register_url"`
	Host                  string `yaml:"host"`
}

func (c *Conf) setValuesFromConfig() {

	auctioneerURL := os.Getenv("AUCTIONEER_URL")
	appHost := os.Getenv("HOST")
	appPort := os.Getenv("PORT")
	bidderDelay := os.Getenv("DELAY")

	if auctioneerURL == "" || appHost == "" || appPort == "" || bidderDelay == "" {
		logger.Log.Error(`Make sure all these environment variables are set. 
		1. AUCTIONEER_URL
		2. PORT
		3. DELAY
		4. HOST

		`)

		logger.Log.Error("Example: AUCTIONEER_URL=localhost:8081 HOST=localhost PORT=8080 DELAY=150 go run main.go")
		// os.Exit(1)
		panic("ENV not set")
	}

	c.AuctioneerRegisterURL = "http://" + auctioneerURL + "/RegisterBidder"
	c.Port = appPort
	bidderDelayInt, _ := strconv.Atoi(bidderDelay)
	c.Delay = bidderDelayInt
	c.Host = appHost
}

// GetConf - GetConfig
func (c *Conf) GetConf() *Conf {
	return c
}

// BidderID - For this Bidder
var BidderID string

// Config - Config
var Config Conf

// SetConstants - to be called once before server start
func SetConstants() {
	setBidderID()
	Config.setValuesFromConfig()
}
