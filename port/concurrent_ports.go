package port

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

// flar permite pasar parametros cuando corremos nuestro programa en Go

// el primer parametro es nombre del flag, el segundo un valor por defecto y el tecero un descrioncion

// es necesario poner en la funcion main del programa utilizar el metodo flag.Parse() para que lea los valores

// para correrlo simplmente hacemos go run main.go --site=scanme.webscantest.com

// func main() {
// 	flag.Parse()
// 	scanner.VerifyPortsCurrently()
// }

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func VerifyPortsCurrently() {
	flag.Parse()
	var wg sync.WaitGroup

	for scannedPort := 0; scannedPort < 65535; scannedPort++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(scannedPort)
	}

	wg.Wait()
}
