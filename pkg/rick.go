package wildcards

import (
	"context"
	"fmt"
	"os/exec"
)

var RickRollFucker = WildCard{
	Level: LEVEL_EASY,
	Description: "Shows GTA VII (7, stupid) trailer.",
	execute: func(ctx context.Context) error {
		OpenLink()

		return nil
	},
}

func OpenLink() {
	cmd := exec.Command("xdg-open", "https://www.youtube.com/watch?v=dQw4w9WgXcQ&list=RDdQw4w9WgXcQ&start_radio=1")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to open link: %v\n", err)
	}
}
