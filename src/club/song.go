package main

import "fmt"

const songTpl = `
-------------------------------------------------------------------
Now playing -> %s / %s
-------------------------------------------------------------------
`

const (
	genreHouse = iota
	genrePop
	genreRnb
	genreBlack
)

type Song struct {
	name  string
	genre int
}

func (s Song) Play() {
	fmt.Printf(songTpl, s.name, s.getGenre())
}

func (s Song) getGenre() string {
	switch s.genre {
	case genreHouse:
		return "House"
	case genrePop:
		return "Pop"
	case genreRnb:
		return "Rnb"
	case genreBlack:
		return "BlackMetal"
	}

	return "Unknown"
}

func NewSong(name string, genre int) *Song {
	return &Song{name, genre}
}
