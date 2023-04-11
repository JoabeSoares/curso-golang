package main

import "fmt"

type vehicle struct{
	doors int
	colour string
}
type truck struct{
	vehicle
	fourWheel bool
}
type sedan struct{
	vehicle
	luxury bool
}

func main()  {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			colour: "Verde",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			colour: "Vermelho",
		},
		luxury: true,
	}
	

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
	fmt.Println(t.colour)
	fmt.Println(s.colour)
}