package main

import (
	"masonwordle/answerflag"
	"masonwordle/game"
)

type converter struct {
}

func (converter) convertToAnswerFlag(g *game.GameState, setAF bool, showAF bool) *answerflag.AnswerInfoHolder {
	return &answerflag.AnswerInfoHolder{
		AnswerList:     g.AnswerList,
		ValidAnswers:   g.ValidAnswers,
		SetAnswerFlag:  setAF,
		ShowAnswerFlag: showAF,
	}
}
