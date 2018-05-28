package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"github.com/Jeffail/gabs"
	"time"
)

var LastRestarted time.Time

func restartAllWorkers() {
	log.Println("Restarting all workers")

	result, err := http.Get("http://localhost:21000/service/inspection/workers")
	if err != nil {
		log.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(result.Body)
	json, _ := gabs.ParseJSON(body)
	workers, _ := json.Path("worker_summaries").Children()

	for _, child := range workers {
		worker_id := child.Path("worker_id").Data().(string)
		req, err := http.NewRequest("DELETE", "http://localhost:21000/service/inspection/workers/" + worker_id, nil)
		http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
	}

	LastRestarted = time.Now()
}

func scanForAssemblyChanges() {
	files, err := ioutil.ReadDir("build/assembly/worker")
	if err != nil {
		log.Fatal(err)
	}

	shouldRestart := false

	for _, file := range files {
		if LastRestarted.Before(file.ModTime()) {
			shouldRestart = true
		}
	}

	if shouldRestart {
		restartAllWorkers()
	}
}

func main() {
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			scanForAssemblyChanges()
		}
	}()
	time.Sleep(240 * time.Hour)
}


