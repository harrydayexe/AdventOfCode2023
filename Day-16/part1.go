package main

func part1(lines []string) int {
	var seen = make(map[[4]int]struct{})
	var queue = [][4]int{{0, -1, 0, 1}}

	for len(queue) > 0 {
		r, c, dr, dc := queue[0][0], queue[0][1], queue[0][2], queue[0][3]
		queue = queue[1:]
		r += dr
		c += dc

		if r < 0 || r >= len(lines) || c < 0 || c >= len(lines[0]) {
			continue
		}

		ch := rune(lines[r][c])
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
