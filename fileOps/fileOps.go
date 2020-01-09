package fileOps

import (
	"bufio"
	"fmt"
	"gym_project/gymlog/dateOps"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//
// Database File Functions
//
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Console(line string) (response string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(line)
	response, err := reader.ReadString('\n')
	check(err)
	return
}

func ReadFile(file string) []string {
	content, err := ioutil.ReadFile(file)
	check(err)
	//	text := string(content)
	return strings.Split(string(content), "\n")
}

func ReadData(file string) []byte {
	content, err := ioutil.ReadFile(file)
	check(err)
	return content
}

func WriteData(file string, data []byte) {
	err := ioutil.WriteFile(file, data, 0644)
	check(err)
}

func WriteFile(file string, data []string) {
	content := ""
	for _, v := range data {
		content += v + "\n"
	}
	bc := []byte(content)
	err := ioutil.WriteFile(file, bc, 0644)
	check(err)
}

func WriteAppend(dataFile string, data []byte) {
	file, err := os.OpenFile(dataFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()

	datawriter := bufio.NewWriter(file)

	_, _ = datawriter.Write(data)

	datawriter.Flush()
}

func LoadDatabase(file string) (db string) {
	db = ""
	mm, dd, yy := 0, 0, 0
	data := string(ReadData(file))
	lines := strings.Split(data, ";")
	for _, v := range lines {
		item := strings.Split(v, ",")
		if len(item) > 1 {
			mm, dd, yy = dateOps.BreakDate(item[0])
			if dateOps.InRange(mm, dd, yy) {
				db += v + ";"
			}
		}
	}
	db = db[:len(db)-1]
	return db
}
