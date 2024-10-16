package cells

import (
	"bytes"
)

type mapsChars struct {
	maps  [][]byte
	alive byte
	empty byte
}

func NewMapsChars(alive, empty byte, width, height int) Maps {
	m := make([][]byte, height)
	for i := range m {
		m[i] = make([]byte, width)
	}
	return &mapsChars{
		maps:  m,
		alive: alive,
		empty: empty,
	}
}

func NewMapsCharsWithData(alive, empty byte, data [][]byte) Maps {
	w := 0
	for _, r := range data {
		if len(r) > w {
			w = len(r)
		}
	}

	m := make([][]byte, len(data))

	for i, r := range data {
		if s := w - len(r); s != 0 {
			m[i] = append(data[i], bytes.Repeat([]byte{empty}, s)...)
		} else {
			m[i] = data[i]
		}
	}

	return &mapsChars{
		maps:  m,
		alive: alive,
		empty: empty,
	}
}

func (m mapsChars) Get(x, y int) bool {
	return m.maps[y][x] != m.empty
}

func (m mapsChars) Set(x, y int, s bool) {
	if s {
		m.maps[y][x] = m.alive
	} else {
		m.maps[y][x] = m.empty
	}
}

func (m mapsChars) String() string {
	buf := bytes.Buffer{}
	for i, raw := range m.maps {
		buf.Write(raw)
		if i != len(m.maps)-1 {
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}

func (m mapsChars) Clone() Maps {
	n := make([][]byte, 0, len(m.maps))
	for _, row := range m.maps {
		r := make([]byte, len(row))
		copy(r, row)
		n = append(n, r)
	}
	return &mapsChars{
		maps:  n,
		alive: m.alive,
		empty: m.empty,
	}
}

func (m mapsChars) Width() int {
	return len(m.maps[0])
}

func (m mapsChars) Height() int {
	return len(m.maps)
}
