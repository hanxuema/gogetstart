package main

func loops() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			println("break at ",i)
			continue
		}
		println("continuing.......")
	} 

	for i := 0; i < 5; i ++{
		println(i)
	}

	slice := []int{1,2,3}
	for i, v := range slice {
		println(i, v)
	}

	ports := map[string]int{"http": 80, "https":443}
	for k, v := range ports {
		println(k,v)
	}
}
