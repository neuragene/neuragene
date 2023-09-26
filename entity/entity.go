// Package entity contains the logic to handle entities.
package entity

import (
	"sync/atomic"

	"github.com/arsham/lify/component"
)

// Mask is used to determine if an entity supports a specific logic.
type Mask uint16

const (
	// Positioned mask indicates that the entity has a position, scale and
	// velocity.
	Positioned Mask = 1 << iota
)

// An Entity is an element in the game that can have at least one component.
// Each component is managed by the component.Manager as a map if entity ID to
// its respective component. When an Entity is removed, its ID is removed from
// all the maps in the Manager. You should not create an Entity directly,
// instead you should use the Manager's NewEntity method.
type Entity struct {
	mask Mask
	ID   uint64
}

// List contains a slice of Entity.
type List []*Entity

// Manager manages all the entities in the game. It is advised to create a new
// Manager by calling the NewManager constructor to pre-allocate memory.
type Manager struct {
	components *component.Manager
	entities   List
	// When a new entity is added, it will first go into this list. At the end
	// of the frame, all the entities in this list will be added to the
	// entities list.
	toAdd List
	// lastID is the last ID that was given to an entity.
	lastID uint64
}

// NewManager returns a new Manager with pre-allocated memory by the given
// size.
func NewManager(size int) *Manager {
	return &Manager{
		components: &component.Manager{
			Position: make(map[uint64]*component.Position, size),
		},
		entities: make(List, 0, size),
		toAdd:    make(List, 0, size),
	}
}

// NewEntity creates a new Entity with the given mask. The entity is not added
// to the available entities, but it will be on the next Update call. You
// should manually set the necessary components for the entity. Your call will
// panic if the component you are trying to set doesn't match the mask.
func (m *Manager) NewEntity(mask Mask) *Entity {
	e := &Entity{
		ID:   atomic.AddUint64(&m.lastID, 1),
		mask: mask,
	}
	m.toAdd = append(m.toAdd, e)
	if mask&Positioned == Positioned {
		m.components.Position[e.ID] = &component.Position{}
	}
	return e
}
