package oracle

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/kendfss/oracle/slices"
	"github.com/kendfss/rules"
)

func Compact[S ~[]E, E comparable](s S) S {
	if len(s) == 0 {
		return s
	}
	i := 1
	last := s[0]
	for _, v := range s[1:] {
		if v != last {
			s[i] = v
			i++
			last = v
		}
	}
	return s[:i]
}

func Factorial[N rules.Int](n N) N {
	switch {
	case n == 0:
		return 1
	case n > 0:
		return n * Factorial(n-1)
	default:
		return Factorial(-n)
	}
}

func Quitf(t *testing.T, msg string, args ...any) {
	t.Fatalf(msg, args...)
	os.Exit(1)
}

func Inequiv(t *testing.T, i int, have, want any) {
	Quitf(t, "#%d have %v, want %v", i, have, want)
}

// the sequence of terms generated repeatedly subtracting
// the divisor from the value and then subtracting any non-zero remainder
func Subtractions[T rules.OrderedNumber](value, divisor T) (out []T) {
	if divisor == 0 {
		return []T{value}
	}
	for value > 0 {
		if value >= divisor {
			out = append(out, divisor)
			value -= divisor
		} else {
			out = append(out, value)
			value = 0
		}
	}
	return out
}

func Triangular[T rules.Integer](n T) T {
	return n*(n+1)/2 - n
}

// recursive computation of the Triangular numbers
func TriangularR[T rules.Integer](n T) T {
	if n == 0 {
		return n
	}
	return TriangularR(n-1) + n
}

func Odds[T rules.Unsigned](n T) T {
	return (n + 1) * (n + 1)
}

func Evens[T rules.Unsigned](n T) T {
	return Odds(n) - n - 1
}

func Mkr(quantity, maximum int) (out []int) {
	for len(out) < quantity {
		out = append(out, rand.Intn(maximum))
		out = slices.Compact(out)
	}
	return out
}

func Mk(length, quantity, maxItem int) [][]int {
	out := make([][]int, quantity)
	for i := 0; i < quantity; i++ {
		out[i] = Mkr(length, maxItem)
	}
	return out
}

func Runes(arg string) []rune {
	return []rune(arg)
}

func Runes2(arg []string) (out [][]rune) {
	for _, str := range arg {
		out = append(out, []rune(str))
	}
	return out
}

func RandBool() bool {
	return rand.Int()%2 == 0
}

func RandRunes(length int) []rune {
	data := make([]int32, length)
	for i := range data {
		data[i] = rand.Int31()
	}
	return data
}

func RandNums[T rules.OrderedNumber](length int) []T {
	data := make([]T, length)
	for i := range data {
		data[i] = T(rand.Int())
	}
	return data
}

func RandBytes(length int) []byte {
	data := make([]uint8, length)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}
	return data
}

// Quotev returns a quoted string of a value without its type info
func Quotev[T any](arg T) string {
	return fmt.Sprintf("%q", fmt.Sprintf("%v", arg))
}

// Quotet returns a quoted string of a value with its type info
func Quotet[T any](arg T) string {
	return fmt.Sprintf("%q", fmt.Sprintf("%#v", arg))
}

func RandStr(maxLen int) string {
	return string(RandRunes(rand.Intn(maxLen)))
}

func RandStrs(count, maxLen int) []string {
	out := make([]string, count)
	for i := range out {
		out[i] = string(RandRunes(rand.Intn(maxLen)))
	}
	return out
}
