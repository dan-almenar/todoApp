package data

import (
	//"time"
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"log"
	"io"
)

type TodoTask struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Urgent bool `json:"urgent"`
}

type Tasks []*TodoTask

func (t *Tasks) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *TodoTask) FromJSON(r io.Reader) error{
	d := json.NewDecoder(r)
	return d.Decode(t)
}

func GetNextID() int{
	taskList := GetTasks()
	if len(taskList) == 0{
		return 1
	}
	lastTask := taskList[len(taskList) -1]
	return lastTask.ID +1
}

func getDataFile() *csv.Writer{
	csvFile, err := os.OpenFile("data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil{
		log.Printf("ERROR: %s", err)
	}
	return csv.NewWriter(csvFile)
}

func createCSVRow(t *TodoTask)([]string, error){
	var row = []string{
		strconv.Itoa(t.ID),
		t.Name,
		strconv.FormatBool(t.Urgent),
	}
	return row, nil
}

func SaveTask(t *TodoTask) error{
	writer := getDataFile()
	row, err := createCSVRow(t)
	if err != nil{
		log.Printf("ERROR: %s", err)
	}
	writer.Write(row)
	writer.Flush()
	return nil
}

func UpdateTasksList(t Tasks) error{
	csvFile, err := os.Create("data.csv")
	if err != nil{
		log.Print("ERROR: %s", err)
		return err
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, task := range t{
		row, err := createCSVRow(task)
		if err != nil{
			log.Printf("ERROR: %s", err)
			return err
		}
		writer.Write(row)
	}
	writer.Flush()
	return nil
}

func GetTasks() Tasks{
	taskList := Tasks{}
	csvFile, err := os.OpenFile("data.csv", os.O_CREATE|os.O_RDONLY, 0644)
	defer csvFile.Close()
	if err != nil{
		log.Printf("ERROR: %s", err)
	}
	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil{
		log.Printf("ERROR reading file: %s", err)
	}
	for _, line := range lines{
		id, _ := strconv.Atoi(line[0])
		urgent, _ := strconv.ParseBool(line[2])
		task := TodoTask{
			ID: id,
			Name: line[1],
			Urgent: urgent,
		}
		taskList = append(taskList, &task)
	}
	return taskList
}
