package main

import "fmt"



func main()  {
	x := make([]string, 27, 27)
	fmt.Println("First Time")
	fmt.Println(len(x), cap(x))
	x = append(x, "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal", "Espirito Santo", "Goiás", "Maranhão", "Mato Grosso", "MAto Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins")
	
	fmt.Println("Second Time")
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	y := make([]string, 27, 27)
	fmt.Println("Third Time with Y")
	fmt.Println(len(y), cap(y))

	estados := []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal", "Espirito Santo", "Goiás", "Maranhão", "Mato Grosso", "MAto Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	for i, v := range estados {
		y[i] = v
	}

	fmt.Println("Fourth Time with Y")
	fmt.Println(y)
	fmt.Println(len(y), cap(y))

	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}
}