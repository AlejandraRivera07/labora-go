package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Type     string
	Content  string
	Complete bool
}

func main() {
	// Crea canales para cada tipo de mensaje
	emailCh := make(chan *Message)
	smsCh := make(chan *Message)
	pushCh := make(chan *Message)

	// Crea una espera para coordinar la finalización de las gorutinas
	var wg sync.WaitGroup

	// Inicia las gorutinas para procesar cada tipo de mensaje
	wg.Add(1)
	go processMessages("email", emailCh, &wg)
	wg.Add(1)
	go processMessages("SMS", smsCh, &wg)
	wg.Add(1)
	go processMessages("push", pushCh, &wg)

	// Agrega mensajes a la cola
	for i := 0; i < 10; i++ {
		msg := &Message{
			Type:    "email",
			Content: fmt.Sprintf("Mensaje de correo electrónico #%d", i),
		}
		emailCh <- msg
	}
	for i := 0; i < 10; i++ {
		msg := &Message{
			Type:    "SMS",
			Content: fmt.Sprintf("Mensaje de SMS #%d", i),
		}
		smsCh <- msg
	}
	for i := 0; i < 10; i++ {
		msg := &Message{
			Type:    "push",
			Content: fmt.Sprintf("Mensaje de notificación push #%d", i),
		}
		pushCh <- msg
	}

	// Espera a que todas las gorutinas finalicen
	wg.Wait()
}

// processMessages procesa los mensajes en orden de llegada
func processMessages(typ string, ch chan *Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		fmt.Printf("Procesando mensaje %q de tipo %q...\n", msg.Content, msg.Type)

		// Simula un tiempo de procesamiento variable
		time.Sleep(time.Duration(len(msg.Content)) * 10 * time.Millisecond)

		// Marca el mensaje como completado
		msg.Complete = true

		fmt.Printf("Mensaje %q de tipo %q completado\n", msg.Content, msg.Type)
	}
	fmt.Printf("Gorutina de procesamiento de mensajes %q finalizada\n", typ)
}
