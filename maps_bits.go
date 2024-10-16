package cells

import (
	"bytes"

	"github.com/wzshiming/bitart"
)

type mapsBits []bitart.Bits

func NewMapsBits(width, height int) Maps {
	h := height / 4
	w := width / 2
	if height%4 != 0 {
		h++
	}
	if width%2 != 0 {
		w++
	}
	m := make(mapsBits, h)
	for i := range m {
		m[i] = make(bitart.Bits, w)
	}
	return m
}

func (m mapsBits) Get(x, y int) bool {
	index := (x%2)*4 + (y%4)%4
	i := y / 4
	j := x / 2
	return (m[i][j] & (1 << index)) != 0
}

func (m mapsBits) Set(x, y int, s bool) {
	index := (x%2)*4 + (y%4)%4
	i := y / 4
	j := x / 2
	if s {
		m[i][j] |= 1 << index
	} else {
		m[i][j] &^= 1 << index
	}
}

func (m mapsBits) String() string {
	buf := bytes.Buffer{}
	for i, raw := range m {
		buf.WriteString(raw.String())
		if i != len(m)-1 {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}

func (m mapsBits) Clone() Maps {
	n := make(mapsBits, 0, len(m))
	for _, row := range m {
		r := make(bitart.Bits, len(row))
		copy(r, row)
		n = append(n, r)
	}
	return n
}

func (m mapsBits) Width() int {
	return len(m[0]) * 2
}

func (m mapsBits) Height() int {
	return len(m) * 4
}
