package day14

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type reaction struct {
	inputs []chemical
	output chemical
	rank   int
}

type chemical struct {
	name  string
	count int
}

//DoTheThing that day 14 is supposed to do
func DoTheThing() {
	reactionsDescription := `13 WDSR, 16 FXQB => 6 BSTCB
185 ORE => 9 BWSCM
1 WDSR => 9 RLFSK
5 LCGL, 7 BWSCM => 9 BSVW
6 NLSL => 3 MJSQ
1 JFGM, 7 BSVW, 7 XRLN => 6 WDSR
3 WZLFV => 3 BZDPT
5 DTHZH, 12 QNTH, 20 BSTCB => 4 BMXF
18 JSJWJ, 6 JLMD, 6 TMTF, 3 XSNL, 3 BWSCM, 83 LQTJ, 29 KDGNL => 1 FUEL
1 LWPD, 28 RTML, 16 FDPM, 8 JSJWJ, 2 TNMTC, 20 DTHZH => 9 JLMD
1 SDVXW => 6 BPTV
180 ORE => 7 JFGM
13 RLFSK, 15 HRKD, 1 RFQWL => 5 QNTH
1 RFQWL, 3 NZHFV, 18 XRLN => 9 HRKD
2 NLSL, 2 JXVZ => 5 GTSJ
19 SDVXW, 2 BSVW, 19 XRLN => 6 QMFV
1 CSKP => 8 LQTJ
4 ZSZBN => 5 RBRZT
8 WZLFV, 3 QNWRZ, 1 DTHZH => 4 RTRN
1 CGXBG, 1 PGXFJ => 3 TNMTC
4 CGCSL => 7 RNFW
9 CGCSL, 1 HGTL, 3 BHJXV => 8 RSVR
5 NGJW => 8 HTDM
21 FPBTN, 1 TNMTC, 2 RBRZT, 8 BDHJ, 28 WXQX, 9 RNFW, 6 RSVR => 1 XSNL
2 WZLFV => 5 BHJXV
10 BSTCB, 4 NLSL => 4 HQLHN
1 JFGM => 7 SDVXW
6 CSKP => 8 FXQB
6 TNMTC, 4 BZDPT, 1 BPTV, 18 JSJWJ, 2 DTHZH, 1 LWPD, 8 RTML => 8 KDGNL
6 XFGWZ => 7 CGCSL
3 GTSJ => 4 LWPD
1 WDSR, 1 QNWRZ => 5 XFGWZ
11 CSKP, 10 SDVXW => 4 QNWRZ
7 BSVW, 4 QMFV => 1 RFQWL
12 QNTH, 10 HTDM, 3 WXQX => 3 FDPM
2 HGTL => 7 PGXFJ
14 SDVXW => 6 CSKP
11 HQLHN, 1 GTSJ, 1 QNTH => 5 TMTF
173 ORE => 9 LCGL
4 WXQX => 9 BDHJ
5 BZDPT => 7 NGJW
1 GTSJ, 23 QNWRZ, 6 LQTJ => 7 JSJWJ
23 NZHFV, 3 HQLHN => 6 DTHZH
2 JFGM => 4 XRLN
20 CGCSL => 9 WXQX
2 BSTCB, 3 HRKD => 9 NLSL
1 MJSQ, 1 BPTV => 8 CGXBG
1 RTRN, 1 RSVR => 3 ZSZBN
2 NZHFV, 1 BSTCB, 20 HRKD => 1 JXVZ
2 BZDPT => 5 HGTL
1 ZSZBN, 14 FDPM => 9 RTML
3 BMXF => 8 FPBTN
1 SDVXW, 8 XRLN => 9 NZHFV
18 QNWRZ, 7 RLFSK => 1 WZLFV`

	reactions := parseReactions(reactionsDescription)
	rankReactions("FUEL", "ORE", &reactions)
	fmt.Printf("%d Ore for 1 Fuel\n", getRequiredOre(reactions, 1))

	//I think this might be cheating for part 2, but...uh whatever!
	guess := guessFuelFor1Trillion(reactions)
	fmt.Printf("Guess %d Fuel from %d Ore\n", guess, 1000000000000)
	//if we go over, drop the guess
	for getRequiredOre(reactions, guess) > 1000000000000 {
		guess--
	}
	//or if we guessed under somehow, incease it
	for getRequiredOre(reactions, guess+1) < 1000000000000 {
		guess++
	}
	fmt.Printf("But it's actually %d", guess)
}

func parseReactions(reactionsDescription string) map[string]reaction {
	reactions := make(map[string]reaction)
	regex := regexp.MustCompile(`(\d+) (\w+)`)
	for _, line := range strings.Split(reactionsDescription, "\n") {
		reaction := getReaction(regex.FindAllStringSubmatch(line, -1))
		reactions[reaction.output.name] = reaction
	}
	return reactions
}

func getReaction(matches [][]string) reaction {
	inputCount := len(matches) - 1
	chemicals := make([]chemical, inputCount)
	for i, chem := range matches {
		count, _ := strconv.Atoi(chem[1])
		chemical := chemical{name: chem[2], count: count}
		if i < inputCount {
			chemicals[i] = chemical
			continue
		}
		return reaction{inputs: chemicals, output: chemical}
	}
	return reaction{}
}

func rankReactions(outputName, baseInputName string, reactions *map[string]reaction) int {
	reaction := (*reactions)[outputName]

	highestRankedInput := 0
	for _, input := range reaction.inputs {
		if input.name != baseInputName {
			highestRankedInput = int(math.Max(float64(rankReactions(input.name, baseInputName, reactions)), float64(highestRankedInput)))
		}
	}
	reaction.rank = highestRankedInput + 1
	(*reactions)[outputName] = reaction
	return reaction.rank
}

func getRequiredOre(reactions map[string]reaction, requiredFuel int) int {
	rankedReactions := make([]reaction, 0)
	for _, reaction := range reactions {
		rankedReactions = append(rankedReactions, reaction)
	}

	//invert the less function to sort descending
	sort.Slice(rankedReactions, func(i, j int) bool { return rankedReactions[i].rank > rankedReactions[j].rank })

	requiredChemicals := map[string]int{
		"FUEL": requiredFuel,
	}
	for _, reaction := range rankedReactions {
		output := reaction.output
		requiredReactions := int(math.Ceil(float64(requiredChemicals[output.name]) / float64(output.count)))
		for _, input := range reaction.inputs {
			_, exists := requiredChemicals[input.name]
			if exists {
				requiredChemicals[input.name] += requiredReactions * input.count
			} else {
				requiredChemicals[input.name] = requiredReactions * input.count
			}
		}
	}

	return requiredChemicals["ORE"]
}

func guessFuelFor1Trillion(reactions map[string]reaction) int {
	rankedReactions := make([]reaction, 0)
	for _, reaction := range reactions {
		rankedReactions = append(rankedReactions, reaction)
	}

	//invert the less function to sort descending
	sort.Slice(rankedReactions, func(i, j int) bool { return rankedReactions[i].rank > rankedReactions[j].rank })

	requiredChemicals := map[string]float64{
		"FUEL": 1,
	}
	for _, reaction := range rankedReactions {
		output := reaction.output
		requiredReactions := float64(requiredChemicals[output.name]) / float64(output.count)
		for _, input := range reaction.inputs {
			_, exists := requiredChemicals[input.name]
			if exists {
				requiredChemicals[input.name] += requiredReactions * float64(input.count)
			} else {
				requiredChemicals[input.name] = requiredReactions * float64(input.count)
			}
		}
	}

	return int(1000000000000 / requiredChemicals["ORE"])
}
