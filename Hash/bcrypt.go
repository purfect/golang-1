package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "passwort123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil { // Error-Handling
		fmt.Println(err)
	}
	fmt.Printf("Klartext:\t%s\n", s)     // gibt noch einmal den zu hash-enden String aus
	fmt.Printf("Bcrypt-Hash:\t%s\n", bs) // gibt den String in ge-hash-ter Form aus

	// Vergleich zwischen ge-hash-ten und unge-hash-ten Strings:

	// Cost-Parameter ermitteln
	/* Anmerkung:
    *   Neuere Passwort-Hashing-Funktionen haben einen Cost-Paramter.append
    *   Dieser gibt an, wieviel "Arbeit" (z.B. CPU-Zeit) aufgewendet wurde um diesen Hash zu erstellen
    *   Definition:
    *   MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
    *   MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
    *   DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
	*/
	c, err := bcrypt.Cost(bs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("Cost:\t\t%d\n", c)
	fmt.Println("[ Check  ]")
	fmt.Printf("Test-Hash:\t%s\n", bs)
	fmt.Printf("Plaintext:\t%s\n", s)
	err = bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil { // Wenn es nicht das richtige Passwort war
		fmt.Printf("Return:\t\tFalse\n")
		return
	}
	fmt.Printf("Return:\t\tTrue\n")

}
