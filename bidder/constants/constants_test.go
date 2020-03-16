package constants

import (
	"os"
	"regexp"
	"testing"
)

func init() {
	os.Setenv("AUCTIONEER_URL", "localhost:8080")
	os.Setenv("PORT", "8081")
	os.Setenv("DELAY", "150")
	os.Setenv("HOST", "localhost")
}

func TestUUID(t *testing.T) {
	SetConstants()
	bidderID := GetBidderID()
	if len(bidderID) != 36 {
		t.Error("bidderID - UUID - string length doesnot match")
	}
	if !isValidUUID(bidderID) {
		t.Errorf(`bidderID - UUID - not valid UUID V4 %s`, bidderID)
	}
}

func TestSetBidderID(t *testing.T) {
	os.Setenv("BIDDER_ID", "5b7c8234-2570-4d83-8790-5f09757f6339")
	setBidderID()
	if BidderID != "5b7c8234-2570-4d83-8790-5f09757f6339" {
		t.Errorf(`Expected BIDDER_ID: 5b7c8234-2570-4d83-8790-5f09757f6339 got %s`, BidderID)
	}
}

func TestSetConfig(t *testing.T) {
	// check if all the config variables are set
	SetConstants()

	if Config.Delay != 150 {
		t.Errorf(`Expected delay should be 150 but got: %v`, Config.Delay)
	}

	if Config.Host != "localhost" {
		printError(t, "Host", "localhost", Config.Host)
	}

	if Config.Port != "8081" {
		printError(t, "Port", "8081", Config.Port)
	}

	if Config.AuctioneerRegisterURL != "http://localhost:8080/RegisterBidder" {
		printError(t, "Auctioneer register URL", "http://localhost:8080/RegisterBidder", Config.AuctioneerRegisterURL)
	}

	// check for the edge cases
	os.Setenv("AUCTIONEER_URL", "")
	configShouldNotBeSet(t, Config)
	os.Setenv("AUCTIONEER_URL", "localhost:8080")

	os.Setenv("PORT", "")
	configShouldNotBeSet(t, Config)
	os.Setenv("PORT", "8081")

	os.Setenv("DELAY", "")
	configShouldNotBeSet(t, Config)
	os.Setenv("DELAY", "150")

	os.Setenv("HOST", "")
	configShouldNotBeSet(t, Config)
	os.Setenv("HOST", "localhost")

}

func printError(t *testing.T, test string, expected string, got string) {
	t.Errorf(`Expected %v should be %v but got %v`, test, expected, got)
}

func configShouldNotBeSet(t *testing.T, config Conf) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	SetConstants()
	if config.AuctioneerRegisterURL != "" || config.Delay != 0 || config.Host != "" || config.Port != "" {
		t.Errorf(`Expected Config should be nil but got %v`, Config)
	}
}

func isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// Benchmarking Testing

func BenchmarkUUID(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SetConstants()
	}
}
