package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAXINT = ^int(0)
)

func klariguLudon(min, max, skip int) {
	fmt.Printf(`Bonvenon ĉe "Kaptu la muson"

La celo de la ludo estas diveni numeron de %d ĝis %d.
Vi povas provi tion unu post la alia.
Tamen atentu! La numero povas ĉu ĉiam resti la sama, ĉu supreniri per %d ĉiufoje, ĉu malsupreniri ĉiufoje per la samo.
Se la numero fariĝas pli alta ol %d, ĝi reiras al %d.
Se la numero fariĝas malpli alta ol %d, ĝi reiras al %d.
Multan sukceson!

`, min, max-1, skip, max-1, min, min, max-1)
}

func getNumber(min, max, skip int, guessed chan bool) (nchan chan int) {
	nchan = make(chan int, 0)
	go func() {
		i := rand.Intn(max-min) + min
		s := rand.Intn(3) - 1
		for {
			select {
			case nchan <- i:
			case <-guessed:
				close(nchan)
			}
			i += s * skip
			switch {
			case i >= max:
				i = min
			case i < min:
				i = max - 1
			}
		}
	}()
	return nchan
}

func play(name string, number int) (won bool) {
	fmt.Printf("Estas la vico de %s: kiu numero estas? ", name)
	var guess int
	fmt.Scan(&guess)
	switch {
	case number == guess:
		fmt.Println("Prave, la numero estas", number)
		return true
	case number > guess:
		fmt.Println("Ne, la numero estas pli alta ol", guess)
	case number < guess:
		fmt.Println("Ne, la numero estas pli malalta ol", guess)

	}
	return false
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	min, max, skip := 0, 100, 1
	players := 2
	klariguLudon(min, max, skip)
	guessed := make(chan bool)
	numbers := getNumber(min, max, skip, guessed)
	for number := range numbers {
		for i := 1; i <= players; i++ {
			name := fmt.Sprintf("Ludanto %d", i)
			if play(name, number) {
				fmt.Printf("Gratulon %s, vi gajnis!\n", name)
				return
			}
		}
	}
}
