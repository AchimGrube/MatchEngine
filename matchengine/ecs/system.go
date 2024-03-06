package ecs

type System interface {
	OnCreate()
	OnStart()
	OnTick()
	OnEnd()
	OnDestroy()
}
