package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sort"
)

type result byte
type team byte

type match struct {
	first team
	second team
	result result
}

type tournament []match

type rating map[team]int

const (
	win result = 'W'
	draw result = 'D'
	loss result = 'L'
)

func (t *tournament) calcRating() rating {
	resultMatch := map[team]int{}
	for _, v := range *t {
		resultMatch[v.first] += addPointsWin(v.result)
		resultMatch[v.second] += addPointsLoss(v.result)
	}
	return resultMatch
}

func addPointsWin(v result) int {
	switch v {
	case win:
		return 3
	case draw:
		return 1
	case loss:
		return 0
	default:
		return 0
	}
}
func addPointsLoss(v result) int {
	switch v {
	case win:
		return 0
	case draw:
		return 1
	case loss:
		return 3
	default:
		return 0
	}
}

	// Sample Input:
	// ABW DCD DAW CBL BDL ACW

func startFootballTournament() {
	src := readString()
	trn := parseTournament(src)
	rt := trn.calcRating()
	rt.print()
}

func readString() string {
	fmt.Println("Enter data")
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ")
	trn := tournament{}
	for _, descr := range descriptions {
		m := parseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}