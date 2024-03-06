package ecs

type Component interface {
	GetName() string
	GetMask() uint64
}
