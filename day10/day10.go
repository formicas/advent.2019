package day10

import (
	"fmt"
	"math"
	"sort"
)

type Vector struct {
	x, y int
}

type asteroid struct {
	position Vector
	angle    float64
	distance int
}

type asteroidCollection struct {
	asteroids []asteroid
	length    int
}

//DoTheThing that day 10 should do
func DoTheThing() {
	lines := []string{
		"#.#.###.#.#....#..##.#....",
		".....#..#..#..#.#..#.....#",
		".##.##.##.##.##..#...#...#",
		"#.#...#.#####...###.#.#.#.",
		".#####.###.#.#.####.#####.",
		"#.#.#.##.#.##...####.#.##.",
		"##....###..#.#..#..#..###.",
		"..##....#.#...##.#.#...###",
		"#.....#.#######..##.##.#..",
		"#.###.#..###.#.#..##.....#",
		"##.#.#.##.#......#####..##",
		"#..##.#.##..###.##.###..##",
		"#..#.###...#.#...#..#.##.#",
		".#..#.#....###.#.#..##.#.#",
		"#.##.#####..###...#.###.##",
		"#...##..#..##.##.#.##..###",
		"#.#.###.###.....####.##..#",
		"######....#.##....###.#..#",
		"..##.#.####.....###..##.#.",
		"#..#..#...#.####..######..",
		"#####.##...#.#....#....#.#",
		".#####.##.#.#####..##.#...",
		"#..##..##.#.##.##.####..##",
		".##..####..#..####.#######",
		"#.#..#.##.#.######....##..",
		".#.##.##.####......#.##.##",
	}

	starMap := createStarMap(lines)
	optimalAsteroid, count := findOptimalAsteroid(starMap)
	fmt.Println(optimalAsteroid)
	fmt.Println(count)

	visibleAsteroidMap := getVisibleAsteroids(starMap, optimalAsteroid)
	visibleAsteroids := make([]asteroid, 0)

	for _, v := range visibleAsteroidMap {
		visibleAsteroids = append(visibleAsteroids, v)
	}

	c := asteroidCollection{asteroids: visibleAsteroids, length: count}

	sort.Sort(c)
	fmt.Println(c.asteroids[199])
}

func (c asteroidCollection) Len() int {
	return c.length
}

func (c asteroidCollection) Less(i, j int) bool {
	return c.asteroids[i].angle < c.asteroids[j].angle
}

func (c asteroidCollection) Swap(i, j int) {
	c.asteroids[i], c.asteroids[j] = c.asteroids[j], c.asteroids[i]
}

func createStarMap(lines []string) []Vector {
	asteroids := make([]Vector, 0)
	for i, line := range lines {
		runes := []rune(line)
		for j, char := range runes {
			if char == '#' {
				asteroids = append(asteroids, Vector{x: j, y: i})
			}
		}
	}

	return asteroids
}

//pick the best asteroid on the star map
func findOptimalAsteroid(asteroids []Vector) (Vector, int) {
	mostVisibleAsteroids := 0
	var optimalAsteroid Vector
	for _, a := range asteroids {
		visibleAsteroids := countVisibleAsteroids(asteroids, a)

		if visibleAsteroids > mostVisibleAsteroids {
			mostVisibleAsteroids = visibleAsteroids
			optimalAsteroid = a
		}
	}

	return optimalAsteroid, mostVisibleAsteroids
}

//count asteroids visible from x,y
func countVisibleAsteroids(asteroids []Vector, v Vector) int {
	failVectors := make(map[Vector]struct{})
	visibleAsteroids := 0

	for _, a := range asteroids {
		//skip itself
		if v.x == a.x && v.y == a.y {
			continue
		}

		diff := Vector{x: a.x - v.x, y: v.y - a.y}
		gcd := gcd(int(math.Abs(float64(diff.x))), int(math.Abs(float64(diff.y))))
		diff.x = diff.x / gcd
		diff.y = diff.y / gcd

		_, exists := failVectors[diff]
		if !exists {
			visibleAsteroids++
			failVectors[diff] = struct{}{}
		}
	}

	return visibleAsteroids
}

func getVisibleAsteroids(asteroids []Vector, v Vector) map[Vector]asteroid {
	visibleAsteroids := make(map[Vector]asteroid)

	for _, a := range asteroids {
		//skip itself
		if v.x == a.x && v.y == a.y {
			continue
		}

		diff := Vector{x: a.x - v.x, y: v.y - a.y}
		gcd := gcd(int(math.Abs(float64(diff.x))), int(math.Abs(float64(diff.y))))
		basicVector := Vector{x: diff.x / gcd, y: diff.y / gcd}
		distance := diff.x*diff.x + diff.y*diff.y

		visibleAsteroid, exists := visibleAsteroids[basicVector]
		if !exists || distance < visibleAsteroid.distance {
			visibleAsteroids[basicVector] = asteroid{position: a, angle: getAngle(basicVector), distance: distance}
		}
	}

	return visibleAsteroids
}

func gcd(a, b int) int {
	if a == b {
		return a
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	left := a
	right := b

	gcd := 1
	r := -1
	for r != 0 {
		left, right = swapify(left, right)
		gcd = right
		r = getRemainder(left, right)
		left = right
		right = r
	}
	return gcd
}

func getRemainder(left, right int) int {
	q := 0

	for left >= right {
		left -= right
		q++
	}
	return left
}

func swapify(left, right int) (int, int) {
	if left < right {
		return right, left
	}
	return left, right
}

//get the angle clockwise from north'
func getAngle(v Vector) float64 {
	intermediate := math.Pi/2 - math.Atan2(float64(v.y), float64(v.x))
	if intermediate >= 0 {
		return intermediate
	}
	return 2*math.Pi + intermediate

}
