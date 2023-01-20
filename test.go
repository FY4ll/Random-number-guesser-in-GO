package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomnum := rand.Intn(11-1) + 1
	var UserInput int
	var Restart string
	try := 3
	for {
		fmt.Printf("Entrez un numéro entre 1 et 10: ")
		fmt.Scanln(&UserInput)
		if UserInput == 0 {
			fmt.Println("Nombre invalide")
			break
		} else if UserInput > 10 {
			fmt.Println("Nombre trop Grand")
			break
		}
		if UserInput == randomnum {
			fmt.Println("Bravo!!")
			fmt.Println("Vous avez gagnez avec", try, "tentative(s) restantes")
			fmt.Printf("Voulez vous recommencer? [Y/N]")
			fmt.Scanln(&Restart)
			if Restart == "Y" || Restart == "y" {
				main()
			} else {
				fmt.Println("GoodBye")
				break
			}
		} else if UserInput < randomnum {
			try = try - 1
			fmt.Println("Votre chiffre est trop petit")
			fmt.Println("Vous avez encore", try, "tentative(s)")
		} else {
			try = try - 1
			fmt.Println("Votre chiffre est trop grand")
			fmt.Println("Vous avez encore", try, "tentative(s)")
		}
		if try == 0 {
			fmt.Println("Vous avez perdu")
			fmt.Println("le nombre mystère était: ", randomnum)
			fmt.Printf("Voulez vous recommencer? [Y/N]")
			fmt.Scanln(&Restart)
			if Restart == "Y" || Restart == "y" {
				main()
			} else {
				fmt.Println("GoodBye")
				break
			}
		}
	}
}
