package dependencysolver

import (
	"testing"
)

var augmenters []Entry

func init() {
	augmenters = make([]Entry, 0)
	augmenters = append(augmenters, Entry{Id: "A"})
	augmenters = append(augmenters, Entry{Id: "B", Deps: []string{"A"}})
	augmenters = append(augmenters, Entry{Id: "C", Deps: []string{"A"}})
	augmenters = append(augmenters, Entry{Id: "D", Deps: []string{"B", "C"}})
}

func TestHasCircularDependency(t *testing.T) {
	if true == HasCircularDependency(augmenters) {
		t.Error("Should not have circular dependencies")
	}
}

func TestLayeredTopologicalSort(t *testing.T) {
	actual := LayeredTopologicalSort(augmenters)
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

func equalSlices(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// last test, add a circular dependency from D to A
func TestHasCircularDependency2(t *testing.T) {
	augmenters[0].Deps = []string{"D"}
	if false == HasCircularDependency(augmenters) {
		t.Error("Should have circular dependencies")
	}
}
