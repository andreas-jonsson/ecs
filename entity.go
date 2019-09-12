// Copyright (c) 2016 Ali Najafizadeh
// Copyright (c) 2019 Andreas T Jonsson

package ecs

type entity struct {
	components     []Component
	componentTypes uint32
}

func (e *entity) indexOfComponent(componentTyp uint32) int {
	//position value is 1, 2, 3, ...
	//in order to conver it into index, we need to decrement it by one
	position := calcBitIndex(e.componentTypes, componentTyp)
	return int(position - 1)
}

func (e *entity) Component(componentType uint32) Component {
	//make sure the component exists in list
	if e.componentTypes&componentType == 0 {
		return nil
	}

	index := e.indexOfComponent(componentType)

	return e.components[index]
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
	e.components = e.components[0 : len(e.components)+1]
	copy(e.components[index:], e.components[index+1:])
	e.components[index] = component
}

func (e *entity) RemoveComponent(componentType uint32) {
	//component doesn't have that component
	if e.componentTypes&componentType == 0 {
		return
	}

	index := e.indexOfComponent(componentType)

	//deleting the component from list
	copy(e.components[index:], e.components[index+1:])
	e.components[len(e.components)-1] = nil
	e.components = e.components[:len(e.components)-1]
}

func (e *entity) HasComponentTypes(componentTypes uint32) bool {
	return e.componentTypes&componentTypes != 0
}

func NewEntity() *entity {
	return &entity{}
}
