package wildcards

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

var RickRollFucker = WildCard{
	Level: LEVEL_EASY,
	Description: "Shows GTA VII (7) trailer at random moments.",
	execute: func(ctx context.Context) error {
		go func() {
			OpenLink()

			for true {
				randTime := rand.Intn(30)

				if randTime < 10 {
					continue
				}

				time.Sleep(time.Duration(randTime) * time.Second)
				fmt.Printf("%d seconds has passed...", randTime)
				fmt.Printf("Let's see the trailer again!!")
				OpenLink()
			}
		}()

		return nil
	},
}

func OpenLink() {
	cmd := exec.Command("xdg-open", "https://www.youtube.com/watch?v=dQw4w9WgXcQ&list=RDdQw4w9WgXcQ&start_radio=1")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to open link: %v\n", err)
	}
}
