package main

func main() {
	s := []string{"foo", "bar"}

	for idx, v := range s {
		println(idx)
		println(v)
	}

	myMap := make(map[string]string)
	myMap["first"] = "xyz"
	myMap["second"] = "abc"

	for k, v := range myMap {
		println(k, v)
	}
}
