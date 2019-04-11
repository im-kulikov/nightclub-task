package main

import (
	"math/rand"
	"time"
)

func main() {
	var (
		club     *Nightclub
		p        *Playlist
		visitor  *Visitor
		visitors []*Visitor
	)

	rand.Seed(time.Now().UnixNano())

	club = NewClub("My NightClub", 2)

	p = new(Playlist)

	p.Add(
		NewSong("Arash feat. Helena - One Day", genrePop),
		NewSong("Бьянка - Ногами Руками", genreRnb),
		NewSong("Inna feat. J Balvin - Cola Song", genreHouse),
		NewSong("Immortal - Blashyrkh (Mighty Ravendark)", genreBlack),
	)

	club.SetPlaylist(p)

	visitors = make([]*Visitor, 0)

	visitor = NewVisitor("Света")
	visitor.AddSkills(Pop{}, Rnb{})
	visitors = append(visitors, visitor)

	visitor = NewVisitor("Даша")
	visitor.AddSkills(House{}, Pop{})
	visitors = append(visitors, visitor)

	visitor = NewVisitor("Вася")
	visitor.AddSkills(Pop{}, Black{})
	visitors = append(visitors, visitor)

	visitor = NewVisitor("Петя")
	visitor.AddSkills(Black{}, House{})
	visitors = append(visitors, visitor)

	visitor = NewVisitor("Коля")
	visitor.AddSkills(House{}, Rnb{})
	visitors = append(visitors, visitor)

	club.AddVisitors(visitors...)

	club.Open()
	club.PlayMusic()
	club.Close()
}
