package handlers

import (
	"github.com/sethvargo/go-limiter"
	"net/http"
	"sync"
	"wildlife-challenge/budget"
)

type LimitHandler struct {
	mu sync.Mutex // guards n
	n  int
	text []string
	Limits []limiter.Store
	acc budget.Balance
}

func (h *LimitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
}
