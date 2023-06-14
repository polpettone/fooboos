package cmd

import (
	"sort"
	"testing"
)

func Test_fuzzySearch(t *testing.T) {
	tests := []struct {
		name     string
		keywords []string
		query    string
		want     []string
	}{

		{
			name:     "0",
			keywords: []string{"one", "two", "three"},
			query:    "one",
			want:     []string{"one"},
		},

		{
			name:     "1",
			keywords: []string{"one", "two", "three"},
			query:    "on",
			want:     []string{"one"},
		},

		{
			name:     "2",
			keywords: []string{"on", "one", "three"},
			query:    "on",
			want:     []string{"on", "one"},
		},

		{
			name:     "3",
			keywords: []string{"one", "on", "three"},
			query:    "onez",
			want:     []string{"one", "on", "three"},
		},

		{
			name:     "4",
			keywords: []string{"one", "on", "three"},
			query:    "thr",
			want:     []string{"three"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuzzySearch(tt.keywords, tt.query); !areSlicesEqual(got, tt.want) {
				t.Errorf("fuzzySearch() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestFoobars_search(t *testing.T) {

	tests := []struct {
		name   string
		fields map[string][]string
		args   string
		want   []string
	}{

		{name: "find exactly",
			fields: map[string][]string{
				"one": {"url0", "url1"},
				"two": {"url0", "url1"},
			},
			args: "one",
			want: []string{"one"}},

		{name: "find fuzzy",
			fields: map[string][]string{
				"one": {"url0", "url1"},
				"two": {"url0", "url1"},
			},
			args: "on",
			want: []string{"one"}},

		{name: "find nothing",
			fields: map[string][]string{
				"one": {"url0", "url1"},
				"two": {"url0", "url1"},
			},
			args: "o",
			want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fooboos := Fooboos{
				Entries: tt.fields,
			}
			if got := fooboos.search(tt.args); !areSlicesEqual(got, tt.want) {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func areSlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	// Sortiere die Slices
	sort.Strings(slice1)
	sort.Strings(slice2)

	// Überprüfe, ob die sortierten Slices gleich sind
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
