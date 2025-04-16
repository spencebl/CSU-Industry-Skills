package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Label struct {
	Job string  `json:"job title"`
	GT  float64 `json:"gt label"`
}

type Job struct {
	JobTitle           string   `json:"job title"`
	CoveredSkills      []string `json:"covered skills"`
	MissingSkills      []string `json:"missing skills"`
	IntersectingSkills []string `json:"intersecting skills"`
	GtLabel            float64  `json:"gt label"`
}

func main() {
	data_encoder := []Label{}
	file_writer, err := os.Create("qualitative_analysis_comparison.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file_writer.Close()

	fileData, err := os.ReadFile("qualitative_analysis.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var data map[string]interface{}
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	for roleName, roleData := range data {
		fmt.Println("Role:", roleName)

		roleMap, ok := roleData.(map[string]interface{})
		if !ok {
			fmt.Println("  Unexpected format for role data.")
			continue
		}

		skillsRaw, found := roleMap["required_skills"]
		if !found {
			fmt.Println("  No required_skills field found.")
			continue
		}

		skillsMap, ok := skillsRaw.(map[string]interface{})
		if !ok {
			fmt.Println("  required_skills is not a map.")
			continue
		}

		total := len(skillsMap)
		countNotNone := 0

		for _, value := range skillsMap {
			strValue, isString := value.(string)
			if isString && strValue != "None" {
				countNotNone++
			}
		}

		percentage := float64(countNotNone) / float64(total)
		fmt.Printf("  -> %d/%d relevant skills (%.10f)\n\n", countNotNone, total, percentage)
		data_encoder = append(data_encoder, Label{
			Job: roleName,
			GT:  percentage,
		})

	}
	encoder := json.NewEncoder(file_writer)
	encoder.SetIndent("", "  ") // Pretty-print
	if err := encoder.Encode(data_encoder); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}

	job_file := "../Skill_Comparison/skill_comparison.json"
	data_file, err := os.ReadFile(job_file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var jobs []Job
	err = json.Unmarshal(data_file, &jobs)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	filename := "compare.csv"

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"job title", "gt", "qa", "percent change"})
	if err != nil {
		panic(err)
	}
	for _, job := range jobs {
		for _, calc := range data_encoder {
			if calc.Job == job.JobTitle {
				fmt.Printf("Job Title: %s\n\tGT Label: %.10f\n\tCaldulated: %.10f\n\tPercent Change: %.3f%%\n", job.JobTitle, job.GtLabel, calc.GT, (100 * (calc.GT - job.GtLabel) / job.GtLabel))
				datacsv := []string{
					job.JobTitle,
					strconv.FormatFloat(job.GtLabel, 'f', -1, 64),
					strconv.FormatFloat(calc.GT, 'f', -1, 64),
					strconv.FormatFloat((100 * (calc.GT - job.GtLabel) / job.GtLabel), 'f', -1, 64),
				}
				err := writer.Write(datacsv)
				if err != nil {
					panic(err)
				}
			}
		}
	}

}
