package itertools

import (
	"iter"
)

func WithPrev[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		prev, ok := next()
		if !ok {
			return
		}
		for val, ok := next(); ok; val, ok = next() {
			if !yield(prev, val) {
				return
			}
			prev = val
		}
	}
}
