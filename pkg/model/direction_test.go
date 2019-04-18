package model

import "testing"

func TestDirection_Clockwise(t *testing.T) {
	d := Left
	if d.Clockwise() != Up {
		t.Error("clockwise does not work")
	}
}

func TestDirection_Invert(t *testing.T) {
	d := Up
	if d.Invert() != Down {
		t.Error("invert failed")
	}
}

func TestDirection_Change(t *testing.T) {
	d := Up
	p := Point{4, 6}
	expected := Point{4, 5}
	if expected != d.Change(p) {
		t.Error("change did not happen")
	}
}
