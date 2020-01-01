package dateOps

import (
	"fmt"
	"time"
)

//
// Start and End Date Variable
var s_mm, s_dd, s_yy int
var e_mm, e_dd, e_yy int

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

func validate(mm, dd, yy int) bool {
	t := time.Now()

	for {
		if mm < 1 && mm > 12 {
			break
		}
		if dd < 1 && dd > 32 {
			break
		}
		if mm < 1900 && mm > t.Year() {
			break
		}
		return true
	}
	fmt.Println("Malformed Date Value:  mm/dd/yy!")
	return false

}

func SetStart(mm, dd, yy int) bool {
	if validate(mm, dd, yy) {
		s_mm = mm
		s_dd = dd
		s_yy = yy
		return true
	}
	return false
}

func SetEnd(mm, dd, yy int) bool {
	if validate(mm, dd, yy) {
		e_mm = mm
		e_dd = dd
		e_yy = yy
		return true
	}
	return false
}

func InRange(mm, dd, yy int) bool {
	if !validate(mm, dd, yy) {
		fmt.Println("Validate Failed")
		return false
	}
	input := fmt.Sprintf("%d-%d-%d", yy, mm, dd)
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, input)
	fmt.Println(t1)

	return true
}
