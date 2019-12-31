package builder

import (
	"fmt"
	"gym_project/gymlog/dataStore"
	"gym_project/gymlog/dateOps"
	"gym_project/gymlog/fileOps"
	"strings"
)

func pad(frag string, width int) (cnt int) {
	l := len(frag)
	cnt = width - l
	return cnt
}

func BuildPage() []string {
	page := make([]string, 0)

	page = append(page, "MON           WED             FRI")
	page = append(page, "==================================\n")

	for i, v := range dataStore.LoadInit() {
		str := ""
		if strings.Contains(v, "MDY") {
			v = fmt.Sprintln("MDY Create_Date ", dateOps.PageDate())
		}

		result := strings.Fields(v)
		if len(result) < 3 || len(result) > 4 {
			fmt.Println("Line [", i, "]: ", v, " improperly formated!  Will be Ignored!")
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
	fileOps.WriteFile("pageFileTest.txt", page)
	return page
}
