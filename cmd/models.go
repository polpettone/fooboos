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

	if len(result) == 1 && result[0] == "" {
		return []string{}
	}

	return result
}

func fuzzySearch(keywords []string, query string) []string {
	bagSizes := []int{1}
	cm := closestmatch.New(keywords, bagSizes)
	result := cm.ClosestN(query, 4)
	return result
}
