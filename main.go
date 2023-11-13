package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fabiokaelin/terminalimage"
)

type (
	country struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
)

func main() {
	fmt.Println("Hello, World!")
	content := readFile()
	contentMap := make(map[string]string)

	err := json.Unmarshal([]byte(content), &contentMap)
	if err != nil {
		log.Fatal(err)
	}

	countires := []country{}

	for k, v := range contentMap {
		countires = append(countires, country{
			Name: v,
			Code: k,
		})
	}

	// fmt.Println(countires)
	// fmt.Println(len(countires))

	correctAnswers := 0
	wrongAnswers := 0

	for _, c := range countires {

		imgString, err := printFlag(c.Code)

		if err != nil {
			continue
		}

		fmt.Println(imgString)
		// generate random number between 0 and len(countires) -1

		rand.Intn(len(countires) - 1)
		country1 := countires[rand.Intn(len(countires)-1)]
		country2 := countires[rand.Intn(len(countires)-1)]
		country3 := countires[rand.Intn(len(countires)-1)]
		country4 := c

		// shuffle countries
		countriesOptions := []country{country1, country2, country3, country4}
		rand.Shuffle(len(countriesOptions), func(i, j int) {
			countriesOptions[i], countriesOptions[j] = countriesOptions[j], countriesOptions[i]
		})

		fmt.Println("1. " + countriesOptions[0].Name)
		fmt.Println("2. " + countriesOptions[1].Name)
		fmt.Println("3. " + countriesOptions[2].Name)
		fmt.Println("4. " + countriesOptions[3].Name)
		// fmt.Println("---", c.Name, "---")
		fmt.Print("Your Answer: ")

		var input string
		fmt.Scanln(&input)

		if input == "q" {
			fmt.Println("")
			fmt.Println("Quit")
			fmt.Println("Correct answers: ", correctAnswers)
			fmt.Println("Wrong answers: ", wrongAnswers)
			os.Exit(0)
		} else {
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				// print in yellow
				fmt.Println("\033[33mWrong input\033[0m")
				continue
			}
			if countriesOptions[inputInt-1].Name == c.Name {
				// print in green
				fmt.Println("\033[32mCorrect!\033[0m")
				correctAnswers++
			} else {
				// print in red
				fmt.Println("\033[31mWrong!\033[0m")
				fmt.Println("\033[31mCorrect answer is: " + c.Name + "\033[0m")
				wrongAnswers++
			}
		}
		fmt.Println("")
		time.Sleep(300 * time.Millisecond)

	}
}

func printFlag(code string) (string, error) {

	imageString, err := terminalimage.ImageToString("flags/"+strings.ToLower(code)+".png", 20, true)
	if err != nil {
		return "", err
	}

	return imageString, nil
}

func readFile() string {
	// file, err := os.Open("countries.json")
	file, err := os.Open("de.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// read file

	content := make([]byte, 10240)
	count, err := file.Read(content)
	if err != nil {
		log.Fatal(err)
	}
	return string(content[:count])
}
