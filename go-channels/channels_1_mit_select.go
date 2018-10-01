package main

import "fmt"

func main() {
	// erstellen der Channels
	gerade := make(chan int)
	ungerade := make(chan int)
	quit := make(chan int)

	// Funktion: senden
	go sende(gerade, ungerade, quit)

	// Funktion: empfangen
	empfange(gerade, ungerade, quit)
}
// Definition der Funktionen
func empfange(e, o, q <-chan int) {
	for {
		select {
		case v := <-e: // falls es aus dem "gerade-Channel" kommt
			fmt.Println("aus gerade-channel:", v)
		case v := <-o: // falls es aus dem "ungerade-Channel" kommt
			fmt.Println("aus ungerade-channel:", v)
		case _, ok := <-q: // Quit mit Status-Meldung
			fmt.Println("Quit:", ok)
			return
		}
	}
}
func sende(e, o, q chan<- int) {
	for i := 0; i < 13; i++ {
		if i%2 == 0 { // wenn durch 2 teilbar in den "gerade-Channel"...
			e <- i
		} else { // ansonsten...
			o <- i
		}
	}
	q <- 0   //wenn fertig: quit
	close(e) // schließen der Channel, da
	close(o) // Channel blocken

}
