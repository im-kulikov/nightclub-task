package main

import "fmt"

type Visitor struct {
	name   string
	skills []Dance
}

func (v *Visitor) AddSkills(skills ...Dance) {
	v.skills = append(v.skills, skills...)
}

func (v *Visitor) DoActions(song *Song) {
	var (
		sk Dance
	)

	for _, skill := range v.skills {
		if GetGenre(skill) == song.genre {
			sk = skill
			break
		}
	}

	action := GoDance(sk)

	fmt.Printf("Посетитель '%s' %s\n", v.name, action)
}

func NewVisitor(name string) *Visitor {
	return &Visitor{name: name}
}
