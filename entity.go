// Copyright (c) 2016 Ali Najafizadeh
// Copyright (c) 2019 Andreas T Jonsson

package ecs

import (
	"math"
)

type entity struct {
	components     []Component
	id             uint64
	componentTypes uint32
}

func (e *entity) indexOfComponent(componentTyp uint32) int {
	//position value is 1, 2, 3, ...
	//in order to conver it into index, we need to decrement it by one
	position := calcBitIndex(e.componentTypes, componentTyp)
	return int(position - 1)
}

func (e *entity) ID() uint64 {
	return e.id
}

func (e *entity) Component(componentType uint32) Component {
	//make sure the component exists in list
	if e.componentTypes&componentType == 0 {
		return nil
	}

	index := e.indexOfComponent(componentType)

	return e.components[index]
}

func (e *entity) Components() []Component {
	return e.components
}

func (e *entity) AddComponent(component Component) {
	//make sure that component always passes as non nil value
	if component == nil {
		return
	}

	componentType := component.ComponentType()

	//component already inside this entity
	if e.componentTypes&componentType != 0 {
		return
	}

	e.componentTypes = e.componentTypes | componentType

	index := e.indexOfComponent(componentType)

	//insert the new component into right index
	e.components = append(e.components, nil)
	copy(e.components[index+1:], e.components[index:])
	e.components[index] = component
}

func (e *entity) RemoveComponent(componentType uint32) Component {
	//component doesn't have that component
	if e.componentTypes&componentType == 0 {
		return nil
	}

	index := e.indexOfComponent(componentType)
	c := e.components[index]

	//deleting the component from list
	copy(e.components[index:], e.components[index+1:])
	e.components = e.components[:len(e.components)-1]
	return c
}

func (e *entity) HasComponentTypes(componentTypes uint32) bool {
	return e.componentTypes&componentTypes != 0
}

var idCounter uint64

func NewEntity() *entity {
	if idCounter == math.MaxUint64 {
		panic("id space exhausted")
	}
	idCounter++
	return &entity{id: idCounter}
}
