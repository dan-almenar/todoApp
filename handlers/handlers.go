package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dan-almenar/todoapp/data"
)

type TaskLogger struct {
	l *log.Logger
}

func NewTaskLogger(l *log.Logger) *TaskLogger {
	return &TaskLogger{l}
}

func (t *TaskLogger) postNewTask(rw http.ResponseWriter, r *http.Request) error {
	task := &data.TodoTask{
		ID: data.GetNextID(),
	}
	err := task.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return err
	}
	t.l.Printf("Task: %#v", task)
	err = data.SaveTask(task)
	if err != nil {
		http.Error(rw, "Cannot save task to database", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (t *TaskLogger) updateTask(rw http.ResponseWriter, r *http.Request) error {
	taskList := data.GetTasks()
	updateTask := &data.TodoTask{}
	err := updateTask.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		return err
	}
	fmt.Printf("id: %d", updateTask.ID)

	for _, task := range taskList {
		if task.ID == updateTask.ID {
			*task = *updateTask
			break
		}
	}
	data.UpdateTasksList(taskList)
	rw.Write([]byte("Task succesfully updated\n"))
	return nil
}

func (t *TaskLogger) deleteTask(rw http.ResponseWriter, r *http.Request) error {
	taskList := data.GetTasks()
	var newTaskList = data.Tasks{}
	deleteTask := &data.TodoTask{}
	err := deleteTask.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusInternalServerError)
		return err
	}
	for _, task := range taskList {
		if task.ID != deleteTask.ID {
			newTaskList = append(newTaskList, task)
		}
	}
	data.UpdateTasksList(newTaskList)
	rw.Write([]byte("Task succesfully deleted\n"))
	return nil
}

func (t *TaskLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	t.l.Printf("Request received: %s\n", r.Method)
	if r.Method == "POST" {
		t.postNewTask(rw, r)
	} else if r.Method == "PUT" {
		t.updateTask(rw, r)
	} else if r.Method == "DELETE" {
		t.deleteTask(rw, r)
	} else {
		taskList := data.GetTasks()
		err := taskList.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		}
	}
}
