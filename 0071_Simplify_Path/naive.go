package main

import (
	"fmt"
	"strings"
	"github.com/ZhengjunHUO/godtype"
)

func simplifyPath(path string) string {
	deque := godtype.NewDeque()
	dirs := strings.Split(path, "/")

	for _,v := range dirs {
		switch v {
			case ".", "":
				// skip	
			case "..":
				deque.PopLast()
			default:
				deque.PushLast(v)			
		}
	}

	rslt := ""
	for !deque.IsEmpty(){
		rslt += "/"
		rslt += deque.PopFirst().(string)
	}

	if rslt == "" {
		return "/"
	}

	return rslt
}

func main() {
	strs := []string{"/home/", "/../", "/home//foo/", "/a/./b/../../c/"}
	for _, v := range strs {
		fmt.Println(simplifyPath(v))
	}
}
