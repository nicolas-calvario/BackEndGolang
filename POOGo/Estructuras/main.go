package main

func main() {
	Go := Course{
		Name:    "Curso desde cero",
		Price:   2323.32,
		IsFree:  false,
		UserIds: []uint{1, 121, 2, 12},
		Classes: map[uint]string{
			1: "Introducion",
			2: "Estructuras",
			3: "Maps",
		},
	}
	// Go2 := Course{
	// 	"Curso desde cero",
	// 	2323.32,
	// 	false,
	// 	[]uint{1, 121, 2, 12},
	// 	map[uint]string{
	// 		1: "Introducion",
	// 		2: "Estructuras",
	// 		3: "Maps",
	// 	},
	// }
	// Css := Course{
	// 	Name:   "Css desde cero",
	// 	IsFree: true,
	// }
	// js := Course{}
	// js.Name = "Javascript"
	// js.UserIds = []uint{232, 23123, 123, 23}
	// fmt.Println(Go.Name)
	// fmt.Println(Go2.Name)
	// fmt.Printf("%+v", Css)
	// fmt.Printf("%+v", js)

	PrintClasses(Go)
	Course.PrintClasses2(Go)

}
