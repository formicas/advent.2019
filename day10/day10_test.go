package day10

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func areEqual(lefts, rights []Vector) bool {
	if len(lefts) != len(rights) {
		return false
	}

	for i, left := range lefts {
		if left.x != rights[i].x || left.y != rights[i].y {
			return false
		}
	}
	return true

}
func Test_gcd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 5, 5},
		{2, 3, 1},
		{2, 4, 2},
		{1071, 462, 21},
	}

	for _, test := range tests {
		got := gcd(test.a, test.b)
		fmt.Printf("gcd(%d,%d) = %d\n", test.a, test.b, got)
		if got != test.expected {
			t.Errorf("expected %d", test.expected)
		}
	}
}

func Test_createStarMap(t *testing.T) {
	tests := []struct {
		lines    []string
		expected []Vector
	}{
		{[]string{".#", ".."}, []Vector{{1, 0}}},
		{[]string{".#", "#."}, []Vector{{1, 0}, {0, 1}}},
		{[]string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##"}, []Vector{{1, 0}, {4, 0}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {4, 3}, {3, 4}, {4, 4}}},
	}

	for _, test := range tests {
		asteroids := createStarMap(test.lines)
		fmt.Println(asteroids)
		if !areEqual(asteroids, test.expected) {
			t.Errorf("expected %v", test.expected)
		}
	}
}

func Test_countVisibleAsteroids(t *testing.T) {
	tests := []struct {
		lines    []string
		asteroid Vector
		expected int
	}{{
		[]string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##",
		}, Vector{3, 4}, 8},
	}

	for _, test := range tests {
		visible := countVisibleAsteroids(createStarMap(test.lines), test.asteroid)
		fmt.Println(test.lines)
		fmt.Println(test.asteroid)
		fmt.Println(visible)
		if visible != test.expected {
			t.Errorf("expected %d", test.expected)
		}
	}

}

func Test_findOptimalAsteroid(t *testing.T) {
	tests := []struct {
		lines            []string
		expectedAsteroid Vector
		expectedTotal    int
	}{{[]string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}, Vector{3, 4}, 8},
		{[]string{
			"......#.#.",
			"#..#.#....",
			"..#######.",
			".#.#.###..",
			".#..#.....",
			"..#....#.#",
			"#..#....#.",
			".##.#..###",
			"##...#..#.",
			".#....####",
		}, Vector{5, 8}, 33},
		{[]string{
			"#.#...#.#.",
			".###....#.",
			".#....#...",
			"##.#.#.#.#",
			"....#.#.#.",
			".##..###.#",
			"..#...##..",
			"..##....##",
			"......#...",
			".####.###.",
		}, Vector{1, 2}, 35},
		{[]string{
			".#..##.###...#######",
			"##.############..##.",
			".#.######.########.#",
			".###.#######.####.#.",
			"#####.##.#.##.###.##",
			"..#####..#.#########",
			"####################",
			"#.####....###.#.#.##",
			"##.#################",
			"#####.##.###..####..",
			"..######..##.#######",
			"####.##.####...##..#",
			".#####..#.######.###",
			"##...#.##########...",
			"#.##########.#######",
			".####.#.###.###.#.##",
			"....##.##.###..#####",
			".#.#.###########.###",
			"#.#.#.#####.####.###",
			"###.##.####.##.#..##",
		}, Vector{11, 13}, 210},
	}

	for _, test := range tests {
		asteroid, total := findOptimalAsteroid(createStarMap(test.lines))
		fmt.Println(asteroid)
		if asteroid.x != test.expectedAsteroid.x || asteroid.y != test.expectedAsteroid.y {
			t.Errorf("expected %v", test.expectedAsteroid)
		}
		fmt.Println(total)
		if total != test.expectedTotal {
			t.Errorf("expected %d", test.expectedTotal)
		}
	}
}

func Test_getAngle(t *testing.T) {
	tests := []struct {
		v        Vector
		expected float64
	}{
		{Vector{x: 0, y: 1}, 0},
		{Vector{x: 1, y: 1}, math.Pi / 4},
		{Vector{x: 1, y: 0}, math.Pi / 2},
		{Vector{x: 1, y: -1}, 3 * math.Pi / 4},
		{Vector{x: 0, y: -1}, math.Pi},
		{Vector{x: -1, y: -1}, 5 * math.Pi / 4},
		{Vector{x: -1, y: 0}, 3 * math.Pi / 2},
		{Vector{x: -1, y: 1}, 7 * math.Pi / 4},
	}

	for _, test := range tests {
		angle := getAngle(test.v)
		if angle != test.expected {
			t.Errorf("Vector %v Got %v, expected %v", test.v, angle, test.expected)
		}
	}
}

func Test_Get200th(t *testing.T) {
	lines := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}

	starMap := createStarMap(lines)
	optimalAsteroid, count := findOptimalAsteroid(starMap)

	visibleAsteroidMap := getVisibleAsteroids(starMap, optimalAsteroid)
	visibleAsteroids := make([]asteroid, 0)

	for _, v := range visibleAsteroidMap {
		visibleAsteroids = append(visibleAsteroids, v)
	}

	c := asteroidCollection{asteroids: visibleAsteroids, length: count}

	sort.Sort(c)

	if c.asteroids[199].position.x != 8 || c.asteroids[199].position.y != 2 {
		t.Errorf("Got {%d,%d}, expected {%d,%d}", c.asteroids[199].position.x, c.asteroids[199].position.y, 8, 2)
	}

}
