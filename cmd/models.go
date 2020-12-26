package cmd

import "github.com/schollz/closestmatch"

type Fooboos struct {
	Entries map[string][]string
}

func (fooboos Fooboos) search(query string) []string {
	var keywords []string
	for k, _ := range fooboos.Entries {
		keywords = append(keywords, k)
	}

	result := fuzzySearch(keywords, query)

	return result
}

func fuzzySearch(keywords []string, query string) []string {
	bagSizes := []int{2}
	cm := closestmatch.New(keywords, bagSizes)
	result := cm.ClosestN(query, 4)
	return result
}