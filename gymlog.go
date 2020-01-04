// Gym Log Program
// This program keep track of daily gym records and provides
// analysis of the records and builds daily workout sheets
package main

import (
	"encoding/json"
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

type dataRecord map[string]map[string][]int
type codeRecord map[string][]int

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

	/*
		dataStore.DateEnd(1, 3, 2020)
		dataStore.DateStart(11, 1, 2019)

		smm, sdd, syy := dataStore.GetDateStart()
		emm, edd, eyy := dataStore.GetDateEnd()

		fmt.Println("Start:  ", smm, "  ", sdd, "  ", syy)
		fmt.Println("End:    ", emm, "  ", edd, "  ", eyy)
	*/

	/*
		dateOps.SetEnd(1, 3, 2020)
		dateOps.SetStart(11, 11, 2019)
		fmt.Println(dateOps.InRange(10, 27, 2019))
	*/

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
						mapD := builder.BuildRecord()
						js, _ := json.Marshal(mapD)
						fmt.Println("mapD = ", mapD)
						fileOps.WriteAppend(dataStore.Name(DATA), js)
						//						fmt.Println("JSON: ", string(json))
						return nil
					},
				},
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed",
			Action: func(c *cli.Context) error {
				fmt.Println("Database Listing: ", c.Args().First())
				//				data := fileOps.LoadDatabase(dataStore.Name(DATA))
				//				fmt.Println("DATA RECORD = ", data)
				dat := codeRecord{}
				newDat := fileOps.ReadData(dataStore.Name(DATA))
				//				if err := json.Unmarshal(newDat, &dat); err != nil {
				//					panic(err)
				//				}
				fmt.Println("newData = ", newDat)
				json.Unmarshal(newDat, &dat)
				fmt.Println("Dat = ", dat)
				for i, v := range dat {
					fmt.Println("[", i, "] ", v)
				}
				return nil
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
