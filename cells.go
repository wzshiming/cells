package cells

type Cells struct {
	width  int
	height int
	maps   Maps

	bound bool
}

type option func(*Cells)

func WithBound(b bool) option {
	return func(cells *Cells) {
		cells.bound = b
	}
}

func NewCells(maps Maps, opts ...option) *Cells {
	c := &Cells{
		maps:   maps,
		width:  maps.Width(),
		height: maps.Height(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Cells) neighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy

			if !c.bound {
				if nx < 0 {
					nx += c.width
				}
				if ny < 0 {
					ny += c.height
				}
				if nx >= c.width {
					nx -= c.width
				}
				if ny >= c.height {
					ny -= c.height
				}
			} else {
				if nx < 0 ||
					ny < 0 ||
					nx >= c.width ||
					ny >= c.height {
					continue
				}
			}

			if c.maps.Get(nx, ny) {
				count++
			}
		}
	}
	return count
}

func (c *Cells) String() string {
	return c.maps.String()
}

func (c *Cells) Maps() Maps {
	return c.maps
}

func (c *Cells) Next() bool {
	n := c.maps.Clone()
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			neighbors := c.neighbors(x, y)
			alive := c.maps.Get(x, y)
			if alive {
				if neighbors < 2 || neighbors > 3 {
					n.Set(x, y, false)
				}
			} else {
				if neighbors == 3 {
					n.Set(x, y, true)
				}
			}
		}
	}
	c.maps = n
	return true
}
