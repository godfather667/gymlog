// Gym Log Program
// This program keep track of daily gym records and provides
// analysis of the records and builds daily workout sheets
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gym_project/gymlog/dataStore"
	_ "gym_project/gymlog/fileOps"

	"github.com/urfave/cli"
)

//
// Type Definitions
//
type FileType int
type FileName string

//
// Constants
const (
	INIT = iota // File Specifers
	DATA
	PAGE
	LIST
)

type Dtype struct {
	Dt int
}

//
// Variable Definitions
//

//
// Functions
//

//
// Display Current Date and return in readable format (ASCII)
//
func displayDate() (curDate string) {
	t := time.Now()
	curDate = fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())
	fmt.Println("Current Date:  ", curDate)
	return curDate
}

//
// Main Function
//

func main() {
	// Display Current Date
	displayDate()

	t := dataStore.Dtype{Dt: INIT}

	//	fmt.Println("Console Input = ", fileOps.Console("Input: "))
	/*
		fmt.Println("Filename = ", dataStore.FileName(INIT))
		fmt.Println("Filename = ", dataStore.FileName(DATA))
		fmt.Println("Filename = ", dataStore.FileName(PAGE))
		fmt.Println("Filename = ", dataStore.FileName(LIST))
	*/
	fmt.Println("Filename = ", dataStore.FileName(INIT))

	fmt.Println("Filename = ", t.Name())
	//
	// CLI Front End
	//
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "Page Storage Commands: ",
			Subcommands: []*cli.Command{
				{
					Name:    "page",
					Aliases: []string{"p"},
					Usage:   "Writes Page File",
					Action: func(c *cli.Context) error {
						/*
						 */
						return nil
					},
				},
				{
					Name:    "data",
					Aliases: []string{"d"},
					Usage:   "Writes Page Data to Data File",
					Action: func(c *cli.Context) error {
						fmt.Println("write data file: ", c.Args().First())
						/*
						 */
						return nil
					},
				},
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List database commands:",
			Subcommands: []*cli.Command{
				{
					Name:    "short_dates",
					Usage:   "mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed",
					Aliases: []string{"s"},
					Action: func(c *cli.Context) error {
						fmt.Println("short dates: ", c.Args().First())
						/*
						 */
						return nil
					},
				},
				{
					Name:    "long_dates",
					Usage:   "mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed",
					Aliases: []string{"l"},
					Action: func(c *cli.Context) error {
						fmt.Println("long dates: ", c.Args().First())
						/*
						 */
						return nil
					},
				},
			},
		},
		{
			Name:    "report",
			Aliases: []string{"r"},
			Usage:   "mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed",
			Action: func(c *cli.Context) error {
				fmt.Println("Report: ", c.Args().First())
				/*
				 */
				return nil
			},
		},
		{
			Name:    "chart",
			Aliases: []string{"c"},
			Usage:   "mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed",
			Action: func(c *cli.Context) error {
				fmt.Println("chart: ", c.Args().First())
				/*
				 */
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
