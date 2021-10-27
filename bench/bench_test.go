package bench

import (
	"os"
	"strconv"
	"testing"
	"wildlife-challenge/src"
)

func BenchmarkTestHello(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull) //Pipes to DevNull

	for i := 0; i < b.N; i++ {
		src.Hello(strconv.Itoa(i))
	}
}
