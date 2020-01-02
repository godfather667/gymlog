package dataStore

import (
	"fmt"
	"gym_project/gymlog/fileOps"
	"strings"
)

type Ftype int
type Fname string

type MyDt int

const (
	INIT = iota
	DATA
	PAGE
	LIST
)

//
// Data Base Structures and Variables
//

var Store map[string][]string
var CodeList []string

var initFile string = "gymlog.ini"
var dataFile string = "dataFile"
var pageFile string = "pageFile.txt"
var listFile string = "listFile.txt"

func Name(d MyDt) string {
	switch d {
	case INIT:
		return initFile
	case DATA:
		return dataFile
	case PAGE:
		return pageFile
	case LIST:
		return listFile
	default:
		fmt.Println("Error: Filetype Unknown: Allowed: INIT, DATA, PAGE, LIST")
		return ""
	}
}

func SetName(d MyDt, n string) bool {
	switch d {
	case INIT:
		fmt.Println("Error: Init Filename can not be changed!")
		return false
	case DATA:
		dataFile = n
		return true
	case PAGE:
		pageFile = n
		return true
	case LIST:
		listFile = n
		return true
	default:
		fmt.Println("Error: Filetype Unknown: Allowed: DATA, PAGE, LIST")
		return false
	}
}

func InitStore() {
	Store = make(map[string][]string)
	CodeList = make([]string, 1)
}

func SetEntry(code string, entry []string) {
	Store[code] = entry
}

func Entry(code string) []string {
	return Store[code]
}

func RemEntry(code string) {
	delete(Store, code)
}

func Codes() []string {
	CodeList = nil
	for c, _ := range Store {
		CodeList = append(CodeList, c)
	}
	return CodeList
}

func LoadInit() (elst []string) {
	content := fileOps.ReadFile(initFile)

	elst = make([]string, 1)
	for _, line := range content {
		if strings.HasPrefix(line, "PAGE ") {
			pageFile = strings.TrimLeft(line, "#PAGE ")
			continue
		}
		if strings.HasPrefix(line, "DBASE ") {
			dataFile = strings.TrimLeft(line, "#DATA ")
			continue
		}
		if strings.HasPrefix(line, "LIST ") {
			listFile = strings.TrimLeft(line, "#LIST ")
			continue
		}
		comment := strings.HasPrefix(line, "#") || strings.HasPrefix(line, " #")
		if !comment && len(line) > 2 {
			elst = append(elst, line)
		}
	}
	return elst
}
