package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/aldinokemal/go-whatsapp-web-multidevice/cmd"
)

func reloadConnection() {
	go func() {
		for range time.Tick(5 * time.Minute) {
			resp, err := http.Get("https://whatsapp-vet.onrender.com")
			if err != nil {
				fmt.Println("Hubo un error al ejecutar la API:", err)
				continue
			}
			defer resp.Body.Close()
			fmt.Println("La API se ejecutó con éxito")
		}
	}()
}

func main() {
	go reloadConnection()
	cmd.Execute()
}

