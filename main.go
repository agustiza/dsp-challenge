package main

import (
	"fmt"
	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
	"log"
	"net/http"
	"time"
	"wildlife-challenge/handlers"
	"wildlife-challenge/middleware"
)


func main() {

	// If you are reading this I'm really sorry
	// This repo is pretty messy code and not really how I code usually.
	// I had to both learn go on the go (heh) and implement a dsp in a day. Also generics when pls.
	s, s2 := configLimiters()

	http.Handle("/bid", middleware.Wrap(&handlers.BidHandler{}, []limiter.Store{s, s2}))
	http.Handle("/imp", &handlers.ImpHandler{})

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func configLimiters() (limiter.Store, limiter.Store) {
	s, err := memorystore.New(&memorystore.Config{
		Tokens:   5,
		Interval: time.Minute * 1,
	})

	if err != nil { // TODO: should use github.com/hashicorp/go-multierror
		log.Fatal(err)
	}

	s2, err := memorystore.New(&memorystore.Config{
		Tokens:   10,
		Interval: time.Minute * 3,
	})

	if err != nil {
		log.Fatal(err)
	}
	return s, s2
}