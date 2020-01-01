// Gym Log Program
// This program keep track of daily gym records and provides
// analysis of the records and builds daily workout sheets
package main

import (
	"fmt"
	"log"
	"os"

	"gym_project/gymlog/builder"
	"gym_project/gymlog/dataStore"
	"gym_project/gymlog/dateOps"
	_ "gym_project/gymlog/extract"
	"gym_project/gymlog/fileOps"

	"github.com/urfave/cli"
)

//
// Type Definitions
//
type FileType int
type FileName string

type MyDt int

//
// Constants
//
const (
	INIT = iota // File Specifers
	DATA
	PAGE
	LIST
)

//
// Variables
//

//
// Main Function
//

func main() {
	// Display Current Date
	fmt.Println("\nCurrent Date: ", dateOps.DisplayDate(), "\n")

	fmt.Println(dateOps.InRange(12, 27, 2019))

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
						page := builder.BuildPage(true)
						fileOps.WriteFile("pageFile.txt", page)
						return nil
					},
				},
				{
					Name:    "data",
					Aliases: []string{"d"},
					Usage:   "Writes Page Data to Data File",
					Action: func(c *cli.Context) error {
						mp := (builder.BuildRecord())
						json := fileOps.ToJson(mp)
						fileOps.WriteJSON(dataStore.Name(DATA), json)
						//fmt.Println("JSON: ", string(json))
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
