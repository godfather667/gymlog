package builder

import (
	"fmt"
	"gym_project/gymlog/dataStore"
	"gym_project/gymlog/dateOps"
	"gym_project/gymlog/extract"
	"strconv"
	"strings"
)

func pad(frag string, width int) (cnt int) {
	l := len(frag)
	cnt = width - l
	return cnt
}

func BreakEx(ex string) (cycle, reps, weight int) {
	cycle, err := strconv.Atoi(extract.Between(ex, "(", "x"))
	if err != nil {
		fmt.Println("Malformed exercise specification:  example: (4x10x30) Acutal: ", ex)
		return 0, 0, 0
	}
	reps, err = strconv.Atoi(extract.Between(ex, "x", ")"))
	if err == nil {
		return cycle, reps, 0 // Form = "(5X200)
	} else {
		if err != nil {
			reps, err = strconv.Atoi(extract.Between(ex, "x", "@"))
			if err != nil {
				fmt.Println("Malformed exercise specification:  example: (4x10x30) Acutal: ", ex)
				return 0, 0, 0
			}
			weight, err = strconv.Atoi(extract.Between(ex, "@", ")"))
			if err != nil {
				fmt.Println("Malformed exercise specification:  example: (4x10x30) Acutal: ", ex)
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

		//BreakEx(result[2])

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

//func BuildRecord() (dataRecord map[string]map[string][]string) {
func BuildRecord() {
	page := BuildPage(false)
	//	codeRecord := make(map[string][]string)

	for _, v := range page {
		if len(v) < 2 {
			continue
		}
		if strings.Contains(v, "MDY") {
			//			date := dateOps.DisplayDate()
			continue
		}
		if strings.Contains(v, "HT") ||
			strings.Contains(v, "PW") ||
			strings.Contains(v, "PC") {
			continue
		}

		fmt.Println(v)
	}
	return
}
