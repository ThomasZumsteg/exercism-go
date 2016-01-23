package tree

import (
	"fmt"
	"sort"
)

//TestVersion is the unit test version this will pass
const TestVersion = 2

// thomas:tree-building$ go test -bench .
// PASS
// BenchmarkTwoTree-4             1        12256812916 ns/op
// BenchmarkTenTree-4             5         281389628 ns/op
// BenchmarkShallowTree-4         2         919570690 ns/op
// ok      _/Users/thomas/programming/exercism/go/tree-building    17.588s

//Record is a database item that specifies a parent record in a tree structure.
type Record struct {
	ID, Parent int
}

//String converts a record to a string.
func (r *Record) String() string {
	return fmt.Sprintf("(ID: %d, Parent: %d)", r.ID, r.Parent)
}

//Interfaces for storting records.
type byID []Record

func (ids byID) Len() int           { return len(ids) }
func (ids byID) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }
func (ids byID) Less(i, j int) bool { return ids[i].ID < ids[j].ID }

//Node is a node in a tree structure.
type Node struct {
	ID       int
	Children []*Node
}

/*Build converts a set of records to a tree of Nodes.*/
func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}
	sort.Sort(byID(records))
	nodes := make([]*Node, len(records))
	for r, rec := range records {
		nodes[r] = &Node{ID: rec.ID}
		if r == 0 && (rec.ID != 0 || rec.Parent != 0) {
			return nil, fmt.Errorf("Not a valid root record %s", rec.String())
		} else if r == 0 {
			continue
		} else if r != rec.ID || rec.ID <= rec.Parent {
			return nil, fmt.Errorf("Not a valid record: %s", rec.String())
		}

		if parent := nodes[rec.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[r])
		}
	}
	return nodes[0], nil
}
