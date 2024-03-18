package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

type Row struct {
	Value_1  string `parquet:"name=value_1, type=UTF8"`
	Value_2  string `parquet:"name=value_2, type=UTF8"`
	Value_3  string `parquet:"name=value_3, type=UTF8"`
	Value_4  string `parquet:"name=value_4, type=UTF8"`
	Value_5  string `parquet:"name=value_5, type=UTF8"`
	Value_6  string `parquet:"name=value_6, type=UTF8"`
	Value_7  string `parquet:"name=value_7, type=UTF8"`
	Value_8  string `parquet:"name=value_8, type=UTF8"`
	Value_9  string `parquet:"name=value_9, type=UTF8"`
	Value_10 string `parquet:"name=value_10, type=UTF8"`
	Value_11 string `parquet:"name=value_11, type=UTF8"`
	Value_12 string `parquet:"name=value_12, type=UTF8"`
	Value_13 string `parquet:"name=value_13, type=UTF8"`
	Value_14 string `parquet:"name=value_14, type=UTF8"`
	Value_15 string `parquet:"name=value_15, type=UTF8"`
	Value_16 string `parquet:"name=value_16, type=UTF8"`
	Value_17 string `parquet:"name=value_17, type=UTF8"`
	Value_18 string `parquet:"name=value_18, type=UTF8"`
	Value_19 string `parquet:"name=value_19, type=UTF8"`
	Value_20 string `parquet:"name=value_20, type=UTF8"`
	Value_21 string `parquet:"name=value_21, type=UTF8"`
	Value_22 string `parquet:"name=value_22, type=UTF8"`
	Value_23 string `parquet:"name=value_23, type=UTF8"`
	Value_24 string `parquet:"name=value_24, type=UTF8"`
	Value_25 string `parquet:"name=value_25, type=UTF8"`
	Value_26 string `parquet:"name=value_26, type=UTF8"`
	Value_27 string `parquet:"name=value_27, type=UTF8"`
	Value_28 string `parquet:"name=value_28, type=UTF8"`
	Value_29 string `parquet:"name=value_29, type=UTF8"`
	Value_30 string `parquet:"name=value_30, type=UTF8"`
	Value_31 string `parquet:"name=value_31, type=UTF8"`
	Value_32 string `parquet:"name=value_32, type=UTF8"`
	Value_33 string `parquet:"name=value_33, type=UTF8"`
	Value_34 string `parquet:"name=value_34, type=UTF8"`
	Value_35 string `parquet:"name=value_35, type=UTF8"`
	Value_36 string `parquet:"name=value_36, type=UTF8"`
	Value_37 string `parquet:"name=value_37, type=UTF8"`
	Value_38 string `parquet:"name=value_38, type=UTF8"`
	Value_39 string `parquet:"name=value_39, type=UTF8"`
	Value_40 string `parquet:"name=value_40, type=UTF8"`
	Value_41 string `parquet:"name=value_41, type=UTF8"`
	Value_42 string `parquet:"name=value_42, type=UTF8"`
	Value_43 string `parquet:"name=value_43, type=UTF8"`
	Value_44 string `parquet:"name=value_44, type=UTF8"`
	Value_45 string `parquet:"name=value_45, type=UTF8"`
	Value_46 string `parquet:"name=value_46, type=UTF8"`
	Value_47 string `parquet:"name=value_47, type=UTF8"`
	Value_48 string `parquet:"name=value_48, type=UTF8"`
	Value_49 string `parquet:"name=value_49, type=UTF8"`
	Value_50 string `parquet:"name=value_50, type=UTF8"`
}

func main() {
	// var wg sync.WaitGroup
	start := time.Now()
	filePath := "./test_dataset10K.parquet"

	fr, err := local.NewLocalFileReader(filePath)
	if err != nil {
		log.Fatal("Can't open file", err)
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, nil, 4)
	if err != nil {
		log.Fatal("Can't create parquet reader", err)
	}
	defer pr.ReadStop()

	rows := make([]Row, 10000)

	if err := pr.Read(&rows); err != nil {
		log.Fatal("Error reading parquet file", err)
	}

	numWorkers := runtime.NumCPU()
	jobs := make(chan Row, 10000)
	results := make(chan int, 10000)

	for w := 0; w < numWorkers; w++ {
		go worker(jobs, results)
	}

	for _, row := range rows {
		jobs <- row
	}
	close(jobs)

	totalSum := 0
	for a := 0; a < len(rows); a++ {
		totalSum += <-results
	}

	fmt.Printf("Total sum: %d\n", totalSum)
	elapsed := time.Since(start)
	fmt.Printf("Function took %s\n", elapsed)

}

func worker(jobs <-chan Row, results chan<- int) {
	for row := range jobs {
		result := processRow(row)
		results <- result
	}
}

func processRow(row Row) int {
	v := reflect.ValueOf(row)
	for i := 0; i < v.NumField()-1; i++ {
		if v.Field(i).String() == "a" {
			if nextValue, err := strconv.Atoi(v.Field(i + 1).String()); err == nil {
				return nextValue
			}
		}
	}
	return 0
}
