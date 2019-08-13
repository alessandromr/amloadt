package ualyzer

import (
	"bufio"
	"log"
	"os"
	"testing"
)

/*
	Test
*/
func TestSuccessURL(t *testing.T) {
	list := getListFromFile("./test/clean_list.txt")
	for _, value := range list {
		_, err := Parse(value)
		if err != nil {
			log.Println("Error on: " + value)
			log.Fatalln(err)
		}
	}
}
func TestFailURL(t *testing.T) {
	list := getListFromFile("./test/fail_list.txt")
	for _, value := range list {
		_, err := Parse(value)
		if err == nil {
			log.Fatalln("Error: URL must fail")
		}
		if err.Error() != "Error: URL not valid" {
			log.Println("Error on: " + value)
			log.Fatalln(err)
		}
	}
}

/*
	Benchmark
*/

/*
	Commons
*/

func getListFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
	return list
}
