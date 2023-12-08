package main

type GraphMaps struct {
	left  map[string]string
	right map[string]string
}

func parseInput(lines []string) GraphMaps {
	var left, right = make(map[string]string), make(map[string]string)
	for _, line := range lines[2:] {
		id := line[:3]
		l := line[7:10]
		r := line[12:15]
		left[id] = l
		right[id] = r
	}

	return GraphMaps{
		left:  left,
		right: right,
	}
}
