package wildcards

import (
	"context"
	"fmt"
)


type Command struct {
	Name string
	Args string
}

type WildCard struct {
	Level string
	Description string
	// Command Command
	execute func (ctx context.Context) error
}


const (

	LEVEL_EASY = "EASY"
	LEVEL_MEDIUM = "MEDIUM"
	LEVEL_HARD = "HARD"
)

func (w *WildCard) Apply(ctx context.Context) error {
	fmt.Printf("Executing command: %s", w.Description)
	//cmd := exec.Command(w.Command.Name, w.Command.Args)

	//_, err := cmd.Output()
	err := w.execute(ctx)
	if err != nil {
		fmt.Println("Error while executing command", err)
		return err
	}

	return nil
}

func GetRandomByLevel(level string) WildCard {
	card := WildCard{
		Level: "easy",
		Description: "Simply prints Hello World.",
		execute: func(ctx context.Context) error {return nil},
		
	}

	return card
}
