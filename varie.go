package main

import (
	"bufio"
	"fmt"
	"os"
)

type operazioni interface {
	salvadati(name string, cogn string, age int, address string)
	recuperadati() utente
	recuperanome() string
}
type utente struct {
	nome      string
	cognome   string
	eta       int
	indirizzo string
}

func (str utente) salvadati(name string, cogn string, age int, address string) {
	str.nome = name
	str.cognome = cogn
	str.eta = age
	str.indirizzo = address
}
func (str utente) recuperadati() utente {
	return utente{str.nome, str.cognome, str.eta, str.indirizzo}
}
func (str utente) recuperanome() string {
	return str.nome
}

func main() {
	var us utente
	fmt.Println("-Inserimento Dati-")
	fmt.Println("<Nuovo utente>")
	fmt.Printf("Inserisci Nome : ")
	fmt.Scan(&us.nome)
	fmt.Printf("Inserisci Cognome : ")
	fmt.Scan(&us.cognome)
	fmt.Printf("Inserisci età : ")
	fmt.Scan(&us.eta)
	fmt.Printf("Inserisci Indirizzo : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	us.indirizzo = scanner.Text()
	var user operazioni = utente{us.nome, us.cognome, us.eta, us.indirizzo}
	user.salvadati(us.nome, us.cognome, us.eta, us.indirizzo)
	prova := user.recuperadati()
	fmt.Println("Utente : ", prova)
	fmt.Printf("Utente - Nome : %s\n", user.recuperanome())
	fmt.Println("Esempio di Interfaccia !!!")

}
