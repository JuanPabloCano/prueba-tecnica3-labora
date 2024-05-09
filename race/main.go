package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//creamos la estructura del objeto competidor

type Competidor struct {
	nombre    string
	velocidad int
}

// Vamos a crear la funcion de simulacion de progreso de los competidores en la carrera

func correrCompetidor(competidor Competidor, pista chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for distancia := 0; distancia < 100; {
		avance := rand.Intn(competidor.velocidad)
		pista <- avance
		distancia += avance
		time.Sleep(100 * time.Millisecond) //Aqui vamos a simular el tiempo de la carrera

	}

	fmt.Printf("%s ha terminado la carrera! \n", competidor.nombre)
}

//El time sleep simula la duracion de la carrera

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//es simplemente un timer enlazado a las señales de reloj del procesador

	competidores := []Competidor{
		{"Competidor 1", rand.Intn(10) + 1},
		{"Competidor 2", rand.Intn(10) + 1}, // Velocidades entre el 1 y 10
		{"Competidor 3", rand.Intn(10) + 1},
		{"Competidor 4", rand.Intn(10) + 1},
	}

	// Creamos un canal para representar la pista
	pista := make(chan int)

	//Nos aseguramos de cerrar el canal cuando se termina la carrera
	defer close(pista)

	//Waitgroup para esperar la finalizacion de todas las go routines
	var wg sync.WaitGroup
	wg.Add(len(competidores))

	//Iniciamos una goroutine para cada competidor

	for _, competidor := range competidores {
		go correrCompetidor(competidor, pista, &wg)
	}

	/*Monitoreamos el progreso de los competidores y definimos
	el ganador*/

	for {
		select {
		case avance := <-pista:
			fmt.Printf("Avance: %d\n", avance)
		case <-time.After(1 * time.Second):
			fmt.Println("No hay avance en 1 segundo...")

		}

		//verificamos si alguien gana la carrera

		select {
		case <-pista:
			fmt.Println("¡Tenemos un ganador!")
			wg.Wait()
			return
		default:
			//No hay ganador y se continua la carrera
		}
	}
}
