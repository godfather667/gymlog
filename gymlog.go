// Gym Log Program
// This program keep track of daily gym records and provides
// analysis of the records and builds daily workout sheets
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
			Usage:   "Prints Page for Log Book: (Formats gymlog.ini for Log Book)\n",
			Action: func(c *cli.Context) error {
				page := builder.BuildPage(true)
				fileOps.WriteFile("pageFile.txt", page)
				return nil
			},
		},
		{
			Name:    "data",
			Aliases: []string{"d"},
			Usage:   "Store Page in Database: (Database Format)\n",
			Action: func(c *cli.Context) error {
				mapD := builder.BuildRecord()
				fmt.Println("mapD = ", mapD)
				b := []byte(mapD)
				fileOps.WriteAppend(dataStore.Name(DATA), b)
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List Contents of Database by range:\n            list mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed\n",
			Action: func(c *cli.Context) error {
				fmt.Println("Database Listing: ", c.Args().First())

				newDat := string(fileOps.ReadData(dataStore.Name(DATA)))

				lines := strings.Split(newDat, ";")
				for i, v := range lines {
					if len(v) > 1 {
						fmt.Println("[", i+1, "] ", v)
					}
				}
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove a record from Database:\n              Remove <record number> (see listing)\n",
			Action: func(c *cli.Context) error {
				fmt.Println("Remove Database Record: ", c.Args().First())

				return nil
			},
		},
		{
			Name:    "chart",
			Aliases: []string{"c"},
			Usage:   "Produces Progress Chart:\n             chart mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed\n",
			Action: func(c *cli.Context) error {
				fmt.Println("Report: ", c.Args().First())
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
