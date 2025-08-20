package pathcrossing

/*
	Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively.
	You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.
	Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. 
	Return false otherwise.

	
	Example 1:

	Input: path = "NES"
	Output: false 
	Explanation: Notice that the path doesn't cross any point more than once.


	Example 2:

	Input: path = "NESWW"
	Output: true
	Explanation: Notice that the path visits the origin twice.
*/

func Solution(path string) bool {
	type coordinate struct {
		x, y int
	}

	var x, y int

	m := map[coordinate]struct{}{
		{x: x, y: y}: {},
	}

	for _, direction := range path {
		switch direction {
		case 'N':
			y++
		case 'S':
			y--
		case 'W':
			x--
		case 'E':
			x++
		}

		c := coordinate{x: x, y: y}

		if _, ok := m[c]; ok {
			return true
		}

		m[c] = struct{}{}
	}

	return false
}
