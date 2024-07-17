package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

type AnimalList struct {
	Animals []string `json:"animals"`
}

func loadAnimalList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var animalList AnimalList
	err = json.NewDecoder(file).Decode(&animalList)
	if err != nil {
		return nil, err
	}

	return animalList.Animals, nil
}

func getRandomAnimal(animals []string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return animals[rand.Intn(len(animals))]
}

func sendPostRequest(endpoint, animal string) (string, error) {
	data := url.Values{}
	data.Set("message", animal)
	encodedData := data.Encode()

	log.Printf("Sending request: %s", encodedData)

	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", bytes.NewBufferString(encodedData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Printf("Received response: %s", string(body))

	return string(body), nil
}

func GetEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func main() {
	animals, err := loadAnimalList("animals.json")
	if err != nil {
		log.Fatalf("Error loading animal list: %v", err)
	}

	webHost := GetEnv("WEB_HOST", "localhost")
	webPort := GetEnv("WEB_PORT", "8081")
	url := fmt.Sprintf("http://%v:%v", webHost, webPort)

	for {
		animal := getRandomAnimal(animals)
		_, err := sendPostRequest(url, animal)
		if err != nil {
			log.Printf("Error sending POST request: %v", err)
		} else {
			fmt.Printf("Sent: %s\n", animal)
		}

		sleepInterval := 10 + rand.Intn(11)
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
