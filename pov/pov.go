package pov

//TestVersion is the unit tests that this will pass
const TestVersion = 1

//Graph is a directed graph of string nodes
type Graph map[string][]string

/*New creates an empty graph.*/
func New() *Graph {
	return &Graph{}
}

/*AddNode adds a node to the graph.*/
func (g *Graph) AddNode(nodeLabel string) {
	(*g)[nodeLabel] = make([]string, 0)
}

/*AddArc adds and edge to the graph*/
func (g *Graph) AddArc(from, to string) {
	children := (*g)[from]
	(*g)[from] = append(children, to)
}

/*ArcList shows all edges in the graph*/
func (g *Graph) ArcList() []string {
	var edges []string
	for from, v := range *g {
		for _, to := range v {
			edges = append(edges, from+" -> "+to)
		}
	}
	return edges
}

/*ChangeRoot reroots the graph.*/
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	path := g.getPath(oldRoot, newRoot)
	for i := 0; i < len(path)-1; i++ {
		oldTo, oldFrom := path[i], path[i+1]
		g.removeArc(oldFrom, oldTo)
		g.AddArc(oldTo, oldFrom)
	}
	return g
}

/*removeArc removes an edge from the graph.*/
func (g *Graph) removeArc(from, to string) {
	children := (*g)[from]
	var newChilderen []string
	for _, child := range children {
		if child != to {
			newChilderen = append(newChilderen, child)
		}
	}
	(*g)[from] = newChilderen
}

/*getPath finds the list from one node to another.*/
func (g *Graph) getPath(from, to string) []string {
	//Recursive depth first search
	if from == to {
		return []string{to}
	}
	for _, child := range (*g)[from] {
		if path := g.getPath(child, to); path != nil {
			return append(path, from)
		}
	}
	return nil
}
