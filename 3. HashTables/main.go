package main

import "fmt"

func main() {
	salarios := map[string]int{"Kaue": 2000, "Nathalia": 0, "Gabriel": 1700}
	fmt.Printf("Pessoa: %s, salário: %d\n", "kaue", salarios["Kaue"]) // é case sensitive!

	delete(salarios, "Gabriel")

	salarios["Gabriel"] = 8700

	salarios_vazios := make(map[string]int) // função make, usada para inicializar map, slice ou chan

	salarios_vazios["Kaue"] = 1888
	println(salarios_vazios["Kaue"])
}
