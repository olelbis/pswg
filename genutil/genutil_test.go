package genutil

import (
	"testing"
)

func TestMeleeEmpty(t *testing.T) {
	melee, err := Melee("")
	if err != nil {
		t.Error("Erroreeeee: ", err)
	}
	t.Log(melee)
}

func TestMeleeFull(t *testing.T) {
	melee, err := Melee("S0n0B1sKu3D0!!")
	if err != nil {
		t.Error("Erroreeeee")
	}
	t.Log("OOOOK", melee)
}
