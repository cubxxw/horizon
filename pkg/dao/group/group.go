package group

import (
	"strings"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name            string
	Path            string
	VisibilityLevel string
	Description     string
	ParentID        uint
	TraversalIDs    string
}

type Groups []*Group

// Len the length of the groups
func (g Groups) Len() int {
	return len(g)
}

// Less sort groups by the size of the traversalIDs array after split by ','
func (g Groups) Less(i, j int) bool {
	return len(strings.Split(g[i].TraversalIDs, ",")) < len(strings.Split(g[j].TraversalIDs, ","))
}

// Swap the two group
func (g Groups) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}