package main

import "fmt"

//nil as default value on interface, function, map, slice, pointer, channel
func main() {
	data := newMap("Iqbal Arrafi")
	if data == nil {
		fmt.Println("Data kosong")
	}
	fmt.Println(data)
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
