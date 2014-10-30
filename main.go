package main

import (
	"fmt"
	"math/rand"
	"os"
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

func askPlayers() (n int) {
	fmt.Print("Kiom da ludantoj estas? [2] ")
	fmt.Scan(&n)
	if n < 1 {
		return 2
	}
	return n
}

func askNames(n int) (names []string) {
	for i := 1; i <= n; i++ {
		name := ""
		fmt.Printf("Ludanto %d, kiel mi nomu vin? [Ludanto %d] ", n, n)
		fmt.Scan(&name)
		if name == "" {
			name = fmt.Sprintf("Ludanto %d", n)
		}
		names = append(names, name)
	}
	return names
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
	switch len(os.Args) {
	case 0:
		panic("This shouldn't happen: os.Args is an empty slice.")
	case 1:
	case 2:
		fmt.Sscan(os.Args[1], &max)
	case 3:
		fmt.Sscan(os.Args[1]+" "+os.Args[2], &min, &max)
	default:
		fmt.Sscan(os.Args[1]+" "+os.Args[2]+" "+os.Args[3], &min, &max, &skip)
	}
	klariguLudon(min, max, skip)
	players := askPlayers()
	names := askNames(players)
	guessed := make(chan bool)
	numbers := getNumber(min, max, skip, guessed)
	for {
		for _, name := range names {
			if play(name, <-numbers) {
				fmt.Printf("Gratulon %s, vi gajnis!\n", name)
				return
			}
		}
	}
}
