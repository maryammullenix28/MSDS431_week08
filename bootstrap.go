package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/gonum/stat"
)

func main() {
	// Read the CSV file
	file, err := os.Open("Celtics_Heat_Game_6_Actual.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip the header row
	if _, err := reader.Read(); err != nil {
		log.Fatal(err)
	}

	// Parse the rest of the data
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the data
	var XDrafted []float64
	var FPTS []float64

	for _, record := range records {

		// Skip the first two columns
		record = record[2:]

		// Parse XDrafted
		xDrafted, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		XDrafted = append(XDrafted, xDrafted)

		// Parse FPTS
		fpts, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		FPTS = append(FPTS, fpts)
	}

	startTime := time.Now()

	R := 10000
	bootstrapCorrelation := bootstrapCorrelation(XDrafted, FPTS, R)

	endTime := time.Now()

	// Print results
	fmt.Println("Time taken:", endTime.Sub(startTime))
	fmt.Println("Bootstrap Correlation:")
	fmt.Println("Mean:", stat.Mean(bootstrapCorrelation, nil))
	fmt.Println("Standard Deviation:", stat.StdDev(bootstrapCorrelation, nil))
	min, max := getRange(bootstrapCorrelation)
	fmt.Println("Range:", min, max)
}

func bootstrapCorrelation(XDrafted, FPTS []float64, R int) []float64 {
	result := make([]float64, R)

	for i := 0; i < R; i++ {
		sample := sampleWithReplacement(XDrafted, FPTS)
		bootstrapCorrelation := stat.Correlation(sample.XDrafted, sample.FPTS, nil)
		result[i] = bootstrapCorrelation
	}

	return result
}

func sampleWithReplacement(XDrafted, FPTS []float64) struct {
	XDrafted []float64
	FPTS     []float64
} {
	n := len(XDrafted)
	result := struct {
		XDrafted []float64
		FPTS     []float64
	}{
		XDrafted: make([]float64, n),
		FPTS:     make([]float64, n),
	}

	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		result.XDrafted[i] = XDrafted[j]
		result.FPTS[i] = FPTS[j]
	}

	return result
}

func getRange(bootstrapCorrelation []float64) (float64, float64) {
	min, max := bootstrapCorrelation[0], bootstrapCorrelation[0]
	for _, v := range bootstrapCorrelation {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
