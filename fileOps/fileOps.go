package fileOps

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//
// Database File Functions
//
func check(e error) {
	if e != nil {
		panic(e)
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
	data, err := json.MarshalIndent(i, "", "   ")
	if err != nil {
		panic(err)
	}

	return data

}

func FromJson(v []byte, vv interface{}) {
	json.Unmarshal(v, &vv)
}

/*

func readInit() (sin stringItem) {
	sin = make(stringItem, 0)
	file, err := os.Open(initFile)
	if err != nil {
		fmt.Println("File Open Err: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "INIT ") {
			initFile = strings.TrimLeft(line, "INIT ")
			continue
		}
		if strings.HasPrefix(line, "PAGE ") {
			pageFile = strings.TrimLeft(line, "PAGE ")
			continue
		}
		if strings.HasPrefix(line, "DBASE ") {
			dataFile = strings.TrimLeft(line, "DBASE ")
			continue
		}
		if strings.HasPrefix(line, "LIST ") {
			listFile = strings.TrimLeft(line, "LIST ")
			continue
		}
		comment := strings.HasPrefix(line, "#") || strings.HasPrefix(line, " #")
		if !comment && len(line) > 2 {
			sin = append(sin, line)
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		fmt.Println("SIN: ", sin)
	}
	return sin
}
func readData() []string {
	data := make([]string, 0)
	file, err := os.Open(dataFile)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func writePage(data []string) {
	_ = os.Remove(pageFile)
	file, err := os.OpenFile(pageFile, os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()

	datawriter := bufio.NewWriter(file)

	for _, line := range data {
		_, _ = datawriter.WriteString(line + "\n")
	}
	datawriter.Flush()
}

func writeInit(data []string) {
	_ = os.Remove(pageFile)
	file, err := os.OpenFile(initFile, os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()

	datawriter := bufio.NewWriter(file)

	for _, line := range data {
		_, _ = datawriter.WriteString(line + "\n")
	}
	datawriter.Flush()
}
func writeData(data string) {
	file, err := os.OpenFile(dataFile, os.O_CREATE|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()

	datawriter := bufio.NewWriter(file)

	_, _ = datawriter.WriteString(data)

	datawriter.Flush()
}

func loadDatabase() []string {
	ldb, err := ioutil.ReadFile(dataFile)
	if err != nil { // Assume No File Present
		writeData("")
		ldb, err = ioutil.ReadFile(dataFile)
		check(err)
	}
	check(err)
	s := string(ldb)
	s = strings.Replace(s, ",,", ",", -1)
	db := strings.Split(s, "\n")
	odb := make([]string, 1)
	for _, v := range db {
		getRecordDate(v) // Get Record Date
		if !InRange() {
			continue
		}
	}
	return odb
}

func storeDatabase(db []interface{}, fileName string) {
	js := toJson(db) //Marshal

	err := ioutil.WriteFile(dataFile, js, 0644)
	check(err)
}


*/
