package wildcards

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

var SoundFucker = WildCard{
	Level: LEVEL_EASY,
	Description: "Plays a really annoying sound at random times.",
	execute: func(ctx context.Context) error {
		go func() {
			for true {
				randTime := rand.Intn(35)

				if randTime < 10 {
					continue
				}

				time.Sleep(time.Duration(randTime) * time.Second)
				fmt.Printf("%d seconds has passed...", randTime)
				PlaySound()
				fmt.Printf("Wait...Did you hear that??\n")
			}
		}()

		return nil
	},
}

func PlaySound() {
	cmd := exec.Command("aplay", "./internal/risada-do-kiko.wav")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to open link: %v\n", err)
	}
}
