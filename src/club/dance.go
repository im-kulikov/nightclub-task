package main

import (
	"strings"
)

type Dance interface {
	moveHead() string
	moveArms() string
	moveBody() string
	moveLegs() string
}

//- Когда играет Rnb на танцполе танцуют те, кто танцуют хип-хоп, рнб.
//- Когда играет Electrohuse на танцполе танцуют те, кто танцуют Electrodance.
//- Когда играет Поп-музыка танцуют те кто умеют танцевать под поп-музыку.
//- Если человек не умеет танцевать под данную музыку, он идет в бар и пьет
//водку.

func GoDance(dance Dance) string {
	if dance == nil {
		return "пьёт водку"
	}

	action := strings.Join([]string{
		dance.moveHead(),
		dance.moveArms(),
		dance.moveBody(),
		dance.moveLegs(),
	}, ", ")

	return "танцует: " + action + "."
}

func GetGenre(dance Dance) int {
	switch dance.(type) {
	case Rnb:
		return genreRnb
	case Pop:
		return genrePop
	case House:
		return genreHouse
	case Black:
		return genreBlack
	}

	return -1
}

type Rnb struct{}

func (d Rnb) moveHead() string {
	return "Двигает головой вперед-назад"
}

func (d Rnb) moveArms() string {
	return "Руки согнуты в локтях"
}

func (d Rnb) moveBody() string {
	return "Покачиваниет телом вперед и назад"
}

func (d Rnb) moveLegs() string {
	return "Ноги в полу-присяде"
}

type House struct{}

func (d House) moveHead() string {
	return "Изредка двигает головой"
}

func (d House) moveArms() string {
	return "Вращает руками"
}

func (d House) moveBody() string {
	return "Покачивает туловищем вперед-назад"
}

func (d House) moveLegs() string {
	return "Ноги двигаются в ритме"
}

type Pop struct{}

func (d Pop) moveHead() string {
	return "Плавные движения головой"
}

func (d Pop) moveArms() string {
	return "Плавные движения руками"
}

func (d Pop) moveBody() string {
	return "Плавные движения туловищем"
}

func (d Pop) moveLegs() string {
	return "Плавные движения ногами"
}

type Black struct{}

func (d Black) moveHead() string {
	return "Круговые движения головой"
}

func (d Black) moveArms() string {
	return `Рука поднята над головой, пальцами демонстрирует "козу"`
}

func (d Black) moveBody() string {
	return "Тело напряжено"
}

func (d Black) moveLegs() string {
	return "Ноги раставлены на ширине плечей"
}
