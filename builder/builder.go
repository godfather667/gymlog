package builder

import (
	"fmt"
	"gym_project/gymlog/dataStore"
	"gym_project/gymlog/dateOps"
	"gym_project/gymlog/extract"
	"gym_project/gymlog/fileOps"

	"strconv"
	"strings"
)

//
// Constants
//
const (
	INIT = iota // File Specifers
	DATA
	PAGE
	LIST
)

func pad(frag string, width int) (cnt int) {
	l := len(frag)
	cnt = width - l
	return cnt
}

func BreakEx(ex string) (cycle, reps, weight int) {
	cycle, err := strconv.Atoi(extract.Between(ex, "(", "x"))
	if err != nil {
		fmt.Println("Malformed exercise specification:  example: (4x10@30) Acutal: ", ex)
		return 0, 0, 0
	}
	reps, err = strconv.Atoi(extract.Between(ex, "x", ")"))
	if err == nil {
		return cycle, reps, 0 // Form = "(5X200)
	} else {
		if err != nil {
			reps, err = strconv.Atoi(extract.Between(ex, "x", "@"))
			if err != nil {
				fmt.Println("Malformed exercise specification:  example: (4x10@30) Acutal: ", ex)
				return 0, 0, 0
			}
			weight, err = strconv.Atoi(extract.Between(ex, "@", ")"))
			if err != nil {
				fmt.Println("Malformed exercise specification:  example: (4x10@30) Acutal: ", ex)
				return 0, 0, 0
			}
		}
	}

	return cycle, reps, weight
}

func BuildPage(title bool) []string {
	page := make([]string, 0)

	if title {
		page = append(page, "MON           WED             FRI")
		page = append(page, "==================================\n")
	}

	for _, v := range dataStore.LoadInit() {
		str := ""
		if strings.Contains(v, "MDY") {
			v = fmt.Sprintln("MDY Create_Date ", dateOps.PageDate())
		}

		result := strings.Fields(v)
		if len(v) < 2 {
			continue
		}
		str = ""
		str = str + result[0]
		for i := 0; i < pad(result[0], 4); i++ {
			if len(result[0]) > 2 {
				str = str + "   "
			} else {
				str = str + "  "
			}
		}
		str += result[1]
		for i := 0; i < pad(result[1], 12); i++ {
			str = str + " "
		}
		str += "  "

		str += result[2]
		if !strings.Contains(result[0], "MDY") {
			if len(result) == 4 {
				str = str + "  " + result[3]
			}
		}
		str += "\n"
		page = append(page, str)
	}
	return page
}

func BuildRecord() (dataRecord string) {
	page := BuildPage(false)
	fmt.Println("PAGE = ", page)
	date := dateOps.PageDate()
	codeRecord := date + "  "

	for _, v := range page {
		if len(v) < 2 {
			continue
		}
		if strings.Contains(v, "MDY") {
			//			date := dateOps.DisplayDate()
			continue
		}

		result := strings.Fields(v)
		if len(result) != 3 && len(result) != 4 {
			fmt.Println("Malformed Record: Improper number of Fields!  Fields = ", len(result))
			continue
		}
		c, r, w := BreakEx(result[2])
		if len(result) == 4 {
			c3, r3, w3 := BreakEx(result[3])
			if c3 > c {
				c, r, w = c3, r3, w3
			}
		}
		codeRecord += " , " + result[0] + " "
		codeRecord += strconv.Itoa(c) + " "
		codeRecord += strconv.Itoa(r) + " "
		codeRecord += strconv.Itoa(w)
	}
	codeRecord = codeRecord[:len(codeRecord)-1]
	codeRecord += "\n;"
	fmt.Println("Data Record = ", codeRecord)
	return codeRecord
}

func RebuildDatabase(rn int) {
	newDat := string(fileOps.ReadData(dataStore.Name(DATA)))
	lines := strings.Split(newDat, ";")

	ll := 0
	for i, v := range lines {
		if len(v) < 2 {
			ll = i
		}
	}

	rn -= 1
	if rn >= ll {
		fmt.Println("Err - Line Number does not Exist")
		return
	}

	ol := ""
	for i := 0; i < rn; i++ {
		ol += lines[i] + ";"
	}
	for j := rn + 1; j < ll; j++ {
		ol += lines[j] + ";"
	}
	b := []byte(ol)
	fileOps.WriteData(dataStore.Name(DATA), b)
	return
}

func BuildChart() {
	newDat := string(fileOps.ReadData(dataStore.Name(DATA)))
	lines := strings.Split(newDat, ";")

	ex := strings.Split(lines[0], ",")
	lex := len(ex)
	fmt.Println("lex = ", lex)

	for i, v := range ex {
		fmt.Println("[", i, "]  ", v)
	}
	date := dateOps.ConvertDate(ex[0])
	fmt.Println(date)

	return
}

/*

		ll := 0
		for i, v := range lines {
			if len(v) < 2 {
				ll = i
			}
		}

		rn -= 1
		if rn >= ll {
			fmt.Println("Err - Line Number does not Exist")
			return
		}

		ol := ""
		for i := 0; i < rn; i++ {
			ol += lines[i] + ";"
		}
		for j := rn + 1; j < ll; j++ {
			ol += lines[j] + ";"
		}
		b := []byte(ol)
		fileOps.WriteData(dataStore.Name(DATA), b)


	return
}
*/
