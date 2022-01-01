package main

import (
	"adventofcode_2015/dinner"
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	in := input
	scanner := bufio.NewScanner(bytes.NewBufferString(in))
	locationSet := map[string]bool{}
	var locations []string
	distances := map[string]int{}
	for scanner.Scan() {
		s := strings.Trim(scanner.Text(), " \r\n\t")
		if s == "" {
			continue
		}
		var from, to string
		var dist int
		_, err := fmt.Sscanf(s, "%s to %s = %d", &from, &to, &dist)
		if err != nil {
			panic(err)
		}
		if _, ok := locationSet[from]; !ok {
			locationSet[from] = true
			locations = append(locations, from)
		}
		if _, ok := locationSet[to]; !ok {
			locationSet[to] = true
			locations = append(locations, to)
		}
		distances[from+":"+to] = dist
		distances[to+":"+from] = dist
	}
	fmt.Printf("locations: %v\n", locations)

	minDist := -1
	var minRoute []string
	maxDist := -1
	var maxRoute []string
	for perm := range dinner.Permutations(locations) {
		var permDist int
		for i := 0; i < len(perm)-1; i++ {
			d := distances[perm[i]+":"+perm[i+1]]
			permDist += d
		}

		if minDist < 0 || permDist < minDist {
			minDist = permDist
			minRoute = perm
		}
		if maxDist < 0 || permDist > maxDist {
			maxDist = permDist
			maxRoute = perm
		}
	}
	fmt.Printf("min-route: %v of length %d\n", minRoute, minDist)
	fmt.Printf("max-route: %v of length %d\n", maxRoute, maxDist)
}

var inputTest = `
London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
`

var input = `
Faerun to Norrath = 129
Faerun to Tristram = 58
Faerun to AlphaCentauri = 13
Faerun to Arbre = 24
Faerun to Snowdin = 60
Faerun to Tambi = 71
Faerun to Straylight = 67
Norrath to Tristram = 142
Norrath to AlphaCentauri = 15
Norrath to Arbre = 135
Norrath to Snowdin = 75
Norrath to Tambi = 82
Norrath to Straylight = 54
Tristram to AlphaCentauri = 118
Tristram to Arbre = 122
Tristram to Snowdin = 103
Tristram to Tambi = 49
Tristram to Straylight = 97
AlphaCentauri to Arbre = 116
AlphaCentauri to Snowdin = 12
AlphaCentauri to Tambi = 18
AlphaCentauri to Straylight = 91
Arbre to Snowdin = 129
Arbre to Tambi = 53
Arbre to Straylight = 40
Snowdin to Tambi = 15
Snowdin to Straylight = 99
Tambi to Straylight = 70
`
