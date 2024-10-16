package cells

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// https://conwaylife.com/wiki/Run_Length_Encoded

func NewMapsWithRLE(data [][]byte) (Maps, error) {
	var maps Maps

	var width, height int
	var x, y int
	re := regexp.MustCompile(`(\d*)([ob$])`)

	for _, l := range data {
		line := string(l)
		if maps == nil {
			if strings.HasPrefix(line, "x =") {
				parts := strings.Split(line, ",")
				for _, part := range parts {
					part = strings.TrimSpace(part)
					if strings.HasPrefix(part, "x = ") {
						width, _ = strconv.Atoi(strings.TrimPrefix(part, "x = "))
					}
					if strings.HasPrefix(part, "y = ") {
						height, _ = strconv.Atoi(strings.TrimPrefix(part, "y = "))
					}
					if strings.HasPrefix(part, "rule = ") {
						rule := strings.TrimPrefix(part, "rule = ")
						if strings.ToLower(rule) != "b3/s23" {
							return nil, fmt.Errorf("unsupported rule %q", rule)
						}
					}
				}

				maps = NewMapsBits(width, height)
				continue
			}
		}

		if maps == nil {
			return nil, fmt.Errorf("invalid format line: %s", line)
		}

		for _, match := range re.FindAllStringSubmatch(line, -1) {
			count := 1
			if match[1] != "" {
				count, _ = strconv.Atoi(match[1])
			}
			switch match[2] {
			case "b":
				x += count
			case "o":
				for i := 0; i < count; i++ {
					if x < width && y < height {
						maps.Set(x, y, true)
					}
					x++
				}
			case "$":
				y += count
				x = 0
			}
		}
	}

	return maps, nil
}
