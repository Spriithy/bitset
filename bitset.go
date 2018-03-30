package bitset

// Type Bitset is used as a collection of bits, that can each represent a state
// It is not optimized for space (otherwise I would have used an integer slice
// to represent all possible bits, and grow it as needed).
//
// It is compatible with untyped integer constants for ease of use with
// enum-like declared values.
type Bitset []bool

// New returns a new, zeroed, Bitset of the given size.
func New(nBits int) Bitset {
	return make(Bitset, nBits)
}

// Initial returns a new Bitset, with the given bits set.
func Initial(nBits int, bits ...int) Bitset {
	bs := New(nBits)
	for _, bit := range bits {
		if bit < nBits {
			bs[bit] = true
		}
	}
	return bs
}

// Resize returns a copy of the actual Bitset, with the given capacity.
func (bs Bitset) Resize(nBits int) Bitset {
	nbs := New(nBits)
	copy(nbs, bs)
	return nbs
}

// Grow returns a copy of the actual Bitset grown by the given amount of bits.
func (bs Bitset) Grow(nBits int) Bitset {
	return bs.Resize(len(bs) + nBits)
}

// Shrink returns a copy of the actual Bitset shrunk by the given amount of bits.
func (bs Bitset) Shrink(nBits int) Bitset {
	return bs.Resize(len(bs) - nBits)
}

// Set is used to set a series of bits.
func (bs Bitset) Set(bits ...int) {
	nBits := len(bs)
	for _, bit := range bits {
		if bit < nBits {
			bs[bit] = true
		}
	}
}

// Clear is used to clear a series of bits.
func (bs Bitset) Clear(bits ...int) {
	nBits := len(bs)
	for _, bit := range bits {
		if bit < nBits {
			bs[bit] = false
		}
	}
}

// Flip is used to flip (toggle) a series of bits.
func (bs Bitset) Flip(bits ...int) {
	nBits := len(bs)
	for _, bit := range bits {
		if bit < nBits {
			bs[bit] = !bs[bit]
		}
	}
}

// All returns whether all given bits are set.
func (bs Bitset) All(bits ...int) bool {
	nBits := len(bs)
	for _, bit := range bits {
		if bit < nBits && !bs[bit] {
			return false
		}
	}
	return true
}

// Any returns whether any given bit is set.
func (bs Bitset) Any(bits ...int) bool {
	nBits := len(bs)
	for _, bit := range bits {
		if bit < nBits && bs[bit] {
			return true
		}
	}
	return false
}

// String returns the Bitset's string representation.
func (bs Bitset) String() string {
	str := "Bitset{"
	for _, bit := range bs {
		switch bit {
		case true:
			str += "1, "
		case false:
			str += "0, "
		}
	}
	return str[:len(str)-2] + "}"
}
