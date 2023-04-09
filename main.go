package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func rgb(r, g, b int) string {
	return fmt.Sprintf("%06X", r<<16|g<<8|b)
}

func main() {
	app := &cli.App{
		Name:  "RGB to Hex",
		Usage: "Convert RGB values to a 6-digit hexadecimal string",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "r",
				Usage:    "Red value (0-255)",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "g",
				Usage:    "Green value (0-255)",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "b",
				Usage:    "Blue value (0-255)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			r := c.Int("r")
			g := c.Int("g")
			b := c.Int("b")

			if 0 <= r && r <= 255 && 0 <= g && g <= 255 && 0 <= b && b <= 255 {
				hexValue := rgb(r, g, b)
				fmt.Printf("Hex value: #%s\n", hexValue)
			} else {
				return errors.New("Error: RGB values must be between 0 and 255")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
