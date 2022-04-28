package algorithm

import (
	"regexp"
	"strings"
	"fmt"
)

//Mengecek apakah input sesuai DNA
func IsSanitizedRegex(text string) bool {
	match1, err := regexp.MatchString("^[ATCG]+$", text)
	if err == nil {
		return match1
	}
	return false
}

//Mengecek text dengan format regex yang disediakan dan fungsi StringRegex
func ReadRegex(text string) SearchRegex {
	var regexformats [11]RegexFormat
	regexformats[0] = RegexFormat{"\\d{2}/\\d{2}/\\d{4}", "/"}
	regexformats[1] = RegexFormat{"\\d{2}-\\d{2}-\\d{4}", "-"}
	regexformats[2] = RegexFormat{"\\d{2}[ \t]+\\d{2}[ \t]+\\d{4}", " "}

	regexformats[3] = RegexFormat{"\\d{2}/.+/\\d{4}", "/"}
	regexformats[4] = RegexFormat{"\\d{2}-.+-\\d{4}", "-"}
	regexformats[5] = RegexFormat{"\\d{2}[ \t].+[ \t]+\\d{4}", " "}

	regexformats[6] = RegexFormat{".+/\\d{4}", "/"}
	regexformats[7] = RegexFormat{".+-\\d{4}", "-"}
	regexformats[8] = RegexFormat{".+[ \t]+\\d{4}", " "}

	regexformats[9] = RegexFormat{"^\\d{4}", "/$"}

	regexformats[10] = RegexFormat{".+", " "}

	var date SearchRegex
	var get bool
	for _, v := range regexformats {
		date, get = StringRegex(text, v.regex, v.splitter)
		fmt.Println(v.regex)
		if get {
			return date
		}

	}
	date = SearchRegex{"null", "null", "null", "null"}
	return date

}
//Memecah informasi text dengan regex kedalam SearchRegex yang lalu di-return bersama bool yang menyatakan apakah regex valid.
func StringRegex(text, regex, splitter string) (SearchRegex, bool) {
	regexcompiled := regexp.MustCompile(regex)

	if regexcompiled.MatchString(text) {
		fmt.Println("TEXT :",text )
		date := regexcompiled.FindString(text)
		res := strings.TrimSpace(regexcompiled.ReplaceAllString(text, "${1} $2"))

		datesplitted := strings.Split(date, splitter)
		fmt.Println("DATESPLITTER = ",len(datesplitted))
		if splitter == " " {
			datesplitted = strings.Fields(date)
		}
		var day string
		var month string
		var year string
		if len(datesplitted) == 3 {
			day = datesplitted[0]
			month = datesplitted[1]
			year = datesplitted[2]
		} else if len(datesplitted) == 2 {
			day = "null"
			month = datesplitted[0]
			year = datesplitted[1]
		} else {
			day = "null"
			regexyear := regexp.MustCompile("\\d{4}")
			if regexyear.MatchString(datesplitted[0]) {
				month = "null"
				year = datesplitted[0]
			} else if(checkifDate(datesplitted[0])){
				month = datesplitted[0]
				year = "null"
			}else{
				month = "null"
				year = "null"
				res = strings.TrimSpace(datesplitted[0])
			}
		}
		if res ==""{
			res = "null"
		}
		fmt.Println("TEXT :",text )
		var x = SearchRegex{day, month, year, res}
		return x, true
	} else {
		var x = SearchRegex{"null", "null", "null", "null"}
		return x, false
	}
}
