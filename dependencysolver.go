package dependencysolver

// Entry is a struct containing information about a task, its ID and the dependency list
type Entry struct {
	ID   string
	Deps []string
}

// HasCircularDependency returns false if there are no cycles in the dependency graph
func HasCircularDependency(entries []Entry) bool {
	return (nil == LayeredTopologicalSort(entries))
}

// LayeredTopologicalSort returns a list of layers of entries,
// the entries within each layer can be executed in parallel
func LayeredTopologicalSort(entries []Entry) (layers [][]string) {
	// build the dependencies graph
	dependenciesToFrom := make(map[string]map[string]bool)
	dependenciesFromTo := make(map[string]map[string]bool)
	for _, entry := range entries {
		dependenciesToFrom[entry.ID] = make(map[string]bool)
		for _, dep := range entry.Deps {
			dependenciesToFrom[entry.ID][dep] = true
			if _, ok := dependenciesFromTo[dep]; !ok {
				dependenciesFromTo[dep] = make(map[string]bool)
			}
			dependenciesFromTo[dep][entry.ID] = true
		}
	}

	for len(dependenciesToFrom) > 0 {
		thisIterationIds := make([]string, 0)
		for k, v := range dependenciesToFrom {
			if 0 == len(v) {
				// if an item has zero dependencies, remove it
				thisIterationIds = append(thisIterationIds, k)
			}
		}
		if 0 == len(thisIterationIds) {
			// if nothing was found to remove, there's no valid sort
			return nil
		}

		layer := make([]string, 0)
		for _, id := range thisIterationIds {
			// Remove the found items from the dictionary
			delete(dependenciesToFrom, id)
			// add them to the overall ordering
			layer = append(layer, id)

			// and remove all outbound edges
			if deps, ok := dependenciesFromTo[id]; ok {
				for dep, _ := range deps {
					delete(dependenciesToFrom[dep], id)
				}
			}
		}
		layers = append(layers, layer)
	}
	return layers
}
