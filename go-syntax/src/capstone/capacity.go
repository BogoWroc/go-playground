package capstone

import "fmt"

func AppendString(value string, count int) {
	data := make([]string, 0, 1)
	for i := 0; i < count; i++ {
		data = append(data, value)
		fmt.Printf("%v has %d lenght and %d capacity\n", data, len(data), cap(data))
	}
}
