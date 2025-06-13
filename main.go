package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("⚠️  Debes especificar al menos un puerto como argumento.")
		fmt.Println("👉  Ejemplo: go run main.go 8080 3000 9000")
		os.Exit(1)
	}

	wslIP := "localhost"

	fmt.Printf("✅ IP de WSL detectada: %s\n", wslIP)

	for _, port := range os.Args[1:] {
		go startProxy(port, wslIP, port)
	}

	fmt.Println("🚀 Proxy iniciado. Presiona Ctrl+C para detener.")
	select {}
}

func startProxy(localPort, wslIP, wslPort string) {
	listenAddr := "0.0.0.0:" + localPort
	targetAddr := wslIP + ":" + wslPort

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Printf("❌ Error al escuchar en el puerto %s: %v\n", localPort, err)
		return
	}
	defer listener.Close()

	fmt.Printf("🔗 Escuchando en %s -> Redirigiendo a %s\n", listenAddr, targetAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("❌ Error al aceptar conexión en el puerto %s: %v\n", localPort, err)
			continue
		}
		go handleConnection(conn, targetAddr)
	}
}

func handleConnection(clientConn net.Conn, targetAddr string) {
	defer clientConn.Close()

	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		fmt.Printf("❌ Error al conectar con el servicio de WSL en %s: %v\n", targetAddr, err)
		return
	}
	defer targetConn.Close()

	go io.Copy(targetConn, clientConn)
	io.Copy(clientConn, targetConn)
}
