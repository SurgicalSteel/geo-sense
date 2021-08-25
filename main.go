package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

func main() {
	const (
		fileName    string = "dataset/worldcities_small.csv"
		workerCount int    = 25
	)

	var (
		orderLocation = Coordinate{
			Latitude:  -7.7315203,
			Longitude: 110.4141519,
		}
	)

	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer logFile.Close()

	multi := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multi)

	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("ERROR Opening dataset file %s. Detail :%s\n", fileName, err.Error())
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatalf("ERROR Reading dataset file. Detail :%s\n", err.Error())
	}

	haversineCalculationData := parseCSV(csvLines, orderLocation)
	resultData := make([]HaversineCalculationData, len(haversineCalculationData))

	//if there's no calculation job, then return
	if len(haversineCalculationData) == 0 {
		return
	}

	calculationStartTime := time.Now()
	log.Printf("[GEO-SENSE] Start Working on Distance Calculation on %s.\n", calculationStartTime.Format("2006-01-02T15:04:05.999999-07:00"))

	jobChan := make(chan HaversineCalculationData, len(haversineCalculationData))
	resultChan := make(chan HaversineCalculationData, len(haversineCalculationData))

	for i := 1; i <= workerCount; i++ {
		go worker(i, jobChan, resultChan)

	}

	for i := 0; i < len(haversineCalculationData); i++ {
		jobChan <- haversineCalculationData[i]
	}

	close(jobChan)

	var wg sync.WaitGroup

	for i := 0; i < len(haversineCalculationData); i++ {
		wg.Add(1)
		tempResult := <-resultChan
		resultData[i] = tempResult
		wg.Done()
	}

	wg.Wait()

	sort.Sort(HaversineCalculationDataOrdering(resultData))

	resultWarehouse := resultData[0]
	log.Println("====================================================================================")
	log.Println(resultWarehouse)
	log.Println("====================================================================================")
	calculationEndTime := time.Now()
	calculationDuration := time.Since(calculationStartTime)
	log.Printf("[GEO-SENSE] Result Closest Warehouse Name %s and distance %.5f Kilometers\n", resultWarehouse.WarehouseName, resultWarehouse.Distance)
	log.Printf("[GEO-SENSE] Finished Working on Distance Calculation on  %s. Duration : %.5f seconds", calculationEndTime.Format("2006-01-02T15:04:05.999999-07:00"), calculationDuration.Seconds())

}

func worker(id int, jobChan <-chan HaversineCalculationData, resultChan chan<- HaversineCalculationData) {
	for job := range jobChan {
		startTime := time.Now()
		log.Printf(
			"[%s] Worker with ID:%d is starting to execute job %+v.\n",
			startTime.Format("2006-01-02T15:04:05.999999-07:00"),
			id,
			job,
		)

		distance := CalculateDistance(job.WarehouseLocation, job.OrderLocation)
		job.Distance = distance
		resultChan <- job

		endTime := time.Now()
		log.Printf(
			"[%s] Worker with ID:%d has finished executing job %+v, result=%.5f\n",
			endTime.Format("2006-01-02T15:04:05.999999-07:00"),
			id,
			job,
			distance,
		)

	}

}

// The willow it weeps today
// A breeze from the distance is calling your name
// Unfurl your black wings and wait
// Across the horizons coming to sweep you away
// It's coming to sweep you away

// Let the wind carry you home
// Blackbird fly away
// May you never be broken again

// The fragile can not endure
// The wrecked and jaded
// A place so impure
// The static of this cruel world
// 'Cause some birds to fly long before they've seen their day
// Long before they've seen their day

// Let the wind carry you home
// Blackbird fly away
// May you never be broken again

// Beyond the suffering you've known
// I hope you find your way
// May you never be broken again

// Ascend, may you find no resistance
// Know that you've made such a difference
// And all you leave behind
// Will live till the end

// The cycle of suffering goes on
// But the memories of you stay strong
// Someday I too will fly and find you again

// Let the wind carry you home
// Blackbird fly away
// May you never be broken again

// Beyond the suffering you've known
// I hope you find your way
// May you never be broken again
// May you never be broken again
