package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println("Printing random Emojis")
	emojis := map[int]string{
		1: "baby",
		2: "child",
		3: "cow",
		4: "girl",
		5: "man",
		6: "unicorn",
		7: "ox",
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 3; i++ {
		key := r1.Intn(7) + 1
		emoji.Println(":", emojis[key], ": ", emojis[key], "!!!")
	}
}
