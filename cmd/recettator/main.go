package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/camembertaulaitcrew/recettator"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "recettator"
	app.Usage = "Generate CALC recipes"
	app.Version = "master"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
		cli.IntFlag{
			Name:  "seed, s",
			Usage: "Set seed value",
		},
		cli.IntFlag{
			Name:  "ingredients, i",
			Usage: "Amount of ingredients",
		},
		cli.IntFlag{
			Name:  "steps",
			Usage: "Amount of steps",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Use JSON output",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("%v", err)
	}
}

func run(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	seed := int64(c.Int("seed"))
	if seed == 0 {
		seed = time.Now().UTC().UnixNano()
	}
	rctt := recettator.New(seed)

	for i := 0; i < c.Int("ingredients"); i++ {
		rctt.AddRandomIngredient()
	}

	for i := 0; i < c.Int("steps"); i++ {
		rctt.AddRandomStep()
	}

	var output string
	var err error

	if c.Bool("json") {
		output = rctt.JSON()
	} else {
		output, err = rctt.Markdown()
		if err != nil {
			return err
		}
	}

	fmt.Println(output)

	return nil
}
