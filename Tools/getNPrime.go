package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
	"os"
	"strconv"
	"log"
)

func main() {
	n := 10
	if len(os.Args) == 2 {
		x, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("A number is needed !")
		}
		n = x
        }

	rslt := gtoolkit.SieveOfEratosthenes(n*n)
	fmt.Println(rslt[:n])
}
