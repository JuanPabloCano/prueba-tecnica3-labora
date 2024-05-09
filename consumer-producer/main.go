package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Función para producir/fabricar numeros aleatorios y enviarlos en un canal

func productor(numeros chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		numero := rand.Intn(100)
		numeros <- numero
		fmt.Printf("Productor: Enviado número %d\n", numero)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) //Simulamos una produccion irregular
	}

	close(numeros)
}

//Funcion  consumidora que recibe numeros del canal y los procesa

func consumidor(numeros <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for numero := range numeros {

		fmt.Printf("Consumidor: Recibido número %d\n", numero)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//Creamos el canal de comunicacion entre productores y consumidores

	numeros := make(chan int)

	// Waitgropu para esperar la finalizacion de las goroutines
	var wg sync.WaitGroup

	// Iniciamos las goroutines de productor y consumidor

	wg.Add(2)
	go productor(numeros, &wg)
	go consumidor(numeros, &wg)

	//Esperamos a que terminen

	wg.Wait()
}
