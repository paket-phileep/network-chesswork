package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fsnotify/fsnotify"
)

type ScanResult struct {
	Hostname string `json:"hostname"`
	Latency  string `json:"latency"`
	Vendor   string `json:"vendor"`
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	//STUB - information about the watcher
	fmt.Println("Successfully began watching results.txt for file changes")

	defer watcher.Close()

	err = watcher.Add("results.txt")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Modified file:", event.Name)
					processFile("results.txt", "result.json")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-done
}

func processFile(inputFile, outputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	results := parseResults(string(data))

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully updated", outputFile)
}

func parseResults(data string) map[string]ScanResult {
	results := make(map[string]ScanResult)
	re := regexp.MustCompile(`Nmap scan report for (.+?)\nHost is up \((.+?)s latency\).\nMAC Address: (.+?) \((.+?)\)`)

	matches := re.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		macAddress := match[3]
		results[macAddress] = ScanResult{
			Hostname: match[1],
			Latency:  match[2],
			Vendor:   match[4],
		}
	}
	return results
}
