package main

import (
	"fmt"
	"github.com/evsyukovmv/piglatin"
	"os"
	"strings"
)

func main () {
	phrase := strings.Join(os.Args[1:], " ")
	fmt.Println(piglatin.PigLatin(phrase))
}
