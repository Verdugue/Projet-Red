package main

import (
	"fmt"

	gamePlay "PROJET-RED/src"
)

func main() {
	// Créer un personnage personnalisé avec la fonction CharCreation
	fmt.Println("BIENVENUE DANS  LE JEUX D'AVENTURE")
	gamePlay.Tuto()
	gamePlay.ClearConsole()
	p := gamePlay.CharCreation()

	// Créer une variable pour suivre si la potion gratuite est disponible
	gamePlay.ClearConsole()
	gamePlay.InfoDebut(p)
	for {
		// Afficher le menu principal
		fmt.Println("\nMenu :")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Equipement")
		fmt.Println("6. Combat")
		fmt.Println("7. Qui-sont'ils ?")
		fmt.Println("0. Quitter")
		/*
		   2. Les equipements ne s'affiche pas et n'octroie pas hp max en plus.  1
		   3. Les objets acheté s'affiche à l'index 1. ==
		   4. Ca n'affiche pas que je suis deja full hp et consomme les potions.
		   5. Pouvoir désequipper des objets dans l'inventaire.4
		   6. Nb Tours bugger apres l'inventaire. 2
		   10. Deux 0 dans l'inventaire.
		   11. Ajouter du graphisme.
		   13. supprimer les lots et ajouter une fonction pour demander combien d'item je veux. 3
		*/

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// Afficher les informations du personnage
			gamePlay.ClearConsole()
			gamePlay.DisplayInfo(p)
		case 2:
			// Afficher le contenu de l'inventaire
			gamePlay.ClearConsole()
			gamePlay.AccessInventory(p)
		case 3:
			// Afficher le menu du marchand
			gamePlay.ClearConsole()
			gamePlay.MarchandMenu(p)
		case 4:
			// Afficher le menu du forgeron
			gamePlay.ClearConsole()
			gamePlay.Forgeron(p)
		case 5:
			gamePlay.ClearConsole()
			gamePlay.Equipe(p)
		case 6:
			gamePlay.ClearConsole()
			gamePlay.ChoisirAdversaire(p)
		case 7:
			gamePlay.Qui()
		case 0:
			// Quitter
			fmt.Println("Au revoir !")
			return
		default:
			// Choix invalide, veuillez réessayer.
			fmt.Println("Choix invalide, veuillez réessayer.")
		}

	}
}
