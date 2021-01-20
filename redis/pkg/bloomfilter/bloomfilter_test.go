package bloomfilter

import (
	"fmt"
	"testing"
)

func TestCalculateFilterLength(t *testing.T) {
	m := CalFilterLength(1000000, 0.1)
	fmt.Println(m / 8 / 1024 / 1024 / 1024)
}
