package bid

import (
	"fmt"
	"math/rand"
)

type User struct {

}

type BidReq struct {
	min int64
	max int64
}

func NewReq(min, max int64) BidReq {
	return BidReq{min: min, max: max}
}

type Bid struct {
	Amount int64
}

type Bidder interface {
	ShouldBid(req BidReq, u User) (Bid, error)
}

type randBidderSrvc struct {

}

func New() Bidder {
	return &randBidderSrvc{}
}

func (b *randBidderSrvc) ShouldBid(req BidReq, u User) (bid Bid, err error) {
	bid = Bid{}

	if req.min >= req.max {
		return bid, fmt.Errorf("req.min %d >= req.max %d", req.min, req.max)
	}

	bid.Amount = rand.Int63n(req.max - req.min) + req.min
	return bid, nil
}


