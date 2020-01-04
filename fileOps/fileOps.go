package fileOps

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type dataRecord map[string]map[string][]int

//
// Database File Functions
//
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func PrintItem(item []string) {
	for i, v := range item {
		fmt.Println("[", i, "] ", v)
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

func storeDatabase(db interface{}, file string) {
	js := ToJson(db) //Marshal
	err := ioutil.WriteFile(file, js, 0644)
	check(err)
}

func loadDatabase(dataFile string) []string {
	ldb, err := ioutil.ReadFile(dataFile)
	if err != nil { // Assume No File Present
		storeDatabase("", "test")
		ldb, err = ioutil.ReadFile(dataFile)
		check(err)
	}
	check(err)
	s := string(ldb)
	s = strings.Replace(s, ",,", ",", -1)
	db := strings.Split(s, "\n")
	odb := make([]string, 1)
	for i, v := range db {
		fmt.Println("[", i, "]  ", v)
	}
	return odb
}

//
// Json Fuctions for MarshalIndent, Marshal Unmarshal
//
func ToJsonIndent(i interface{}) []byte {
	data, err := json.MarshalIndent(i, "", "   ")
	if err != nil {
		panic(err)
	}

	return data

}

func ToJson(i interface{}) []byte {
	data, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return data

}

func FromJson(v []byte, vv interface{}) {
	json.Unmarshal(v, &vv)
}

func LoadDatabase(file string) (db []string) {
	data := ReadFile(file)
	//	fmt.Println(data)
	return data
}
