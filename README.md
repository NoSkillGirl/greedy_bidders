# Greedy Bidders

This repo represents an auction system which can auction bids and select the winning bid such that it always responds before a specific time interval.

## Prerequisites

- Docker, Docker Compose
- Golang

## Getting Up and Running

- `git clone https://github.com/NoSkillGirl/greedy_bidders.git`
- `cd greedy_bidders`
- `docker-compose rm -f && docker-compose up --remove-orphans --build`

## API Documentation

This repo is having 4 public endpoints:

1. New Auction
2. RegisterBidder
3. GetActiveRegisteredBidders
4. BidderRequest

### NewAuction (GET Request)

#### Request

- Type: JSON
- PARMAS: auction_id

Example:

```JSON
{
	"auction_id": "PEN656"
}
```

#### Response

- Type: JSON
- Response Structure

```JSON
{
	"bidder_id": "id of the bidder",
	"price": "bid bidded by the bidder"
}
```

Example:

```JSON
{
	"bidder_id": "0905b10e-5e28-11ea-9aa1-0242ac130004",
	"price": 75.65
}
```

### RegisterBidder (POST Request)

#### Request

- Type: JSON
- PARMAS: bidder_id, host

Example:

```JSON
{
	"bidder_id": "0905b10e-5e28-11ea-9aa1-0242ac130004",
	"host": "http://localhost:8090"
}
```

#### Response

- Type: JSON
- Response Structure

```JSON
{}
```

### GetActiveRegisteredBidders (POST Request)

#### Request

- Type: JSON
- PARMAS: Empty JSON

Example:

```JSON
{}
```

#### Response

- Type: JSON
- Response Structure

```JSON
{
    "bidder_ids": [
        "id of bidder1"
        "id of bidder2"
        "id of bidder3"
    ]
}
```

Example:

```JSON
{
    "bidder_ids": [
        "0905b10e-5e28-11ea-9aa1-0242ac130004"
        "8f392adc-5dff-11ea-bb7c-a683e76d0373"
        "9a4ff752-5dff-11ea-ab32-a683e76d0373"
    ]
}
```

### BidderRequest (GET Request)

#### Request

- Type: JSON
- PARMAS: auction_id

Example:

```JSON
{
	"auction_id": "test"
}
```

#### Response

- Type: JSON
- Response Structure

```JSON
{
	"bidder_id": "id of the bidder",
	"price": "bid bidded by the bidder"
}
```

Example:

```JSON
{
	"bidder_id": "0905b10e-5e28-11ea-9aa1-0242ac130004",
	"price": 75.65
}
```
