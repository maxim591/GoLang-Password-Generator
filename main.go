package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	groese := 16
	eingabe := 0
	eingabe2 := 0
	gb, kb, sz, z := false, false, false, false
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Printf(`
	[1] Generieren
	[2] Einstellungen
	`)
		fmt.Scan(&eingabe)

		buchstabenG := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		buchstabenK := "abcdefghijklmnopqrstuvwxyz"
		zahlen := "0123456789"
		zeichen := "!@#$%^&*()_+-=[]{}|;:,.<>?"

		switch eingabe {
		case 1:
			// Generieren
			charset := ""
			if gb {
				charset += buchstabenG
			}
			if kb {
				charset += buchstabenK
			}
			if z {
				charset += zahlen
			}
			if sz {
				charset += zeichen
			}

			if charset == "" {
				fmt.Println("Warnung: Keine Zeichentypen ausgewählt!")
				continue
			}
			password := make([]byte, groese)
			for i := 0; i < groese; i++ {
				randomIndex := rand.Intn(len(charset))
				password[i] = charset[randomIndex]
			}
			fmt.Printf("Dein Password: %s\n", string(password))
		case 2:
			fmt.Printf(`
		[1] Länge ändern 	   [%v]
		[2] Benutze Großbuchstaben [%v]
		[3] Benutze Kleinbuchstben [%v]
		[4] Benutze Zahlen	   [%v]
		[5] Benutze Sonderzeichen  [%v]
		`, groese, gb, kb, z, sz)
			fmt.Scan(&eingabe2)
			if eingabe2 == 1 {
				fmt.Printf("Wie Lange soll es sein: ")
				fmt.Scan(&groese)
			}
			if eingabe2 == 2 {
				gb = !gb
			}
			if eingabe2 == 3 {
				kb = !kb
			}
			if eingabe2 == 4 {
				z = !z
			}
			if eingabe2 == 5 {
				sz = !sz
			}
		default:
			fmt.Printf("Ungültige eingabe")
		}

	}

}
