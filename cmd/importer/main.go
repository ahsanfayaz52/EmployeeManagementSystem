package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"

	runtime "github.com/ahsanfayaz52/EmployeeManagementSystem"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/config"
	"github.com/ahsanfayaz52/EmployeeManagementSystem/models"
)

func main() {
	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	var rows, _ = readCsv(viper.GetString(config.Filepath))
	importData(rows, rt)
}

func importData(rows [][]string, rt *runtime.Runtime) {
	for _, row := range rows {
		_, err := rt.Service().AddEmployee(&models.Employee{
			Name:    row[0],
			Address: row[1],
			Age:     atoi(row[2]),
			Salary:  atoi(row[3]),
			Phone:   atoi(row[4]),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readCsv(name string) ([][]string, error) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	firstRow, err := bufio.NewReader(f).ReadSlice('\n')
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Seek(int64(len(firstRow)), io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	return csv.NewReader(f).ReadAll()
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
