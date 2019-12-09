package orbcal

import (
	"math"
	"strings"
)

type planet struct {
	Name, ParentName  string
	totalOrbits       int
	orbitsCalculated  bool
	parents           []string
	parentsCalculated bool
}

func (p planet) getTotalOrbits(planetChart *map[string]planet) int {
	if p.orbitsCalculated {
		return p.totalOrbits
	}

	if p.ParentName == "COM" {
		p.totalOrbits = 1
		p.orbitsCalculated = true
	} else {
		parent := (*planetChart)[p.ParentName]
		p.totalOrbits = parent.getTotalOrbits(planetChart) + 1
		p.orbitsCalculated = true
	}
	return p.totalOrbits
}

func (p planet) getParents(planetChart *map[string]planet) []string {
	if p.parentsCalculated {
		return p.parents
	}
	if p.ParentName == "COM" {
		p.parentsCalculated = true
		p.parents = []string{"COM"}
		return p.parents
	}

	parent := (*planetChart)[p.ParentName]
	parents := parent.getParents(planetChart)
	p.parents = append(parents, p.ParentName)
	p.parentsCalculated = true
	return p.parents
}

func BuildChart(definition []string) *map[string]planet {
	planetMap := make(map[string]planet)

	for _, relationship := range definition {
		split := strings.Split(relationship, ")")
		parent := split[0]
		name := split[1]
		planet := planet{Name: name, ParentName: parent}
		planetMap[name] = planet
	}
	return &planetMap
}

func CheckSum(planetMap *map[string]planet) int {
	sum := 0
	for _, planet := range *planetMap {
		sum += planet.getTotalOrbits(planetMap)
	}

	return sum
}

func ShortestPath(a, b string, planetMap *map[string]planet) int {
	var commonParent planet

	planetA := (*planetMap)[a]
	planetB := (*planetMap)[b]
	//by putting the name on the end, we guarantee we'll trigger the interesting case in the for loop
	aParents := append(planetA.getParents(planetMap), planetA.Name)
	bParents := append(planetB.getParents(planetMap), planetB.Name)
	commonParentIsCom := false

	//todo handle the case where one is a descendant
	for i := 0; i < int(math.Min(float64(len(aParents)), float64(len(bParents)))); i++ {
		if aParents[i] != bParents[i] {
			commonParentName := aParents[i-1]
			if commonParentName == "COM" {
				commonParentIsCom = true
			} else {
				commonParent = (*planetMap)[aParents[i-1]]
				//fmt.Printf("Parent: %v\n", commonParent.Name)
			}

			break
		}
	}

	if commonParentIsCom { //i dunno
		return planetA.getTotalOrbits(planetMap) + planetB.getTotalOrbits(planetMap) - 2
	}
	//fmt.Printf("%d %d %d", planetA.getTotalOrbits(planetMap), planetB.getTotalOrbits(planetMap), commonParent.getTotalOrbits(planetMap))
	return planetA.getTotalOrbits(planetMap) + planetB.getTotalOrbits(planetMap) - 2*commonParent.getTotalOrbits(planetMap) - 2
}
