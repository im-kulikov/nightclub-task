package main

func main() {
	var (
		club     *Nightclub
		p        *Playlist
		visitor  *Visitor
		visitors []*Visitor
	)

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

	visitor = NewVisitor("Вася")
	visitor.AddSkills(Pop{}, Black{})
	visitors = append(visitors, visitor)

	visitor = NewVisitor("Петя")
	visitor.AddSkills(Pop{}, House{})
	visitors = append(visitors, visitor)

	visitor = NewVisitor("Коля")
	visitor.AddSkills(House{}, Rnb{})
	visitors = append(visitors, visitor)

	club.AddVisitors(visitors...)

	club.Open()
	club.PlayMusic()
	club.Close()
}
