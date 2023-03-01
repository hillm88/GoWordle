package main_test

import (
	"masonwordle/game"
	"testing"
	// . "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

func TestLetterPositionSetter(t *testing.T) {

	z, _ := game.LetterPositionSetter("hello", "jello", false)
	if z {
		t.Fatalf("failed")
	}
	x, y := game.LetterPositionSetter("hello", "hello", false)
	if !x {
		t.Fatalf("failed")
	}
	b := [...]int{3, 3, 3, 3, 3}

	for i := 0; i < len(y); i++ {
		if y[i] != b[i] {
			t.Fatalf("failed")
		}
	}

	a, c := game.LetterPositionSetter("sewer", "river", false)
	if a {
		t.Fatalf("failed")
	}
	m := [...]int{1, 1, 1, 3, 3}

	for l := 0; l < len(y); l++ {
		if m[l] != c[l] {
			t.Fatalf("failed")
		}
	}

	w, g := game.LetterPositionSetter("river", "sewer", true)
	if w {
		t.Fatalf("failed")
	}
	j := [...]int{1, 1, 1, 3, 4}

	for o := 0; o < len(g); o++ {
		if g[o] != j[o] {
			t.Fatalf("failed")
		}
	}

	ff, gg := game.LetterPositionSetter("crane", "facet", true)
	if ff {
		t.Fatalf("failed")
	}
	jj := [...]int{1, 2, 2, 2, 1}
	for kk := 0; kk < len(gg); kk++ {
		if gg[kk] != jj[kk] {
			t.Fatalf("failed")
		}
	}

}
