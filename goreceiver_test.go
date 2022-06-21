/*
Requirements:
Write unit tests that ensure:
- Health & energy cannot go above their maximums
- Health & energy are set correctly

If any of your tests fail, make the necessary corrections
in the copy of your solution file
*/

package main

import "testing"

func newPlayer() Player {
	return Player{100, 100, 100, 100, "Test"}
}

func TestSetHealth(t *testing.T) {
	p := newPlayer()
	p.setHealth(p.maxHealth + 100)
	if p.health > p.maxHealth {
		t.Fatalf(
			"Health exceeded limit. Expected %v, not %v",
			p.maxHealth,
			p.health,
		)
	}
	p.setHealth(20)
	if p.health != 20 {
		t.Fatalf(
			"Health incorrect. Expected %v, not %v",
			20,
			p.health,
		)
	}
}

func TestSetEnergy(t *testing.T) {
	p := newPlayer()
	p.setEnergy(p.maxEnergy + 100)
	if p.energy > p.maxEnergy {
		t.Fatalf(
			"Energy exceeded limit. Expected %v, not %v",
			p.maxEnergy,
			p.health,
		)
	}
	p.setEnergy(20)
	if p.energy != 20 {
		t.Fatalf(
			"Energy incorrect. Expected %v, not %v",
			20,
			p.energy,
		)
	}
}
