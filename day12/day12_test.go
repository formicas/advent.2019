package day12

import "testing"

func Test_simulateMoons(t *testing.T) {
	tests := []struct {
		moons, expected []moon
		steps           int
	}{
		{
			moons: []moon{
				moon{p: vector3{x: -1, y: 0, z: 2}, name: "Io"},
				moon{p: vector3{x: 2, y: -10, z: -7}, name: "Europa"},
				moon{p: vector3{x: 4, y: -8, z: 8}, name: "Ganymede"},
				moon{p: vector3{x: 3, y: 5, z: -1}, name: "Callisto"},
			},
			expected: []moon{
				moon{p: vector3{x: 2, y: 1, z: -3}, v: vector3{x: -3, y: -2, z: 1}, name: "Io"},
				moon{p: vector3{x: 1, y: -8, z: 0}, v: vector3{x: -1, y: 1, z: 3}, name: "Europa"},
				moon{p: vector3{x: 3, y: -6, z: 1}, v: vector3{x: 3, y: 2, z: -3}, name: "Ganymede"},
				moon{p: vector3{x: 2, y: 0, z: 4}, v: vector3{x: 1, y: -1, z: -1}, name: "Callisto"},
			},
			steps: 10,
		},
		{
			moons: []moon{
				moon{p: vector3{x: -8, y: -10, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: 5, z: 10}, name: "Europa"},
				moon{p: vector3{x: 2, y: -7, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: 9, y: -8, z: -3}, name: "Callisto"},
			},
			expected: []moon{
				moon{p: vector3{x: -9, y: -10, z: 1}, v: vector3{x: -2, y: -2, z: -1}, name: "Io"},
				moon{p: vector3{x: 4, y: 10, z: 9}, v: vector3{x: -3, y: 7, z: -2}, name: "Europa"},
				moon{p: vector3{x: 8, y: -10, z: -3}, v: vector3{x: 5, y: -1, z: -2}, name: "Ganymede"},
				moon{p: vector3{x: 5, y: -10, z: 3}, v: vector3{x: 0, y: -4, z: 5}, name: "Callisto"},
			},
			steps: 10,
		},
		{
			moons: []moon{
				moon{p: vector3{x: -8, y: -10, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: 5, z: 10}, name: "Europa"},
				moon{p: vector3{x: 2, y: -7, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: 9, y: -8, z: -3}, name: "Callisto"},
			},
			expected: []moon{
				moon{p: vector3{x: -10, y: 3, z: -4}, v: vector3{x: -5, y: 2, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: -25, z: 6}, v: vector3{x: 1, y: 1, z: -4}, name: "Europa"},
				moon{p: vector3{x: 13, y: 1, z: 1}, v: vector3{x: 5, y: -2, z: 2}, name: "Ganymede"},
				moon{p: vector3{x: 0, y: 1, z: 7}, v: vector3{x: -1, y: -1, z: 2}, name: "Callisto"},
			},
			steps: 20,
		},
		{
			moons: []moon{
				moon{p: vector3{x: -8, y: -10, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: 5, z: 10}, name: "Europa"},
				moon{p: vector3{x: 2, y: -7, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: 9, y: -8, z: -3}, name: "Callisto"},
			},
			expected: []moon{
				moon{p: vector3{x: 15, y: -6, z: -9}, v: vector3{x: -5, y: 4, z: 0}, name: "Io"},
				moon{p: vector3{x: -4, y: -11, z: 3}, v: vector3{x: -3, y: -10, z: 0}, name: "Europa"},
				moon{p: vector3{x: 0, y: -1, z: 11}, v: vector3{x: 7, y: 4, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: -3, y: -2, z: 5}, v: vector3{x: 1, y: 2, z: -3}, name: "Callisto"},
			},
			steps: 30,
		},
		{
			moons: []moon{
				moon{p: vector3{x: -8, y: -10, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: 5, z: 10}, name: "Europa"},
				moon{p: vector3{x: 2, y: -7, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: 9, y: -8, z: -3}, name: "Callisto"},
			},
			expected: []moon{
				moon{p: vector3{x: 14, y: -12, z: -4}, v: vector3{x: 11, y: 3, z: 0}, name: "Io"},
				moon{p: vector3{x: -1, y: 18, z: 8}, v: vector3{x: -5, y: 2, z: 3}, name: "Europa"},
				moon{p: vector3{x: -5, y: -14, z: 8}, v: vector3{x: 1, y: -2, z: 0}, name: "Ganymede"},
				moon{p: vector3{x: 0, y: -12, z: -2}, v: vector3{x: -7, y: -3, z: -3}, name: "Callisto"},
			},
			steps: 40,
		},
		{
			moons: []moon{
				moon{p: vector3{x: -8, y: -10, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: 5, z: 10}, name: "Europa"},
				moon{p: vector3{x: 2, y: -7, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: 9, y: -8, z: -3}, name: "Callisto"},
			},
			expected: []moon{
				moon{p: vector3{x: -23, y: 4, z: 1}, v: vector3{x: -7, y: -1, z: 2}, name: "Io"},
				moon{p: vector3{x: 20, y: -31, z: 13}, v: vector3{x: 5, y: 3, z: 4}, name: "Europa"},
				moon{p: vector3{x: -4, y: 6, z: 1}, v: vector3{x: -1, y: 1, z: -3}, name: "Ganymede"},
				moon{p: vector3{x: 15, y: 1, z: -5}, v: vector3{x: 3, y: -3, z: -3}, name: "Callisto"},
			},
			steps: 50,
		},
		{
			moons: []moon{
				moon{p: vector3{x: -8, y: -10, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: 5, z: 10}, name: "Europa"},
				moon{p: vector3{x: 2, y: -7, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: 9, y: -8, z: -3}, name: "Callisto"},
			},
			expected: []moon{
				moon{p: vector3{x: 8, y: -12, z: -9}, v: vector3{x: -7, y: 3, z: 0}, name: "Io"},
				moon{p: vector3{x: 13, y: 16, z: -3}, v: vector3{x: 3, y: -11, z: -5}, name: "Europa"},
				moon{p: vector3{x: -29, y: -11, z: -1}, v: vector3{x: -3, y: 7, z: 4}, name: "Ganymede"},
				moon{p: vector3{x: 16, y: -13, z: 23}, v: vector3{x: 7, y: 1, z: 1}, name: "Callisto"},
			},
			steps: 100,
		},
	}

	for j, test := range tests {
		moons := test.moons
		expected := test.expected
		for i, moon := range simulateMoons(moons, test.steps) {
			if moon != expected[i] {
				t.Errorf("Test %d: Expected %v, got %v", j, expected[i], moon)
			}
		}

	}

}

func Test_getTotalEnergy(t *testing.T) {
	tests := []struct {
		moons    []moon
		steps    int
		expected int
	}{
		{
			moons: []moon{
				moon{p: vector3{x: -1, y: 0, z: 2}, name: "Io"},
				moon{p: vector3{x: 2, y: -10, z: -7}, name: "Europa"},
				moon{p: vector3{x: 4, y: -8, z: 8}, name: "Ganymede"},
				moon{p: vector3{x: 3, y: 5, z: -1}, name: "Callisto"},
			},
			steps:    10,
			expected: 179,
		},
		{
			moons: []moon{
				moon{p: vector3{x: -8, y: -10, z: 0}, name: "Io"},
				moon{p: vector3{x: 5, y: 5, z: 10}, name: "Europa"},
				moon{p: vector3{x: 2, y: -7, z: 3}, name: "Ganymede"},
				moon{p: vector3{x: 9, y: -8, z: -3}, name: "Callisto"},
			},
			steps:    100,
			expected: 1940,
		},
	}

	for i, test := range tests {
		expected := test.expected
		moons := test.moons

		got := getTotalEnergy(simulateMoons(moons, test.steps))

		if got != expected {
			t.Errorf("Test %d: Expected %d, got %d", i, expected, got)
		}
	}

}
