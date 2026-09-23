package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"

	wildcards "github.com/slenderb13/rolada-russa/pkg"
)

type Dice struct {
	Value int
}

var ctx = context.Background()

func (d *Dice) Roll() {
	randValue := rand.Intn(20)+1
	d.Value = randValue
	fmt.Printf("Dice rolled: D%d\n", d.Value)
}
func (d *Dice) Decide() bool {
	if d.Value < 15	&& d.Value > 10 {
		return true
	}
	if d.Value < 10	&& d.Value > 5 {}
	if d.Value > 15 {
		fmt.Println("You are a lucky bastard!!!")
	}
	return false
}

type ItemCard struct {
	IsActive bool
	Wildcard wildcards.WildCard
}


type Game struct {
	Cards []*ItemCard
	Dice Dice
}



func (g *Game) GetRandomCard() *ItemCard {
	cards := g.GetAvailableCards()

	if len(cards) == 0 {
		return nil
	}

	cardIdx := rand.Intn(len(cards))
		
	return cards[cardIdx]
}

func (g *Game) GetAvailableCards() []*ItemCard {
	cards := make([]*ItemCard, 0, len(g.Cards))
	for _, w := range g.Cards {
		if !w.IsActive {
			cards = append(cards, w)
		}
	}

	return cards
}


func (g *Game) Decide() {
	if g.Dice.Decide() {
		itemCard := g.GetRandomCard()
		if itemCard == nil {
			fmt.Printf("It's over bro.. all cards are being used")
			return
		}
		err := itemCard.Wildcard.Apply(ctx)
		if err != nil {
			fmt.Sprintf("Failed to apply wildcard", err)
		}
		itemCard.IsActive = true
	}
}

func main() {
	cards := []*ItemCard{
		&ItemCard{ Wildcard: wildcards.MouseFuckerCard },
		&ItemCard{ Wildcard: wildcards.MoveFuckerCard },
		&ItemCard{ Wildcard: wildcards.RickRollFucker },
		&ItemCard{ Wildcard: wildcards.SoundFucker },
	}
	session := Game{
		cards,
		Dice{ Value: 0 },
	}

	scanner := bufio.NewScanner(os.Stdin)
	shouldRun := true
	
	for shouldRun {
		fmt.Println("\nPress r to roll the dice or e to exit...")
		if scanner.Scan() {
		input := scanner.Text()
			switch input {
			case "r":
				session.Dice.Roll()
				session.Decide()
			case "e":
				shouldRun = false
			}
		}
	}
}
