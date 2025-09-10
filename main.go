package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Print("\033[2J\033[H")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
		groese := 16
	for {

		eingabe := 0
		eingabe2 := 0
		eingabe3 := 0
		gb, kb, sz, z := true, true, true, true
		clearScreen()
		fmt.Printf(`
[1] Generieren
[2] Einstellungen
`)
		fmt.Print("\n-> ")
		fmt.Scan(&eingabe)

		buchstabenG := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		buchstabenK := "abcdefghijklmnopqrstuvwxyz"
		zahlen := "0123456789"
		zeichen := "!@#$%^&*()_+-=[]{}|;:,.<>?"

		switch eingabe {
		case 1:
			clearScreen()
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

			fmt.Printf("\nDein Password: %s\n", string(password))
			fmt.Scan(&eingabe3)
		case 2:
			for eingabe2 != 9 {
				clearScreen()
				fmt.Printf(`
[1] Länge ändern 	   [%v]
[2] Benutze Großbuchstaben [%v]
[3] Benutze Kleinbuchstben [%v]
[4] Benutze Zahlen	   [%v]
[5] Benutze Sonderzeichen  [%v]

[9] Back
`, groese, gb, kb, z, sz)
				fmt.Print("\n-> ")
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
			}

		default:
			fmt.Printf("Ungültige eingabe")
		}

	}

}

