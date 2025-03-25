package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	codes := make(map[string][]string)
	columnIndex := 1

	f, err := excelize.OpenFile("../industry-codes/all-industry-codes.xlsx")
	// f, err := excelize.OpenFile("../csu-codes/all-csu-codes.xlsx")

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Printf("Error reading rows in sheet %s: %v", sheet, err)
			continue
		}
		for i := 1; i < len(rows); i++ {
			row := rows[i]
			if len(row) > columnIndex {
				codes[row[columnIndex]] = append(codes[row[columnIndex]], sheet)

			}
		}
	}
	// delete(codes, "Spencer Baloga Loufek")
	// delete(codes, "03-24-25")
	// delete(codes, "Masters Project")
	// delete(codes, "CS152 - Modules.")

	// for key, value := range codes {
	// 	fmt.Printf("%s: %d\n", key, len(value))
	// }
	file, err := os.Create("../industry-codes/all-industry-codes.csv")
	// file, err := os.Create("../csu-codes/all-csu-codes.csv")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '|'
	defer writer.Flush()
	for key, values := range codes {
		line := fmt.Sprintf("%s||[%s]", key, joinArray(values))
		if _, err := file.WriteString(line + "\n"); err != nil {
			log.Fatalf("Error writing row to CSV: %v", err)
		}
	}

	fmt.Println("CSV file 'output.csv' has been created successfully!")
}
func joinArray(arr []string) string {
	return fmt.Sprintf("%s", formatArray(arr))
}

// Format array values into a comma-separated string
func formatArray(arr []string) string {
	result := ""
	for i, v := range arr {
		if i > 0 {
			result += ","
		}
		result += v
	}
	return result
}
