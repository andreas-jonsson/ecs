// Copyright (c) 2016 Ali Najafizadeh
// Copyright (c) 2019 Andreas T Jonsson

package ecs

import (
	"time"
)

type World struct {
	systemTypes uint32
	systems     []System
	entities    []Entity
}

func (w *World) indexOfSystem(systemType uint32) int {
	position := calcBitIndex(w.systemTypes, systemType)
	return int(position - 1)
}

func (w *World) AddSystem(system System) {
	//make sure that system always passes as non nil value
	if system == nil {
		return
	}

	systemType := system.SystemType()

	//system already inside this entity32
	if w.systemTypes&systemType != 0 {
		return
	}

	w.systemTypes = w.systemTypes | systemType

	index := w.indexOfSystem(systemType)

	//insert the new system into right index
	w.systems = append(w.systems, nil)
	copy(w.systems[index+1:], w.systems[index:])
	w.systems[index] = system
}

func (w *World) RemoveSystem(systemType uint32) {
	//system doesn't have that system
	if w.systemTypes&systemType == 0 {
		return
	}

	index := w.indexOfSystem(systemType)

	//deleting the system from list
	copy(w.systems[index:], w.systems[index+1:])
	w.systems = w.systems[:len(w.systems)-1]
}

func (w *World) Update(stage int, delta time.Duration) {
	for _, system := range w.systems {
		if system.Active(stage) {
			system.Update(delta, w)
		}
	}
}

func (w *World) Entities(componentTypes uint32) []Entity {
	var entities []Entity

	for _, entity := range w.entities {
		if entity.HasComponentTypes(componentTypes) {
			entities = append(entities, entity)
		}
	}

	return entities
}

func (w *World) System(systemType uint32) System {
	if w.systemTypes&systemType == 0 {
		return nil
	}

	index := w.indexOfSystem(systemType)
	return w.systems[index]
}

func (w *World) AddEntity(entity Entity) {
	w.entities = append(w.entities, entity)
}

func (w *World) RemoveEntity(target Entity) {
	index := -1
	for i, entity := range w.entities {
		if entity == target {
			index = i
			break
		}
	}

	if index > -1 {
		copy(w.entities[index:], w.entities[index+1:])
		w.entities = w.entities[:len(w.entities)-1]
	}
}
