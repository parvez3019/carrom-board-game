package player

import "github.com/google/uuid"

type Player struct {
	id                  string
	name                string
	score               int
	successiveFoulCount int
}

func NewPlayer(name string) *Player {
	return &Player{name: name, id: uuid.New().String()}
}

func (p *Player) UpdateScore(pointsScored int) *Player {
	p.score += pointsScored
	if pointsScored > 0 {
		p.successiveFoulCount = 0
	}
	return p
}

func (p *Player) Foul() *Player {
	p.successiveFoulCount += 1
	p.score -= 1
	if p.successiveFoulCount%3 == 0 {
		p.score -= 1
	}
	return p
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Score() int {
	return p.score
}

func (p *Player) Id() string {
	return p.id
}