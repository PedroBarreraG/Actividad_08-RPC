package main

import (
	"fmt"
	"net/rpc"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Agregar la calificación de un alumno")
		fmt.Println("2) Obtener el promedio del alumno")
		fmt.Println("3) Obtener el promedio de todos los alumnos")
		fmt.Println("4) Obtener el promedio por materia")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var nombre string
			var materia string
			var calificaion string
			fmt.Print("Nombre del Alumno: ")
			fmt.Scanln(&nombre)
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			fmt.Print("Calificaión: ")
			fmt.Scanln(&calificaion)
			var datos string
			datos = nombre + "," + materia + "," + calificaion
			var result string
			err = c.Call("Server.AgregarCalificacion", datos, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 2:
			var nombre string
			fmt.Print("Nombre del Alumno: ")
			fmt.Scanln(&nombre)

			var result float64
			err = c.Call("Server.PromedioAlumno", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio es: ", result)
			}
		case 3:
			var algo string
			algo = "1"
			var result float64
			err = c.Call("Server.PromedioGeneral", algo, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio general es: ", result)
			}
		case 4:
			var nombre string
			fmt.Print("Materia: ")
			fmt.Scanln(&nombre)

			var result float64
			err = c.Call("Server.PromedioMateria", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio es: ", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
