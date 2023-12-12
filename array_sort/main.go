package main

import "sort"

type Fighter struct {
	Name   string
	Weight int
	Height int
}

func main() {
	fighters := []*Fighter{
		{"Anderson Silva", 185, 188},
		{"Jon Jones", 205, 193},
		{"Georges St-Pierre", 170, 178},
		{"Jose Aldo", 145, 170},
		{"Demetrious Johnson", 125, 160},
		{"Cain Velasquez", 265, 185},
		{"Junior Dos Santos", 265, 193},
		{"Daniel Cormier", 205, 180},
		{"Fabricio Werdum", 265, 193},
	}

	sort.Slice(fighters, func(i, j int) bool {
		if fighters[i].Weight < fighters[j].Weight {
			return true
		}

		if fighters[i].Height < fighters[j].Height {
			return true
		}
		return false
	})

	sort.Slice(fighters, func(i, j int) bool {
		if fighters[i].Weight == fighters[j].Weight {
			return fighters[i].Height < fighters[j].Height
		}
		return fighters[i].Weight < fighters[j].Weight
	})
	for _, f := range fighters {
		println("weight:", f.Weight, "\t", "height:", f.Height)
	}

}
