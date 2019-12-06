package main

import "testing"

//TestMap1 ensure the most basic planet has length 1
func TestMap1(t *testing.T) {
	orbitMap := []string{"COM)FOO"}

	planetMap := buildChart(&orbitMap)
	got := (*planetMap)["FOO"].TotalOrbits(planetMap)
	if got != 1 {
		t.Errorf("TotalOrbits = %d; want 1", got)
	}
}

//TestMap2 ensure the most basic planet has length 1
func TestMap2(t *testing.T) {
	orbitMap := []string{"COM)FOO", "FOO)BAR"}

	planetMap := buildChart(&orbitMap)
	got := (*planetMap)["BAR"].TotalOrbits(planetMap)
	if got != 2 {
		t.Errorf("TotalOrbits = %d; want 2", got)
	}
}
