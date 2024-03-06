package components

type positionComponent struct {
	X, Y int
}

func NewPositionComponent(x, y int) *positionComponent {
	return &positionComponent{
		X: x,
		Y: y,
	}
}

func (c *positionComponent) GetName() string {
	return PositionComponentName
}

func (c *positionComponent) GetMask() uint64 {
	return PositionComponentMask
}
