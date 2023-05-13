package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// saldo inicial
var balance int = 1000
var wg sync.WaitGroup // WaitGroup para sincronizar las goroutines

func main() {
	// imprimo saldo inicial
	fmt.Printf("El saldo inicial de la cuenta es :%d\n", balance)
	// Se establece una semilla para generar números aleatorios
	rand.Seed(time.Now().UnixNano())

	var w sync.WaitGroup
	var m sync.Mutex

	// Se crean varias goroutines que realizan transacciones aleatorias
	for i := 0; i < 10; i++ {
		w.Add(1)
		go transaction(&w, &m)
	}

	w.Wait() // Se espera a que todas las goroutines terminen
	fmt.Printf("El nuevo saldo de la cuenta es: %d\n", balance)
}

func transaction(wg *sync.WaitGroup, m *sync.Mutex) {

	m.Lock()
	// Se genera una transacción aleatoria
	amount := rand.Intn(501) // Genera un número aleatorio entre 1 y 500
	// variable booleana de 0 o 1 para sortear si es un retiro o un depósito
	isWithdrawal := rand.Intn(2) == 0

	// Se realiza la transacción correspondiente
	if isWithdrawal {
		fmt.Printf("Retiro de %d\n", amount)
		if balance < amount {
			fmt.Printf("La cuenta no tiene suficientes fondos para que retire %d\n", amount)
		} else {
			balance -= amount
		}

	} else {
		fmt.Printf("Depósito de %d\n", amount)
		balance += amount
	}
	m.Unlock()
	wg.Done()
}
