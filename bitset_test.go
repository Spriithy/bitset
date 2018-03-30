package bitset

import "testing"

const (
	A = iota
	B
	C
	D
	E
	n
)

func TestNew(t *testing.T) {
	bs := New(n)
	if len(bs) != n || cap(bs) != n {
		t.Fail()
	}
}

func TestInitial(t *testing.T) {
	bs := Initial(n, A, C, E)
	if len(bs) != n || cap(bs) != n {
		t.Fail()
	}

	if !bs.All(A, C, E) {
		t.Fail()
	}

	if bs.Any(B, D) {
		t.Fail()
	}
}
