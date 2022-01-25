package main

import (
	"strings"
)

func capitalizeTitle(title string) string {
	ss := strings.Split(title, " ")
	rslt := ""
	for i := range ss {
		if len(ss[i]) < 3 {
			rslt += strings.ToLower(ss[i])
		}else{ 
			temp := strings.ToLower(ss[i])
			rslt += string(temp[0] - 32) + temp[1:]
		}

		if i < len(ss) - 1 {
			rslt += " "
		}
	}

	return rslt
}
