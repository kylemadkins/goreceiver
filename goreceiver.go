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
	health    uint
	maxHealth uint
	energy    uint
	maxEnergy uint
	name      string
}

func (player *Player) setHealth(health uint) {
	if health > player.maxHealth {
		player.health = player.maxHealth
	} else {
		player.health = health
	}
	fmt.Println(player.name, "'s health has been set to", player.health)
}

func (player *Player) setEnergy(energy uint) {
	if energy > player.maxEnergy {
		player.energy = player.maxEnergy
	} else {
		player.energy = energy
	}
	fmt.Println(player.name, "'s energy has been set to", player.energy)
}

func main() {
	warlock := Player{100, 100, 100, 100, "Abe"}
	warlock.setHealth(warlock.health - 20)
	warlock.setEnergy(warlock.energy - 50)
}
