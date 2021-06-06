package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"log"
)

func sieveOfEratosthenes(n int) []int {
        isNotPrime := make([]bool, n+1)
        prime := []int{}
        pivot := int(math.Ceil(math.Sqrt(float64(n))))

        for i:=2; i<n+1; i++ {
                if isNotPrime[i] == false {
                        prime = append(prime, i)
                        if i<=pivot {
                                for j:=i*i; j<=n; j+=i {
                                        isNotPrime[j] = true
                                }
                        }
                }
        }

        return prime
}

func main() {
	n := 5
	if len(os.Args) == 2 {
		x, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("A number is needed !")
		}
		n = x
        }

	rslt := sieveOfEratosthenes(n*n)
	fmt.Println(rslt[:n])
	for i := range rslt[:n] {
		fmt.Printf("%d,", rslt[i])
	}
	fmt.Println()
}
