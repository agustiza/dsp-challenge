# Wip DSP Challenge

Challenge: Implement a crude DSP in a day (while learning go).
 * Implement /bid route that recieves bid information
 * Implement bidding
 * Delegate to bidder implementation
 * Set up redis
 * Properly handle currency through fixed point math/big decimals
 * Implement impressions callback count route / count (impid=102)
 * Ideally load balancers should assign requests to nodes by round-robin or hash sharded userId
 * Ideally avoid locking on global balance by partitioning balance by nodes?
ous
 * At most 5 impressions per minute, per user basis
 * At most 10 impressions every 3 minutes, per user basis
 * A maximum of 10 USD to spend per day, system wide

## Caveats

(**NOT PRODUCTION READY**)
* Absolutely not ready for production usage. Nor anything!
* It never checks the Accept header of the request, the errors responses are lacking additional details and proper envelopes. 
* HTTP server also never validates the data the client sends
* Should eventually switch to an already made HttpRestApi server such as Gin, Mux or Chi
* Should dockerize
* Also has unbounded memory caches which should eventually be swapped for Redis
* Should use probably implement a sliding window sorted set for Balance/Ratelimits like Redis ZSet with a TTL
* Bidding is completely random, bidding service should be split and should include a learning phase from historical data
* Learning phase could use Bayesian statistics, neural networks, association rule learning,
  clustering, or other techniques to calculate the bonuses from historical data. These other methods
  may offer more mathematically rigorous results, which may result in better
  CTRs, and ultimately more money earned. See `/rtb-opt.pdf`
* Random uses Math.Rand for bid generation which could lead to bid prediction by an adversary.
* I will eventually move the balance to its own Redis server. I'm worried on lock contention
* Redis IS NOT not configured for durability guarantees. No SLA guarantee for money handling system. Discrepancies in balance are expected.
* This could be remediated by `set appendfsync` at a performance tradeoff.


# References
* [OpenRTB Specification](https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf)

* [Pitaya, very cool, very good reference](https://github.com/topfreegames/pitaya)

* [cpp DSP implementation and framework](https://github.com/venediktov/vanilla-rtb/blob/master/examples/bidder/multi_bidder.cpp)
