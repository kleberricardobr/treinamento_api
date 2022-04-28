package main

import (
	_ "github.com/lib/pq"
)

/*
func main() {
	log.Println("Conectando no banco de dados...")
	database.ConnectDataBase()
	routes.GetRoutes()

}
*/

/*
func main() {
	//exemplo de slices
	var vSlice []int
	vSlice = append(vSlice, 1)
	vSlice = append(vSlice, 2)
	vSlice = append(vSlice, 3)

	//varrendo o slice
	for _, it := range vSlice {
		fmt.Println(it)
	}

	vSlice2 := []int{1, 2, 3, 4}
	for _, it := range vSlice2 {
		fmt.Println(it)
	}

	vSlice3 := [5]string{"Item1", "Item2", "Item3", "Item4", "Item5"}
	vSliceDoSlice3 := vSlice3[1:4]

	fmt.Printf("O tamanho do vSlice3 %d \n", len(vSlice3))
	fmt.Printf("O tamanho do vSliceDoSlice3 %d \n", len(vSliceDoSlice3))
	fmt.Printf("%v \n", vSliceDoSlice3)

	//capacidade = 10, tamanho = 2
	vSlice4 := make([]string, 2, 10)
	fmt.Println("Tamanho do vSlice4 " + strconv.Itoa(len(vSlice4)))
	fmt.Println("Capacidade do vSlice4 " + strconv.Itoa(cap(vSlice4)))

	vSlice4[0] = "Item1"
	vSlice4[1] = "Item2"

	fmt.Printf("%v \n", vSlice4)

	//ponteiros
	var vPoint *int
	vValor := 3

	vPoint = &vValor

	fmt.Printf("Antes de alterar pelo ponteiro: %d \n", vValor)
	*vPoint = 4
	fmt.Printf("Após alterar pelo ponteiro: %d \n", vValor)

}

*/
