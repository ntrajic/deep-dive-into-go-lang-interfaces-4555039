package tempus

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func ExampleMarshal() {
	d := Date{2024, 02, 17}

	fmt.Println("- JSON")
	data, err := json.Marshal(d)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))

	var d2 Date
	if err := json.Unmarshal(data, &d2); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(d2)

	fmt.Println("- XML")
	data, err = xml.Marshal(d)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))

	var d3 Date
	if err := xml.Unmarshal(data, &d3); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(d3)

	// Output:
	// - JSON
	// "2024-02-17"
	// {2024 2 17}
	// - XML
	// <Date>2024-02-17</Date>
	// {2024 2 17}
}
