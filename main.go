package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	task "github.com/DiegoMaes17/CLI-CRUD/tasks"
)

// Colores
const (
	Rojo           = "\033[31m"
	Verde          = "\033[32m"
	Azul           = "\033[34m"
	ReiniciarColor = "\033[0m"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)

		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}

	}

	for {
		if len(os.Args) < 2 {
			opcion := MenuOpciones()
			if opcion == 7 {
				fmt.Println("Gracias por usar este programa 🌐")
				break
			}
			tasks = Menu(opcion, tasks, file)
		}
	}
}

func Menu(opcion int, tasks []task.Task, file *os.File) []task.Task {
	LimpiarPantalla()
	switch opcion {

	case 1:
		LimpiarPantalla()
		task.ListarTareas(tasks)
		return tasks

	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nNombre de tarea a añadir")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AgregarTareas(tasks, name)
		task.Guardar(file, tasks)
		fmt.Print("Tarea añadida con exito ✅")
		return tasks

	case 3:
		var id int
		LimpiarPantalla()
		task.ListarTareas(tasks)
		fmt.Println("\nIntroduce el id de la tarea que deseas completar:")

		_, err := fmt.Scanln(&id)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Debes proporcionar un ID para completar")
			} else {
				fmt.Println("El id debe ser un numero valido")
			}
			return tasks
		}

		tasks = task.CompletarTarea(tasks, id)
		task.Guardar(file, tasks)
		fmt.Println("Tarea completada con exito ✅")
		return tasks

	case 4:
		var id int
		LimpiarPantalla()
		task.ListarTareas(tasks)
		fmt.Println("\nIntroduce el id de la tarea que deseas desmarcar:")
		_, err := fmt.Scanln(&id)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Debes proporcionar un ID para desmarcar")
			} else {
				fmt.Println("El id debe ser un numero valido")
			}
			return tasks
		}

		tasks = task.DesmarcarTarea(tasks, id)
		task.Guardar(file, tasks)
		fmt.Println("Tarea desmarcada con exito ✅")
		return tasks

	case 5:
		var id int
		LimpiarPantalla()
		task.ListarTareas(tasks)
		fmt.Println("\nIntroduce el id de la tarea que deseas borrar:")
		_, err := fmt.Scanln(&id)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Debes proporcionar un ID para modificar")
			} else {
				fmt.Println("El id debe ser un numero valido")
			}
			return tasks
		}

		tasks = task.EliminarTarea(tasks, id)
		task.Guardar(file, tasks)
		fmt.Println("Tarea borrada con exito")
		return tasks

	case 6:
		LimpiarPantalla()
		return tasks
	default:
		return tasks

	}

}

func MenuOpciones() int {
	fmt.Printf(`%s
	CLI-CRUD Menu:
	 1- Listar
	 2- Añadir
	 3- Completar
	 4- Desmarcar
	 5- Borrar
	 6- Limpiar pantalla
	 7- Salir
	 %s`, Azul, ReiniciarColor)

	fmt.Printf(`%sIngrese algun numero (1-7): %s`, Verde, ReiniciarColor)

	var opcion int

	_, err := fmt.Scanln(&opcion)
	if err != nil || opcion < 1 || opcion > 7 {
		fmt.Printf("%s\n Opcion invalida%s\n", Rojo, ReiniciarColor)
	}

	return opcion
}

func LimpiarPantalla() {
	fmt.Print("\033[H\033[2J")
}
