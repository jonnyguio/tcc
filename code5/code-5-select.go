package main

func readTemperature(temperatureCh chan float32) {
	// ...
	temperatureCh <- getLastTemperatureValue() // 
}

func readMoisture(moistureCh chan float32) {
	// ...
	moistureCh  <- getLastTemperatureValue() // 
}

func main() {
	tempChannel := make(chan float32)
	moistureChannel := make(chan float32)
	for {
		go readMoisture(moistureChannel)
		go readTemperature(tempChannel)
		select {
		case temp := <-tempChannel:
			processTemp(temp) //
		case moisture := <-moistureChannel:
			processMoisture(moisture) // 
		}
	}
}
