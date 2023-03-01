package main_test

import (
	"masonwordle/game"
	"testing"
	// . "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

func TestIsValidAnswerOrGuess(t *testing.T) {
	testMap := make(map[string]string)
	testMap["test"] = "test"
	if !game.IsValidGuessOrAnswer(testMap, "test") {
		t.Errorf("Failed on valid input test.")
	}
	testMap1 := make(map[string]string)
	testMap1["t"] = "test"
	if game.IsValidGuessOrAnswer(testMap1, "test") {
		t.Errorf("Failed on valid input test.")
	}
	testMap2 := make(map[string]string)
	testMap2["test"] = "t"
	if game.IsValidGuessOrAnswer(testMap2, "test") {
		t.Errorf("Failed on valid input test.")
	}
	testMap3 := make(map[string]string)
	testMap3["test"] = "test"
	if game.IsValidGuessOrAnswer(testMap3, "t") {
		t.Errorf("Failed on valid input test.")
	}
	testMap4 := make(map[string]string)
	testMap4["excep"] = "exc"
	testMap4["test"] = "test"
	if !game.IsValidGuessOrAnswer(testMap4, "test") {
		t.Errorf("Failed on valid input test.")
	}

}

func TestIsValidInput(t *testing.T) {
	if game.IsValidInput("hello", 4) {
		t.Errorf("Failed on valid input test.")
	}

	if game.IsValidInput("hello", 6) {
		t.Errorf("Failed on valid input test.")
	}
	if game.IsValidInput("", 6) {
		t.Errorf("Failed on valid input test.")
	}
}

func TestIsValidListSize(t *testing.T) {
	x := []string{"hello"}
	if !game.IsValidListSize(x) {
		t.Errorf("IsValidListSize test failed.")
	}
	b := []string{}
	if game.IsValidListSize(b) {
		t.Errorf("IsValidListSize test failed.")
	}
}

func TestIsValidWord(t *testing.T) {
	x := "hello"
	if !game.IsValidWord(x) {
		t.Errorf("IsValidWord test failed.")
	}
	b := ""
	if game.IsValidWord(b) {
		t.Errorf("IsValidWord test failed.")
	}

	c := "hel"
	if game.IsValidWord(c) {
		t.Errorf("IsValidWord test failed.")
	}

	d := "hellos"
	if game.IsValidWord(d) {
		t.Errorf("IsValidWord test failed.")
	}

}
