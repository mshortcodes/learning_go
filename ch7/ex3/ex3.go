package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func (l League) MatchResult(team1, team2 string, score1, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}
	if _, ok := l.Teams[team2]; !ok {
		return
	}
	if score1 == score2 {
		return
	}
	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for name := range l.Teams {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	vals := r.Ranking()
	for _, v := range vals {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func main() {
	league := League{
		Teams: map[string]Team{
			"France": {
				Name:    "France",
				Players: []string{"p1, p2, p3, p4, p5"},
			},
			"Italy": {
				Name:    "Italy",
				Players: []string{"p1, p2, p3, p4, p5"},
			},
			"USA": {
				Name:    "USA",
				Players: []string{"p1, p2, p3, p4, p5"},
			},
		},
		Wins: map[string]int{},
	}

	league.MatchResult("France", "Italy", 80, 20)
	league.MatchResult("France", "USA", 70, 30)
	league.MatchResult("USA", "Italy", 100, 0)
	league.MatchResult("Italy", "France", 50, 50)
	fmt.Println(league.Ranking())
	RankPrinter(league, os.Stdout)
}
