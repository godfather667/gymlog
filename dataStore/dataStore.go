package dataStore

import (
	"fmt"
	"gym_project/gymlog/fileOps"
	"strings"
)

type Ftype int
type Fname string

type MyDt int

type dataRecord map[string]map[string][]int

const (
	INIT = iota
	DATA
	PAGE
	LIST
	DATE
)

//
// Data Base Structures and Variables
//

var dataDB dataRecord

var Store map[string][]string

var CodeList []string

var initFile string = "gymlog.ini"
var dataFile string = "dataFile"
var pageFile string = "pageFile.txt"
var listFile string = "listFile.txt"
var forceDate string = ""

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
	case DATE:
		return forceDate
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
	case DATE:
		forceDate = n
		return true
	default:
		fmt.Println("Error: Filetype Unknown: Allowed: DATA, PAGE, LIST")
		return false
	}
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
