package main

func main() {
	var str string
	var data []byte
	// START OMIT
	str = "..." + str
	data = append([]byte{'.', '.', '.'}, data...)
	// END OMIT
	_ = str
	_ = data
}
