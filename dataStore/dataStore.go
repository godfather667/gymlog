package dataStore

import (
	"fmt"
)

type Ftype int
type Fname string

type Dtype struct {
	Dt int
}

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

var initFile string = "gymlog.ini"
var dataFile string = "dataFile"
var pageFile string = "pageFile.txt"
var listFile string = "listFile.txt"

func (d MyDt) Name() string {
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
		fmt.Println("Error: Filetype Unknown: Allowed: DATA, PAGE, LIST")
		return ""
	}
}
