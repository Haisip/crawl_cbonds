package main

import (
	"github.com/daominah/crawler_xsmb/pkg/driver/source_cbondshnxvn"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	err := godotenv.Load(`.env`)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	refreshDataIntervalStr := os.Getenv("REFRESH_DATA_INTERVAL_SECONDS")
	refreshDataIntervalInt, err := strconv.Atoi(refreshDataIntervalStr)
	if err != nil {
		log.Fatalf("bad env REFRESH_DATA_INTERVAL_SECONDS, strconv.Atoi error: %v", err)
	}
	refreshDataInterval := time.Duration(refreshDataIntervalInt) * time.Second
	log.Printf("recognized env REFRESH_DATA_INTERVAL_SECONDS: %v", refreshDataInterval)

	outputPath := os.Getenv("OUTPUT_FILE")
	log.Printf("outputPath: %v", outputPath)

	var dataSource source_cbondshnxvn.DataSourceCbonds
	job := func() {
		freshData, err := dataSource.DownloadHTML()
		if err != nil {
			log.Printf("error sourceXsktcomvn.DownloadHTML: %v", err)
			return
		}
		rows, err := dataSource.ParseToResultStockNews(string(freshData))
		if err != nil {
			log.Printf("error ParseXsktcomvn: %v", err)
			return
		}
		log.Printf("rows: %+v", rows)

		err = source_cbondshnxvn.WriteToCsv(rows, outputPath)
		if err != nil {
			log.Printf("error WriteToCsv: %v", err)
		} else {
			log.Printf("succesfully wrote to %v", outputPath)
		}
	}

	for {
		log.Printf("begin job")
		job()
		log.Printf("end job")
		time.Sleep(refreshDataInterval)
	}
}
