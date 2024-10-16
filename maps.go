package cells

type Maps interface {
	Get(x, y int) bool
	Set(x, y int, s bool)
	String() string
	Clone() Maps
	Width() int
	Height() int
}

func Copy(dist, src Maps) {
	dw := dist.Width()
	dh := dist.Height()

	sw := src.Width()
	sh := src.Height()

	woff := (dw - sw) / 2
	hoff := (dh - sh) / 2

	for i := 0; i < dw; i++ {
		for j := 0; j < dh; j++ {
			if i < 0 || j < 0 || i >= sw || j >= sh {
				continue
			}

			x := i + woff
			y := j + hoff
			if x < 0 || y < 0 || x >= dw || y >= dh {
				continue
			}

			dist.Set(x, y, src.Get(i, j))
		}
	}
}
