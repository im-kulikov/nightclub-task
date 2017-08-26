package main

import (
	"fmt"
	"time"
)

var clubTpl = `
===================================================================
                   Клуб '%s' %s
===================================================================
`

type Nightclub struct {
	title    string
	pause    time.Duration
	playlist *Playlist
	visitors []*Visitor
}

func (nc *Nightclub) AddVisitors(visitors ...*Visitor) {
	nc.visitors = append(nc.visitors, visitors...)
}

func (nc *Nightclub) SetPlaylist(playlist *Playlist) {
	nc.playlist = playlist
}

func (nc *Nightclub) Open() {
	fmt.Printf(clubTpl, nc.title, "открылся")
}

func (nc *Nightclub) Close() {
	fmt.Printf(clubTpl, nc.title, "закрылся")
}

func (nc *Nightclub) PlayMusic() {
	for _, song := range nc.playlist.Songs {
		song.Play()

		// visitors->DoActions(song)
		for _, visitor := range nc.visitors {
			visitor.DoActions(song)
		}

		// Wait:
		time.Sleep(nc.pause)
	}
}

func NewClub(name string, pause int) *Nightclub {
	return &Nightclub{
		title: name,
		pause: time.Duration(pause) * time.Second,
	}
}
