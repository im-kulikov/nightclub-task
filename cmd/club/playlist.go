package main

type Playlist struct {
	Songs []*Song
}

func (pl *Playlist) Add(songs ...*Song) {
	pl.Songs = append(pl.Songs, songs...)
}
