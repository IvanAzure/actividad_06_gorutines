package main

import (
	"fmt"
	"time"
)
var killed_process_id int = -1
var current_process_id int = 0
var flag bool = false

func menu() string {
	var opt string
	fmt.Println("[Administrador de procesos]\nIngresa una opcion\n1 - Agregar proceso\n2 - Mostrar procesos\n3 - Terminar proceso\n4 - Salir")
	fmt.Scan(&opt)
	return opt
}

func Proceso(id int) {
	i := uint64(0)
	for {
		if(flag){
			fmt.Println(id,":",i);
		}
		i = i + 1
		time.Sleep(time.Millisecond * 500)

		if killed_process_id == id {
			return
		}
	}
}

func kill_process() {
	var id int
	fmt.Print("Ingrese el ID del proceso:")
	fmt.Scan(&id)
	killed_process_id = id
	fmt.Print("Proceso eliminado")
}

func main() {
	var opt string
	for {
		opt = menu()
		switch opt {
		case "1":
			current_process_id++
			go Proceso(current_process_id)
		case "2":
			if flag == false{
			flag = true
			}else{
			flag = false
			}
		case "3":
			kill_process()
		case "4":
			return
		default:
			fmt.Println("")
		}
	}
}