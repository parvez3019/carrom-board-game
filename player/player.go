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

func (this *Player) UpdateScore(pointsScored int) *Player {
	this.score += pointsScored
	if pointsScored > 0 {
		this.successiveFoulCount = 0
	}
	return this
}

func (this *Player) Foul() *Player {
	this.successiveFoulCount += 1
	this.score -= 1
	if this.successiveFoulCount%3 == 0 {
		this.score -= 1
	}
	return this
}

func (this *Player) Name() string {
	return this.name
}

func (this *Player) Score() int {
	return this.score
}

func (this *Player) Id() string {
	return this.id
}