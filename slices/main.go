package slices

import "github.com/kendfss/rules"

// Returns the outcome of successive applications of
// a function, f, as a binary operator over the slice, s.
func Reduce[S ~[]E, E any](f func(E, E) E, s S) (out E) {
	switch len(s) {
	case 0:
		out = *new(E)
	case 1:
		out = s[0]
	default:
		out = s[0]
		for _, e := range s[1:] {
			out = f(out, e)
		}
	}
	return out
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}

	// return append(S([]E{}), s...)

	// out := make(S, len(s))
	// for i, e := range s {
	// 	out[i] = e
	// }
	// return out

	return append(S{}, s...)
}

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s; it does not create a new slice.
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

// Compacted clones the slice and runs Compact on said clone
func Compacted[S ~[]E, E comparable](s S) S {
	c := Clone(s)
	return Compact(c)
}

// Indices of a slice with given length
func Rangen[T rules.Integer](stop T) []T {
	return Range(0, stop, 1)
}

// Consecutive ints, including start, smaller than stop, and separated by step
func Range[T rules.Integer](start, stop, step T) (out []T) {
	if stop <= start && step >= 0 {
		return
	}
	for i := start; i < (stop); i += step {
		out = append(out, i)
	}

	return out
}
