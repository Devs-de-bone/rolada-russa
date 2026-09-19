package wildcards

import (
	"context"
	"math/rand"
	"time"

	"fmt"
	"os/exec"

	"github.com/go-vgo/robotgo"
)

var (
	sx, sy = robotgo.GetScreenSize()
)

var MouseFuckerCard = WildCard{
	Level: LEVEL_EASY,
	Description: "Binds the mouse to the screen brightness control. Move left to decrease and right to increase brightness.",
	execute: func(ctx context.Context) error {
		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				default:
					x, _ := robotgo.Location()
					brightnessPerc := float32(x) / float32(sx) * 100.0
					setBrightness(brightnessPerc)
 				}

			}
		}()

		return nil
	},
}

func setBrightness(delta float32) {
	// Executes: brightnessctl set +5% or 5%-
	brightness := fmt.Sprintf("%v%", delta)
	cmd := exec.Command("brightnessctl", "set", brightness)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to run brightnessctl: %v\n", err)
	}
}



var MoveFuckerCard = WildCard{
	Level: LEVEL_HARD,
	Description: "Moves mouse to left for 2 seconds at random times",
	execute: func(ctx context.Context) error {

		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				default:

					count := 0
					t := time.Tick(1 * time.Second)
					for count < 2 {
						select {
						case <- t:
							count++
						default:
							robotgo.DragSmooth(10, sy/2)
						}
					}
				}
				duration := time.Duration(rand.Intn(10))
				time.Sleep(duration * time.Second)
			}
		}()
		return nil
	},

}
