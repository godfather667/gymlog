package dateOps

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var mm_s, dd_s, yy_s, mm_e, dd_e, yy_e int

func DateStart(mm, dd, yy int) bool {
	if Validate(mm, dd, yy) {
		mm_s = mm
		dd_s = dd
		yy_s = yy
		return true
	}
	return false
}

func DateEnd(mm, dd, yy int) bool {
	if Validate(mm, dd, yy) {
		mm_e = mm
		dd_e = dd
		yy_e = yy
		return true
	}
	return false
}

func GetDateStart() (r bool, mm, dd, yy int) {
	mm = mm_s
	dd = dd_s
	yy = yy_s
	if mm == 0 || dd == 0 || yy == 0 {
		return false, mm, dd, yy
	}
	return true, mm, dd, yy
}
func GetDateEnd() (r bool, mm, dd, yy int) {
	mm = mm_e
	dd = dd_e
	yy = yy_e
	if mm == 0 || dd == 0 || yy == 0 {
		return false, mm, dd, yy
	}
	return true, mm, dd, yy
}

func LoadCmdDate(c *cli.Context) {
	if c.NArg() > 2 {
		mm, _ := strconv.Atoi(c.Args().Get(0))
		dd, _ := strconv.Atoi(c.Args().Get(1))
		yy, _ := strconv.Atoi(c.Args().Get(2))
		if !Validate(mm, dd, yy) {
			return
		}
		DateStart(mm, dd, yy)

		if c.NArg() == 6 {
			mm, _ := strconv.Atoi(c.Args().Get(3))
			dd, _ := strconv.Atoi(c.Args().Get(4))
			yy, _ := strconv.Atoi(c.Args().Get(5))
			if Validate(mm, dd, yy) {
				DateEnd(mm, dd, yy)
			}
		}
	}

	return
}

//
// Convert Display Date (MMxDD@2020) into MM/DD/YY
//
func ConvertDate(dDate string) (oDate string) {
	oDate = strings.TrimLeft(dDate, "(")
	oDate = strings.Trim(oDate, " )")
	oDate = strings.Replace(oDate, "x", "/", 1)
	oDate = strings.Replace(oDate, "@", "/", 1)
	return oDate
}

func BreakDate(date string) (cycle, reps, weight int) {
	date = strings.TrimLeft(date, "(")
	date = strings.Trim(date, " )")
	date = strings.Replace(date, "x", " ", 1)
	date = strings.Replace(date, "@", " ", 1)
	b := strings.Fields(date)
	if len(b) == 2 {
		b = append(b, "0")
	}
	m1, err1 := strconv.Atoi(b[0])
	m2, err2 := strconv.Atoi(b[1])
	m3, err3 := strconv.Atoi(b[2])
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("BreakDate Failure")
	}
	return m1, m2, m3

}

//
// Display Current Date and return in readable format (ASCII)
//
func DisplayDate() (curDate string) {
	t := time.Now()
	return fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())
}

func PageDate() (result string) {
	t := time.Now()
	return fmt.Sprintf("(%dx%d@%d)", t.Month(), t.Day(), t.Year())
}

func Validate(mm, dd, yy int) bool {
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	d := fmt.Sprintf("%02d/%02d/%04d", dd, mm, yy)
	return re.MatchString(d)
}

func InRange(mm, dd, yy int) bool {
	if !Validate(mm, dd, yy) {
		fmt.Println("Validate Failed")
		return false
	}
	tm := time.Month(mm)
	target := time.Date(yy, tm, dd, 0, 0, 0, 0, time.UTC)

	r, s_mm, s_dd, s_yy := GetDateStart()
	if !r {
		return true // No ranging possible!
	}

	sm := time.Month(s_mm)
	start := time.Date(s_yy, sm, s_dd, 0, 0, 0, 0, time.UTC)
	if target.Before(start) {
		return false
	}

	r, e_mm, e_dd, e_yy := GetDateEnd()
	if !r {
		return true // No End Date -
	}

	em := time.Month(e_mm)
	end := time.Date(e_yy, em, e_dd, 0, 0, 0, 0, time.UTC)
	if target.After(end) {
		return false
	}
	return true
}
