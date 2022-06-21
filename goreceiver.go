/*
Summary:
Implement receiver functions to create stat modifications
for a video game character.

Requirements:
Implement a player having the following statistics:
- Health, Max Health
- Energy, Max Energy
- Name

Implement receiver functions to modify the `Health` and `Energy`
statistics of the player.
- Print out the statistic change within each function
- Execute each function at least once
*/

package main

import "fmt"

type Player struct {
	health    int
	maxHealth int
	energy    int
	maxEnergy int
	name      string
}

func (player *Player) setHealth(health int) {
	h := 0
	if health < 0 {
		h = 0
	} else if health > player.maxHealth {
		h = player.maxHealth
	} else {
		h = health
	}
	player.health = h
	fmt.Println(player.name, "'s health has been set to", h)
}

func (player *Player) setEnergy(energy int) {
	e := 0
	if energy < 0 {
		e = 0
	} else if energy > player.maxEnergy {
		e = player.maxEnergy
	} else {
		e = energy
	}
	player.energy = e
	fmt.Println(player.name, "'s energy has been set to", e)
}

func main() {
	warlock := Player{100, 100, 100, 100, "Abe"}
	warlock.setHealth(warlock.health - 20)
	warlock.setEnergy(warlock.energy - 50)
}
