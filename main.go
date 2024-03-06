package main

import (
	"fmt"
	"fussball2000/matchengine/ecs"
	"fussball2000/matchengine/ecs/components"
	"fussball2000/matchengine/loop"
)

func main() {
	loop, err := loop.NewMatchLoop(90, (loop.PenaltyShootout), 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loop)

	registry := ecs.NewRegistry()

	entity1, _ := registry.CreateEntity()
	entity2, _ := registry.CreateEntity()
	fmt.Println("(1) Entities:", entity1, entity2)

	positionComponent1 := components.NewPositionComponent(1, 1)
	positionComponent2 := components.NewPositionComponent(2, 2)
	fmt.Println("(2) PositionComponent:", positionComponent1)
	fmt.Println("(2) PositionComponent:", positionComponent2)

	registry.AddComponent(entity1, positionComponent1)
	registry.AddComponent(entity2, positionComponent2)
	registry.AddComponent(entity2, positionComponent2)
	entitiesWithPositionComponent, _ := registry.GetEntitiesWithComponent("PositionComponent")
	fmt.Println("(3) Entities with PositionComponent:", entitiesWithPositionComponent)

	for _, entity := range entitiesWithPositionComponent {
		fmt.Println("(4) Entity:", entity)
		component, _ := registry.GetComponent(entity, "PositionComponent")
		fmt.Println("(4) Component:", component)
	}

	positionComponentFromEntity1, _ := registry.GetComponent(entity1, "PositionComponent")
	positionComponentFromEntity2, _ := registry.GetComponent(entity2, "PositionComponent")
	fmt.Println("(5) ", positionComponentFromEntity1)
	fmt.Println("(5) ", positionComponentFromEntity2)
}
