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
			resp, err := http.Get("https://whatsappvet3-tmohe8jg.b4a.run")
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

