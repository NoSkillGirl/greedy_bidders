NAME := auctioneer
DOCKER_REPO := ajoop/$(NAME)

default: run

run:
	@go run .

docker:
	docker build -t $(NAME) -f Dockerfile .

docker-push: TAG ?= latest
docker-push: docker
	docker tag $(NAME) $(DOCKER_REPO):$(TAG)
	docker push $(DOCKER_REPO):$(TAG)

doc:
	go doc -all github.com/NoSkillGirl/greedy_bidders/auctioneer/constants
	go doc -all github.com/NoSkillGirl/greedy_bidders/auctioneer/controllers
	go doc -all github.com/NoSkillGirl/greedy_bidders/auctioneer/log
	go doc -all github.com/NoSkillGirl/greedy_bidders/auctioneer/models
	go doc -all github.com/NoSkillGirl/greedy_bidders/auctioneer/routers