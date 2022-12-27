package port

import (
	"fmt"
	"net"
)

func VerifyPorts() {
	for port := 0; port < 65535; port++ {
		// 1, 2, ..4, 99
		// sitio:1, sitio:2, sitio:99, ...,
		// 1 -> Open,  2 -> Closed, ...

		// el metodo Dial hace  (o intenta hacer una conexion) a un puerto de un dominio
		// parametors tipo de conexion + dominio
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))
		// si no hay error es qeu logramos conectar
		if err != nil {
			continue
		}
		// cierra la conexion, solo imprimimos el puerto descubuierto
		conn.Close()
		fmt.Printf("Port %d is open\n", port)
	}
}
