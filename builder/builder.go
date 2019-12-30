package builder

func BuildPage() []string {
	page := make([]string, 1)
	//	blanks := "                                                                                "

	page = append(page, "MON		WED		     FRI\n")
	page = append(page, "==================================\n")
	return page
}
