package budget

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

//https://www.evanjones.ca/floating-point-money.html#:~:text=A%2064%2Dbit%20floating%2Dpoint,digits%20after%20the%20decimal%20place.&text=Floating%2Dpoint%20numbers%20seem%20like,in%20most%20cases%2C%20they%20will.

// Uses float64 which for a test is generally good enough but should be switched to arbitrary precision decimals
//fixed point or any budget handling library See:
//https://github.com/ericlagergren/decimal
//https://github.com/Rhymond/go-money

func (m currency) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

type Balance struct {
	mu sync.Mutex
	balance, out currency

	sweepInterval time.Duration // Todo: implement a proper store
	sweepMinTTL   uint64 // Todo

	unackd map[string] currency
}

/*sweepInterval time.Duration
sweepMinTTL   uint64

data map[string]*bucket*/

func New(initial currency) (a Balance) {
	return Balance{balance: initial}
}

// Get gets the current limit and remaining tokens for the provided key. It
// does not change any of the values.
func (a *Balance) Get(ctx context.Context, key string) (balance, outstanding currency, err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.balance, a.out, err
}

func (a *Balance) Take(ctx context.Context, key string, amount currency) (balance, outstanding currency, err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(key) == 0 {
		return 0, 0, errors.New("no key provided")
	}
	if a.balance < amount {
		return 0, 0, errors.New("not enough balance")
	}

	a.unackd[key] = amount
	a.balance -= amount
	a.out += amount

	return a.balance, a.out, nil
}

func (a *Balance) Settle(ctx context.Context, key string) (balance, outstanding currency, err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	out, in := a.unackd[key]
	if in != true  {
		return 0, 0, fmt.Errorf("no outstanding for key: %s", key)
	}

	a.out -= out
	delete(a.unackd, key)

	return a.balance, a.out, nil
}