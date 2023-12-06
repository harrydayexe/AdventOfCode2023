package main

import (
	"sort"
)

func part2(lines []string) int {
	var p = cleanData2(lines)

	for _, block := range p.blocks {
		var news [][2]int
		for len(p.seeds) > 0 {
			var start, end int
			start, end, p.seeds = p.seeds[len(p.seeds)-1][0], p.seeds[len(p.seeds)-1][1], p.seeds[:len(p.seeds)-1]
			var found = false
			for _, r := range block {
				overlapStart := max(start, r[1])
				overlapEnd := min(end, r[1]+r[2])
				if overlapStart < overlapEnd {
					news = append(news, [2]int{overlapStart - r[1] + r[0], overlapEnd - r[1] + r[0]})
					if overlapStart > start {
						p.seeds = append(p.seeds, [2]int{start, overlapStart})
					}
					if end > overlapEnd {
						p.seeds = append(p.seeds, [2]int{overlapEnd, end})
					}
					found = true
					break
				}
			}
			if !found {
				news = append(news, [2]int{start, end})
			}
		}
		p.seeds = news
	}

	sort.Slice(p.seeds, func(i, j int) bool {
		return p.seeds[i][0] < p.seeds[j][0]
	})
	//fmt.Println(p.seeds)

	return p.seeds[0][0]
}
