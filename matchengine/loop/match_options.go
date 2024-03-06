package loop

type MatchOptions uint8

const (
	None      MatchOptions = 0
	ExtraTime MatchOptions = 1 << iota
	PenaltyShootout
)
