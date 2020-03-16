package constants

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_NAME", "greedy_bidder")
}

func TestSetValuesFromConfig(t *testing.T) {
	DbConfig.setValuesFromConfig()
	// check if all the config variables are set

	if DbConfig.UserName != "root" {
		t.Errorf(`Expected delay should be 150 but got: %v`, DbConfig.UserName)
	}

	if DbConfig.Host != "localhost" {
		printError(t, "Host", "localhost", DbConfig.Host)
	}

	if DbConfig.Port != "3306" {
		printError(t, "Port", "3306", DbConfig.Port)
	}

	if DbConfig.Name != "greedy_bidder" {
		printError(t, "Database Name", "greedy_bidder", DbConfig.Name)
	}

	// check for the edge cases
	os.Setenv("DB_HOST", "")
	configShouldNotBeSet(t, DbConfig)
	os.Setenv("DB_HOST", "localhost")

	os.Setenv("DB_PORT", "")
	configShouldNotBeSet(t, DbConfig)
	os.Setenv("DB_PORT", "3306")

	os.Setenv("DB_USER", "")
	configShouldNotBeSet(t, DbConfig)
	os.Setenv("DB_USER", "root")

	os.Setenv("DB_NAME", "")
	configShouldNotBeSet(t, DbConfig)
	os.Setenv("DB_NAME", "greedy_bidder")

}

func printError(t *testing.T, test string, expected string, got string) {
	t.Errorf(`Expected %v should be %v but got %v`, test, expected, got)
}

func configShouldNotBeSet(t *testing.T, dbConfig Database) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	DbConfig.setValuesFromConfig()
	if dbConfig.Host == "" || dbConfig.Name == "" || dbConfig.Port == "" || dbConfig.UserName == "" {
		t.Errorf(`Expected Config should be nil but got %v`, dbConfig)
	}
}
