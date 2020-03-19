module github.com/NoSkillGirl/greedy_bidders/auctioneer

go 1.14

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1
	github.com/sirupsen/logrus v1.4.2 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	github.com/NoSkillGirl/greedy_bidders/log
)
replace github.com/NoSkillGirl/greedy_bidders/log => ../log

