package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type suit int

const (
	suitHeart suit = iota
	suitClub
	suitDiamond
	suitSpade
)

type card struct {
	suit   suit
	number int
}

func main() {
	marks := map[suit]string{
		suitHeart:   "♥",
		suitClub:    "♣",
		suitDiamond: "◆",
		suitSpade:   "♠",
	}

	//make all the cards

	all := make([]*card, 0, 13*4)
	for s := suitHeart; s <= suitSpade; s++ {
		for n := 2; n <= 14; n++ {
			all = append(all, &card{
				suit:   s,
				number: n,
			})
		}
	}

	//make random numbers
	t := time.Now().UnixNano()
	rand.Seed(t)

	//shuffle all the cards
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	//Raw capital for gambling
	coins := 100
	var useCoins int

	//start the game loop
	for coins > 0 && len(all) >= 5 {
		//Place your bets

		for {

			fmt.Printf("Please give your bet, up to %d coins", coins)
			fmt.Printf("> ")
			fmt.Scanln(&useCoins)
			if useCoins > 0 && useCoins <= coins {
				break
			}
			fmt.Printf("Please enter the correct number of bet.")
		}
		//separate hand cards and the remaining cards
		handCards := all[:5]
		all = all[5:]

		//sort the hand cards by number
		sort.Slice(handCards, func(i, j int) bool {
			return handCards[i].number < handCards[j].number
		})

		//display the handcards
		fmt.Println("Following are your hand cards:")
		for _, c := range handCards {
			fmt.Print(marks[c.suit] + " ")
			switch c.number {
			case 11:
				fmt.Println("J")
			case 12:
				fmt.Println("Q")
			case 13:
				fmt.Println("K")
			case 1:
				fmt.Println("A")
			default:
				fmt.Println(c.number)
			}
		}

		//replace some of the hand cards with remaining cards
		var replace int
		for {
			fmt.Println("Rest cards: ", len(all))
			fmt.Println("How many hand cards would you like to replace with the new cards? (UP TO 5)")
			fmt.Print("> ")
			fmt.Scanln(&replace)
			if replace >= 0 && replace <= 5 {
				break
			}
			println("Please enter the correct number")
		}

		handCards = append(handCards[:5-replace], all[:replace]...)
		all = all[replace:]

		//resort the handcards
		sort.Slice(handCards, func(i, j int) bool {
			return handCards[i].number < handCards[j].number
		})

		//reprint the handcards
		fmt.Println("Following are your hand cards:")
		for _, c := range handCards {
			fmt.Print(marks[c.suit] + " ")
			switch c.number {
			case 11:
				fmt.Println("J")
			case 12:
				fmt.Println("Q")
			case 13:
				fmt.Println("K")
			case 1:
				fmt.Println("A")
			default:
				fmt.Println(c.number)
			}
		}

		//count the points of handcards
		numCount := make(map[int]int)

		var maxSame int
		isStraight := true
		isFlash := true
		for i := 0; i < len(handCards); i++ {
			numCount[handCards[i].number]++
			if maxSame < numCount[handCards[i].number] {
				maxSame = numCount[handCards[i].number]
			}
			if i > 0 {
				isStraight = isStraight && handCards[i].number-handCards[i-1].number == 1
				isFlash = isFlash && handCards[i].suit == handCards[i-1].suit
			}
		}

		//calculate the ratio!
		var ratio int
		switch {
		case isStraight && isFlash && handCards[0].number == 10:
			fmt.Println("ROYAL FLASH!!!")
			ratio = 100
		case isStraight && isFlash:
			fmt.Println("Straight Flash! WOW")
			ratio = 50

		case maxSame == 4:
			fmt.Println("You got 4 cards the same!")
			ratio = 20
		case len(numCount) == 2:
			fmt.Println("Half half!")
			ratio = 7
		case isFlash:
			fmt.Println("You got a FLASH!")
			ratio = 5
		case isStraight:
			fmt.Println("You got a STRAIGHT!")
			ratio = 5
		case maxSame == 3:
			fmt.Println("You got Triplets!")
			ratio = 3
		case len(numCount) == 3:
			fmt.Println("You got two pair")
			ratio = 2
		case len(numCount) == 4:
			fmt.Println("You got one pair")
			ratio = 1
		default:
			fmt.Println("役無し")
		}

		increase := useCoins * ratio
		newCoins := coins - useCoins + increase
		fmt.Printf("%d * %d = %d\n", useCoins, ratio, increase)
		fmt.Printf("Your fortune: %d >>> %d\n", coins, newCoins)
		coins = newCoins

	}
	fmt.Println("Sorry, your fortune has gone out =v=")
	fmt.Println("Your final fortune:", coins)

}
