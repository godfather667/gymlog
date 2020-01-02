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

var mm_s, dd_s, yy_s, mm_e, dd_e, yy_e int

func DateStart(mm, dd, yy int) {
	mm_s = mm
	dd_s = dd
	yy_s = yy
	return
}

func DateEnd(mm, dd, yy int) {
	mm_e = mm
	dd_e = dd
	yy_e = yy
	return
}

func GetDateStart() (mm, dd, yy int) {
	mm = mm_s
	dd = dd_s
	yy = yy_s
	return mm, dd, yy
}

func GetDateEnd() (mm, dd, yy int) {
	mm = mm_e
	dd = dd_e
	yy = yy_e
	return mm, dd, yy
}

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
