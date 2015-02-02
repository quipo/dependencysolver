package dependencysolver

import (
	"sort"
	"testing"
)

var entries []Entry

func init() {
	entries = make([]Entry, 0)
	entries = append(entries, Entry{ID: "A"})
	entries = append(entries, Entry{ID: "B", Deps: []string{"A"}})
	entries = append(entries, Entry{ID: "C", Deps: []string{"A"}})
	entries = append(entries, Entry{ID: "D", Deps: []string{"B", "C"}})
}

func TestHasCircularDependency(t *testing.T) {
	if true == HasCircularDependency(entries) {
		t.Error("Should not have circular dependencies")
	}
}

func TestLayeredTopologicalSort(t *testing.T) {
	actual := LayeredTopologicalSort(entries)
	if nil == actual {
		t.Error("Failed calculating dependency tree")
	}
	if len(actual) < 3 {
		t.Errorf("Should have had 3 layers, %d found.", len(actual))
	}
	if !equalSlices(actual[0], []string{"A"}) {
		t.Errorf("Expecting [A] at the 1st layer, actual: %v", actual[0])
	}
	if !equalSlices(actual[1], []string{"B", "C"}) {
		t.Errorf("Expecting [B,C] at the 2nd layer, actual: %v", actual[1])
	}
	if !equalSlices(actual[2], []string{"D"}) {
		t.Errorf("Expecting [D] at the 3rd layer, actual: %v", actual[2])
	}
}

func TestLayeredTopologicalSortNoDeps(t *testing.T) {
	var entries []Entry
	entries = append(entries, Entry{ID: "A"})
	entries = append(entries, Entry{ID: "B"})
	entries = append(entries, Entry{ID: "C"})

	actual := LayeredTopologicalSort(entries)
	if nil == actual {
		t.Error("Failed calculating dependency tree")
	}
	if len(actual) != 1 {
		t.Errorf("Should have had 1 layer, %d found.", len(actual))
	}
	if !equalSlices(actual[0], []string{"A", "B", "C"}) {
		t.Errorf("Expecting [A,B,C] at the 1st layer, actual: %v", actual[0])
	}
}

func equalSlices(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// last test, add a circular dependency from D to A
func TestHasCircularDependency2(t *testing.T) {
	entries[0].Deps = []string{"D"}
	if false == HasCircularDependency(entries) {
		t.Error("Should have circular dependencies")
	}
}
