package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

var materias = make(map[string]map[string]float64)
var alumnos = make(map[string]map[string]float64)

var alum []string
var mat []string

type Server struct{}

func (this *Server) AgregarCalificacion(datos string, reply *string) error {
	dat := strings.Split(datos, ",")
	nombre := dat[0]
	mate := dat[1]
	calificacionString := dat[2]
	calificacion, _ := strconv.ParseFloat(calificacionString, 64)

	if alumnos[nombre][mate] == 0 {
		bandA := 0
		bandM := 0
		for _, name := range alum {
			if name == nombre {
				bandA = 1
			}
		}
		for _, name := range mat {
			if name == mate {
				bandM = 1
			}
		}

		if bandM == 0 {
			mat = append(mat, mate)
			var alumno = make(map[string]float64)
			alumno[nombre] = calificacion
			materias[mate] = alumno
		} else {
			materias[mate][nombre] = calificacion
		}

		if bandA == 0 {
			alum = append(alum, nombre)
			var materia = make(map[string]float64)
			materia[mate] = calificacion
			alumnos[nombre] = materia
		} else {
			alumnos[nombre][mate] = calificacion
		}
		*reply = "La calificación del alumno fue registrada con exito"
	} else {
		*reply = "El alumno ya tiene calificación"
	}
	return nil
}

func (this *Server) PromedioAlumno(nombre string, reply *float64) error {
	var promedio float64
	var numMaterias float64
	promedio = 0
	numMaterias = 0
	for materia, calificacion := range alumnos[nombre] {
		fmt.Println(materia+":", calificacion)
		promedio = promedio + calificacion
		numMaterias = numMaterias + 1
	}
	fmt.Println("-------------------------")
	promedio = promedio / numMaterias
	*reply = promedio
	return nil
}

func (this *Server) PromedioGeneral(algo string, reply *float64) error {
	var promedio float64
	var numMaterias float64

	var numAlumnos float64
	numAlumnos = 0
	var promediogeneral float64
	promediogeneral = 0

	for _, name := range alum {
		promedio = 0
		numMaterias = 0
		fmt.Println("Alumno: ", name)
		for materia, calificacion := range alumnos[name] {
			fmt.Println(materia+":", calificacion)
			promedio = promedio + calificacion
			numMaterias = numMaterias + 1
		}
		promedio = promedio / numMaterias
		fmt.Println("Promedio: ", promedio)
		promediogeneral = promediogeneral + promedio
		numAlumnos = numAlumnos + 1
	}
	fmt.Println("-------------------------")
	promediogeneral = promediogeneral / numAlumnos
	*reply = promediogeneral
	return nil
}

func (this *Server) PromedioMateria(mate string, reply *float64) error {
	var promedio float64
	var numAlumnos float64
	promedio = 0
	numAlumnos = 0
	for alumno, calificacion := range materias[mate] {
		fmt.Println(alumno+":", calificacion)
		promedio = promedio + calificacion
		numAlumnos = numAlumnos + 1
	}
	fmt.Println("-------------------------")
	promedio = promedio / numAlumnos
	*reply = promedio
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
