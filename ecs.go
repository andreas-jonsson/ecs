// Copyright (c) 2016 Ali Najafizadeh
// Copyright (c) 2019 Andreas T Jonsson

package ecs

import (
	"time"
)

// Destructible defines a component or entity that needs cleanup.
type Destructible interface {
	Destroy()
}

// Destroy calls the Destroy function (if it exist) on a entity and all it's components.
// If o is a component it will call Destroy if it exist.
func Destroy(o interface{}) {
	if e, ok := o.(*entity); ok {
		for _, c := range e.components {
			Destroy(c)
		}
	}
	if d, ok := o.(Destructible); ok {
		d.Destroy()
	}
}

// Component represents the abstract version of data.
// each component will associate with a meaningful data.
// because of underneath structure, we can have upto
// 32 components in total per game.
// hope that enough...
type Component interface {
	ComponentType() uint32
}

// Entity is a shell around multiple components.
// components can be added or removed from entity either at
// runtime or compile time.
type Entity interface {
	Component(typ uint32) Component
	AddComponent(component Component)
	RemoveComponent(componentType uint32) Component
	HasComponentTypes(componentTypes uint32) bool
}

// Query is a bridge between System and Entity.
// query can be used to fetch array of entities
// based on the data that they have or accessing
// an already exist system.
type Query interface {
	Entities(componentTypes uint32) []Entity
	ForAllEntities(componentTypes uint32, f func(e Entity) bool)
	AddEntity(Entity)
	RemoveEntity(Entity)
	System(systemType uint32) System
}

// System contains all the methods which will needed for
// hopfully all the scenarioes. There is an Update method
// that is being called by World and was given a Query object
// for manipulating entities.
type System interface {
	SystemType() uint32
	Active(stage int) bool
	Update(delta time.Duration, query Query)
}

// Manager is a simple abstraction of minimum requirement of
// method which needed for any game.
type Manager interface {
	AddSystem(system System)
	RemoveSystem(systemType uint32) System
	Update(stage int, delta time.Duration)
}
