// Gym Log Program
// This program keep track of daily gym records and provides
// analysis of the records and builds daily workout sheets
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
// Helper Function
//

//
// Main Function
//

func main() {

	dateOps.DateStart(0, 0, 0)
	dateOps.DateEnd(0, 0, 0)

	//
	// CLI Front End
	//
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "\n   page",
			Aliases: []string{"p"},
			Usage:   "  Prints Page for Log Book: (Formats gymlog.ini for Log Book)\n",
			Action: func(c *cli.Context) error {
				page := builder.BuildPage(true)
				fileOps.WriteFile("pageFile.txt", page)
				return nil
			},
		},
		{
			Name:    "data",
			Aliases: []string{"d"},
			Usage:   "  Store Page in Database: (Database Format)\n",
			Action: func(c *cli.Context) error {
				mapD := builder.BuildRecord()
				//				b := []byte(mapD)
				fileOps.WriteAppend(dataStore.Name(DATA), []byte(mapD))
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "  List Contents of Database by range:\n              list mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed\n",
			Action: func(c *cli.Context) error {
				dateOps.LoadCmdDate(c)
				newDat := fileOps.LoadDatabase(dataStore.Name(DATA))

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
				recNumber := c.Args().First()
				rn, err := strconv.Atoi(recNumber)
				if err != nil {
					fmt.Println("Err - Improper Number - Request Ignored!")
					return nil
				}
				builder.RebuildDatabase(rn)
				return nil
			},
		},
		{
			Name:    " chart",
			Aliases: []string{"c"},
			Usage:   "Produces Progress Chart:\n              chart mm dd yyyy mm dd yyyy -No dates = all dates, otherwise range is processed\n",
			Action: func(c *cli.Context) error {
				fmt.Println("\n")
				dateOps.LoadCmdDate(c)

				newDat := fileOps.LoadDatabase(dataStore.Name(DATA))
				lines := strings.Split(newDat, ";")

				for _, v := range lines {
					if newLine, ok := builder.BuildLine(v); ok {
						fmt.Println(newLine)
						fmt.Println(builder.Spacer)
					}
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
