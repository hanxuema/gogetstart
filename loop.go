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

}
