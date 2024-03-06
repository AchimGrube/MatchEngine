package loop

import (
	"errors"
	"time"
)

type matchLoop struct {
	fullTime               time.Duration
	halfTime               time.Duration
	extraTime              time.Duration
	matchTime              time.Duration
	stoppageTime           time.Duration
	canHaveExtraTime       bool
	canHavePenaltyshootout bool
}

func NewMatchLoop(minutes time.Duration, options MatchOptions, extraTimeMinutes ...time.Duration) (*matchLoop, error) {
	if options&ExtraTime != 0 && (len(extraTimeMinutes) == 0 || extraTimeMinutes[0] <= 0) {
		return nil, errors.New("when ExtraTime is set, extraTimeMinutes must be provided and greater then 0")
	} else if options&ExtraTime == 0 && len(extraTimeMinutes) > 0 {
		return nil, errors.New("extraTimeMinutes provided but ExtraTime not set")
	}

	var extraTime time.Duration
	if len(extraTimeMinutes) > 0 {
		extraTime = extraTimeMinutes[0]
	}

	return &matchLoop{
		fullTime:               minutes * time.Minute,
		halfTime:               (minutes * time.Minute) / 2,
		extraTime:              extraTime * time.Minute,
		matchTime:              0 * time.Minute,
		stoppageTime:           0 * time.Minute,
		canHaveExtraTime:       options&ExtraTime != 0,
		canHavePenaltyshootout: options&PenaltyShootout != 0,
	}, nil
}

func (m *matchLoop) Run() {
	for {

		//call ecs systems etc

		if m.isHalfTime() {
			m.stoppageTime = 0 * time.Minute
			m.matchTime = m.halfTime
		} else if m.isFullTimeOver() {
			if m.canHaveExtraTime {
				m.stoppageTime = 0 * time.Minute
				m.matchTime = m.fullTime
				m.canHaveExtraTime = false
			} else if m.canHavePenaltyshootout && !m.canHaveExtraTime {
				m.canHavePenaltyshootout = false
			} else {
				break
			}
		}

		m.matchTime += 1 * time.Second
	}
}

func (m *matchLoop) isHalfTime() bool {
	return m.matchTime == m.halfTime+m.stoppageTime
}

func (m *matchLoop) isFullTimeOver() bool {
	return m.matchTime >= m.fullTime+m.stoppageTime+m.extraTime
}
