package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListarTareas(tasks []Task) {

	fmt.Printf(`
LISTA DE TAREAS ACTUALES
───────────────────────────────────
ESTADO | ID | TAREA
───────────────────────────────────`)

	if len(tasks) == 0 {
		fmt.Println("\nNo hay tareas")
		return
	}

	for _, task := range tasks {

		status := " "
		if task.Complete {
			status = "🟢"
		} else {
			status = "⚪"
		}

		fmt.Printf("\n[ %s ]   %-4d %-30s\n", status, task.ID, task.Name)
	}
}

func AgregarTareas(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       ObtenerId(tasks),
		Name:     name,
		Complete: false,
	}

	return append(tasks, newTask)
}

func EliminarTarea(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func CompletarTarea(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}

	return tasks
}

func DesmarcarTarea(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = false
			break
		}
	}
	return tasks
}

func Guardar(file *os.File, tasks []Task) {
	byte, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(byte)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}

func ObtenerId(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
