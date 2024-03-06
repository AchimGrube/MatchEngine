package ecs

type registry struct {
	entities       map[Entity]uint64
	nextEntity     uint8
	components     map[Entity]map[string]Component
	componentMasks map[string]uint64
}

func NewRegistry() *registry {
	return &registry{
		entities:       make(map[Entity]uint64),
		nextEntity:     1,
		components:     make(map[Entity]map[string]Component),
		componentMasks: make(map[string]uint64),
	}
}

func (r *registry) CreateEntity() (Entity, bool) {
	ok := r.nextEntity+1 <= 255
	if ok {
		entity := Entity(r.nextEntity)
		r.components[entity] = make(map[string]Component)
		r.entities[entity] = 0
		r.nextEntity++
		return entity, ok
	}
	return 0, ok
}

func (r *registry) DestroyEntity(entity Entity) bool {
	_, entityOk := r.entities[entity]
	_, componentOk := r.components[entity]
	if entityOk && componentOk {
		delete(r.entities, entity)
		delete(r.components, entity)
	}
	return entityOk && componentOk
}

func (r *registry) AddComponent(entity Entity, component Component) bool {
	_, ok := r.componentMasks[component.GetName()]
	if !ok {
		r.componentMasks[component.GetName()] = component.GetMask()
	}
	_, ok = r.components[entity][component.GetName()]
	if !ok {
		r.components[entity][component.GetName()] = component
		mask := r.entities[entity]
		mask |= component.GetMask()
		r.entities[entity] = mask
	}
	return ok
}

func (r *registry) RemoveComponent(entity Entity, componentName string) bool {
	component, ok := r.components[entity][componentName]
	if ok {
		delete(r.components[entity], componentName)
		mask := r.entities[entity]
		mask &= ^(component.GetMask())
		r.entities[entity] = mask
	}
	return ok
}

func (r *registry) GetComponent(entity Entity, componentName string) (Component, bool) {
	ok := r.hasEntityComponent(entity, componentName)
	if ok {
		component := r.components[entity][componentName]
		return component, ok
	}
	return nil, ok
}

func (r *registry) GetEntitiesWithComponent(componentName string) ([]Entity, int) {
	var entities []Entity
	for entity, mask := range r.entities {
		if (mask & r.componentMasks[componentName]) != 0 {
			entities = append(entities, entity)
		}
	}
	return entities, len(entities)
}

func (r *registry) hasEntityComponent(entity Entity, componentName string) bool {
	mask, ok := r.entities[entity]
	if ok {
		return (mask & r.componentMasks[componentName]) != 0
	}
	return false
}
