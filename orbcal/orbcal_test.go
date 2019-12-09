package orbcal

import "testing"

//TestMap1 ensure the most basic planet has length 1
func TestMap1(t *testing.T) {
	orbitMap := []string{"COM)FOO"}

	planetMap := buildChart(orbitMap)
	got := (*planetMap)["FOO"].getTotalOrbits(planetMap)
	if got != 1 {
		t.Errorf("TotalOrbits = %d; want 1", got)
	}
}

//TestMap2 ensure the most basic planet has length 1
func TestMap2(t *testing.T) {
	orbitMap := []string{"COM)FOO", "FOO)BAR"}

	planetMap := buildChart(orbitMap)
	got := (*planetMap)["BAR"].getTotalOrbits(planetMap)
	if got != 2 {
		t.Errorf("TotalOrbits = %d; want 2", got)
	}
}

func TestShortestPath(t *testing.T) {
	orbitMap := []string{"COM)FOO", "COM)BAR"}

	planetMap := buildChart(orbitMap)

	got := shortestPath("FOO", "BAR", planetMap)

	if got != 0 {
		t.Errorf("Shortest path = %d; want 0", got)
	}
}

func TestShortestPath2(t *testing.T) {
	orbitMap := []string{"COM)FOO", "FOO)BAR", "FOO)BUZZ"}

	planetMap := buildChart(orbitMap)

	got := shortestPath("BUZZ", "BAR", planetMap)

	if got != 0 {
		t.Errorf("Shortest path = %d; want 0", got)
	}
}

func TestShortestPath3(t *testing.T) {
	orbitMap := []string{"COM)FOO", "FOO)BAR", "COM)BUZZ"}

	planetMap := buildChart(orbitMap)

	got := shortestPath("BUZZ", "BAR", planetMap)

	if got != 1 {
		t.Errorf("Shortest path = %d; want 1", got)
	}
}

func TestShortestPath4(t *testing.T) {
	orbitMap := []string{"COM)FOO", "FOO)BAR", "FOO)BUZZ", "BAR)BIZZ", "BUZZ)BAZZ"}
	//             BUZZ - BAZZ
	//           /
	// COM - FOO - BAR - BIZZ
	planetMap := buildChart(orbitMap)

	got := shortestPath("BIZZ", "BAZZ", planetMap)

	if got != 2 {
		t.Errorf("Shortest path = %d; want 2", got)
	}
}
