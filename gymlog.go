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
	fmt.Println("\nCurrent Display: ", dateOps.DisplayDate(), "\n")

	// TEST FileOps:
	fmt.Println("Filename = ", dataStore.SetName(INIT, "newInit"))
	fmt.Println("Filename = ", dataStore.SetName(DATA, "newDataFile.txt"))

	fmt.Println("Filename = ", dataStore.Name(INIT))
	fmt.Println("Filename = ", dataStore.Name(DATA))
	fmt.Println("Filename = ", dataStore.Name(PAGE))
	fmt.Println("Filename = ", dataStore.Name(LIST))
	fmt.Println("")

	tst1 := []string{"FW Curl", "4", "10", "20"}
	tst2 := []string{"Leg Press", "4", "10", "220"}
	tst3 := []string{"LAT", "4", "10", "90"}

	dataStore.InitStore()
	dataStore.SetEntry("FWC", tst1)
	dataStore.SetEntry("LP", tst2)
	dataStore.SetEntry("LAT", tst3)

	fmt.Println("Entry = ", dataStore.Entry("FWC"))
	fmt.Println("Entry = ", dataStore.Entry("LP"))
	fmt.Println("Entry = ", dataStore.Entry("LAT"))

	fmt.Println("Codes = ", dataStore.Codes())
	dataStore.RemEntry("LP")
	fmt.Println("Codes = ", dataStore.Codes())

	dataStore.SetEntry("LP", tst2)
	codes := dataStore.Codes()
	fileOps.WriteFile("test.bin", codes)
	retTest := fileOps.ReadFile("test.bin")
	fmt.Println("RetTest = ", retTest)

	fmt.Println("\n---------- INIT FILE DATA -------------------")
	//	fmt.Println("INI = ", fileOps.ReadFile(dataStore.Name(INIT)))

	elst := dataStore.LoadInit()
	fmt.Println("INI = \n", elst)

	fmt.Println("Page Builder:")
	page := builder.BuildPage()
	for _, v := range page {
		fmt.Print(v)
	}

	fmt.Println("\n-------- END OF TEST FUNCTIONS ------------\n")
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
