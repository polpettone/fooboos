package cmd

import (
	"reflect"
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
			keywords: []string{"one", "on", "three"},
			query:    "on",
			want:     []string{"on", "one"},
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuzzySearch(tt.keywords, tt.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fuzzySearch() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestFooboos_search(t *testing.T) {

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fooboos := Fooboos{
				Entries: tt.fields,
			}
			if got := fooboos.search(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

