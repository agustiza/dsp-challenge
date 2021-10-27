package handlers

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/url"
	"wildlife-challenge/bid"
	"wildlife-challenge/model"
	"wildlife-challenge/util"
)

type BidHandler struct {
	BidReq model.BidPostReq
}

func (b *BidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Print(b.BidReq)
	fmt.Printf("\nid: %s\n", b.BidReq.Id)

	bidId := uuid.NewV4().String()

	encoder := json.NewEncoder(w)
	params := url.Values{}
	params.Add("bidId", bidId)
	params.Add("id", b.BidReq.Id)
	nurl := util.GetLocalIP() + ":8080" + "/imp?" + params.Encode()

	bidder := bid.New()
	price, err := bidder.ShouldBid(bid.NewReq(int64(b.BidReq.Imp.Bidfloor), 10), struct{}{})
	if err != nil {
		return
	}

	res := model.BidPostResponse{
		Id:    b.BidReq.Id,
		BidId: bidId,
		Bid: model.BidData{
			Price: float64(price.Amount),
			Nurl:  nurl,
		},
	}
	if err = encoder.Encode(&res); err != nil {
		fmt.Println(err)
	} else {
		return
	}
}
