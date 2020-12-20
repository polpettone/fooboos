package cmd

import (
	"testing"
)

func Test_loadFooboo(t *testing.T) {
	fooboos, err := loadFooboos("test-data/fooboos.yaml")
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(fooboos.Entries) != 3 {
		t.Errorf("expected %d bookmarks got %d", 2, len(fooboos.Entries))
	}
}
