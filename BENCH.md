$ go test -bench . import-path-of-the-model-package

$ go test -bench . ./... # to run all tests and benchmarks under cwd

$ go test -bench Model ./... # for example

$ go test -run=NOTEST -bench=. -benchtime=100000x -benchmem -count=20 -timeout=60m  > x.txt
$ benchstat x.txt

	// start 100ms timer
	// if !ratelimited
	// if !impId
	// save impId
	// doBid
	// 	- getBalance
	//  - shouldBid(user, outstanding)
	// 		return bidAmount
	// 	save bid:bidamount:impId
	//  getBalance slidingWindow of sumOfLockedBalances last 5 minutes + sum of balance last 24 hours
	// 	balance -> lockedBalance
	// onImpression
	// 	lockedBalance -= bid:bidAmount
	//  impressionCount++

	//go test -test.bench=".*"