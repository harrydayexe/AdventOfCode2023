package main

func part2(lines []string) int {
	var maxVal int

	for i := 0; i < len(lines); i++ {
		maxVal = max(maxVal, calcNum(i, -1, 0, 1, lines))
		maxVal = max(maxVal, calcNum(i, len(lines[0]), 0, -1, lines))
	}

	for i := 0; i < len(lines[0]); i++ {
		maxVal = max(maxVal, calcNum(-1, i, 1, 0, lines))
		maxVal = max(maxVal, calcNum(len(lines), i, -1, 0, lines))
	}

	return maxVal
}

func calcNum(ir, ic, idr, idc int, grid []string) int {
	var seen = make(map[[4]int]struct{})
	var queue = [][4]int{{ir, ic, idr, idc}}

	for len(queue) > 0 {
		r, c, dr, dc := queue[0][0], queue[0][1], queue[0][2], queue[0][3]
		queue = queue[1:]
		r += dr
		c += dc

		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}

		ch := rune(grid[r][c])
		if ch == '.' || (ch == '-' && dc != 0) || (ch == '|' && dr != 0) {
			_, prs := seen[[4]int{r, c, dr, dc}]
			if !prs {
				seen[[4]int{r, c, dr, dc}] = struct{}{}
				queue = append(queue, [4]int{r, c, dr, dc})
			}
		} else if ch == '/' {
			dr, dc = -dc, -dr
			_, prs := seen[[4]int{r, c, dr, dc}]
			if !prs {
				seen[[4]int{r, c, dr, dc}] = struct{}{}
				queue = append(queue, [4]int{r, c, dr, dc})
			}
		} else if ch == '\\' {
			dr, dc = dc, dr
			_, prs := seen[[4]int{r, c, dr, dc}]
			if !prs {
				seen[[4]int{r, c, dr, dc}] = struct{}{}
				queue = append(queue, [4]int{r, c, dr, dc})
			}
		} else {
			var deltas [2][2]int
			if ch == '|' {
				deltas = [2][2]int{{-1, 0}, {1, 0}}
			} else {
				deltas = [2][2]int{{0, -1}, {0, 1}}
			}

			for _, delta := range deltas {
				_, prs := seen[[4]int{r, c, delta[0], delta[1]}]
				if !prs {
					seen[[4]int{r, c, delta[0], delta[1]}] = struct{}{}
					queue = append(queue, [4]int{r, c, delta[0], delta[1]})
				}
			}
		}
	}

	var seenCoords = make(map[[2]int]struct{})
	for coords, _ := range seen {
		seenCoords[[2]int{coords[0], coords[1]}] = struct{}{}
	}
	return len(seenCoords)
}
