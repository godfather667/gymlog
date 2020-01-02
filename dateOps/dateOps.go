package dateOps

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

const (
	DATE_1 = iota // File Specifers
	DATE_2
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

func getDateStart() (mm, dd, yy int) {
	mm = mm_s
	dd = dd_s
	yy = yy_s
	return mm, dd, yy
}

func getDateEnd() (mm, dd, yy int) {
	mm = mm_e
	dd = dd_e
	yy = yy_e
	return mm, dd, yy
}

func LoadRecordDate(dateSelect int, c *cli.Context) bool {
	if dateSelect == DATE_1 && c.NArg() > 2 {
		mm, _ := strconv.Atoi(c.Args().Get(0))
		dd, _ := strconv.Atoi(c.Args().Get(1))
		yy, _ := strconv.Atoi(c.Args().Get(2))
		if Validate(mm, dd, yy) {
			DateStart(mm, dd, yy)
			return true
		}
	}

	if dateSelect == DATE_2 && c.NArg() == 6 {
		mm, _ := strconv.Atoi(c.Args().Get(0))
		dd, _ := strconv.Atoi(c.Args().Get(1))
		yy, _ := strconv.Atoi(c.Args().Get(2))
		if Validate(mm, dd, yy) {
			DateStart(mm, dd, yy)
			return true
		}
	}
	return false
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
	input := fmt.Sprintf("%d-%02d-%02d", yy, mm, dd)
	input += "T15:04:05.000-07:00"
	t, _ := time.Parse("2006-01-02T15:04:05.000-07:00", input)

	s_yy, s_mm, s_dd := getDateStart()
	input_s := fmt.Sprintf("%d-%02d-%02d", s_yy, s_mm, s_dd)
	input_s += "T15:04:05.000-07:00"
	t_s, _ := time.Parse("2006-01-02T15:04:05.000-07:00", input_s)

	e_yy, e_mm, e_dd := getDateEnd()
	input_e := fmt.Sprintf("%d-%02d-%02d", e_yy, e_mm, e_dd)
	input_e += "T15:04:05.000-07:00"
	t_e, _ := time.Parse("2006-01-02T15:04:05.000-07:00", input_e)

	fmt.Println("t = ", t, "  t_s = ", t_s, "  t_e", t_e)

	if t.Before(t_s) {
		fmt.Println("Date before Start!")
		return false
	}
	if t.After(t_e) {
		fmt.Println("Date After End!")
		return false
	}

	return true
}
