package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "6fad5235c96f6fc4443e48dd6a3a3c40"

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s", city) // Se escribe el mensaje en el canal
	return data
}

func main() {
	startNow := time.Now()
	cities := []string{"Toronto", "London", "Paris", "Tokyo"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1) // Se agrega un grupo de acuerdo a la cantidad de nodos requeridos
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait() // Se espera a que se finalice la goroutine, de lo contrario el sistema continuaria el flujo. Este es un proceso bloqueante
		close(ch) // Se cierra el canal para no tener fugas
	}()

	for result := range ch {
		fmt.Println(result)
	}
	// De estas dos formas se puede leer o suscribirse al canal para obtener los datos escritos en el
	result := <-ch
	fmt.Println("Result: ", result)

	fmt.Println("This operation took:", time.Since(startNow))
}
