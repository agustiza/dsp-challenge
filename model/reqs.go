package model

// Simplified OpenRTB types
// https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf

type BidPostReq struct {
	Id  string `json:"id"`
	Imp struct {
		Bidfloor float64 `json:"bidfloor"`
	} `json:"imp"`
	Device struct {
		Ip string `json:"ip"`
	} `json:"device"`
	User struct {
		Id string `json:"id"`
	} `json:"user"`
}

type BidData struct {
	Price float64    `json:"price"`
	Nurl  string `json:"nurl"`
}

type BidPostResponse struct {
	Id    string  `json:"id"`
	BidId string  `json:"bidid"`
	Bid   BidData `json:"bid"`
}

