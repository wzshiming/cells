package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/wzshiming/cells"
	"github.com/wzshiming/cursor"
	"golang.org/x/term"
)

var (
	interval time.Duration
	width    int
	height   int
	mapsFile string
	origin   bool
	bound    bool
)

func init() {
	var err error
	width, height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 86
		height = 24
	}
	width *= 2
	height *= 4

	flag.DurationVar(&interval, "interval", time.Second/10, "step interval")
	flag.IntVar(&width, "width", width, "maps width")
	flag.IntVar(&height, "height", height, "maps height")
	flag.StringVar(&mapsFile, "file", "", "maps file path")
	flag.BoolVar(&origin, "origin", false, "use origin")
	flag.BoolVar(&bound, "bound", false, "use bound")
	flag.Parse()
}

func main() {
	var maps cells.Maps

	if mapsFile == "" {
		maps = cells.NewMapsBits(width, height)
		// Put random cells
		count := width * height / 2

		for i := 0; i != count; i++ {
			x := rand.Intn(width)
			y := rand.Intn(height)
			maps.Set(x, y, true)
		}
	} else {
		f, err := os.ReadFile(mapsFile)
		if err != nil {
			panic(err)
		}

		data := bytes.Split(f, []byte{'\n'})
		data = clearSpaces(data)
		if strings.HasSuffix(mapsFile, ".rle") {
			data = clearComment(data, '#')
			maps, err = cells.NewMapsWithRLE(data)
			if err != nil {
				panic(err)
			}
		} else if strings.HasSuffix(mapsFile, ".cells") {
			data = clearComment(data, '!')
			maps = cells.NewMapsCharsWithData('O', '.', data)
		} else {
			maps = cells.NewMapsCharsWithData('*', ' ', data)
		}

		if !origin {
			ms := cells.NewMapsBits(width, height)
			cells.Copy(ms, maps)
			maps = ms
		}
	}

	world := cells.NewCells(maps,
		cells.WithBound(bound),
	)
	for {
		cursor.Clear()
		fmt.Printf("\n%v", world)
		time.Sleep(interval)
		if !world.Next() {
			break
		}
	}
}

func clearComment(data [][]byte, prefix byte) [][]byte {
	for i, row := range data {
		if row[0] != prefix {
			return data[i:]
		}
	}
	return data
}

func clearSpaces(data [][]byte) [][]byte {
	for i := range data {
		data[i] = bytes.TrimRightFunc(data[i], unicode.IsSpace)
	}
	return data
}
