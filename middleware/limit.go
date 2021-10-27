package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-limiter"
	"net/http"
	"wildlife-challenge/handlers"
	"wildlife-challenge/model"
)

func Wrap(next *handlers.BidHandler, limits []limiter.Store) http.Handler {

	h := handlers.LimitHandler{
		Limits:  limits,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		bidReq := model.BidPostReq{}
		err := decoder.Decode(&bidReq)
		if err != nil {
			fmt.Println(err)
		}

		for _, l := range h.Limits {
			_, _, _, ok, _ := l.Take(ctx, bidReq.User.Id)
			if ok != true {
				fmt.Printf("limits hit by %s", bidReq.User.Id)
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
		}

		next.BidReq = bidReq
		next.ServeHTTP(w, r)
	})
}