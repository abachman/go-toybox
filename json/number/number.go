package main

import (
	"encoding/json"
	"fmt"
)

type DataPoint struct {
	Value json.Number `json:"value,omitempty"`
}

func main() {
	var list []DataPoint

	list = append(list, DataPoint{Value: json.Number("1")})
	list = append(list, DataPoint{Value: json.Number("1.1")})

	for _, v := range list {
		// print as various types
		i, ierr := v.Value.Int64()
		if ierr != nil {
			fmt.Println("error converting", v.Value, "to int64")
		} else {
			fmt.Println("int64  ", i)
		}

		f, ferr := v.Value.Float64()
		if ferr != nil {
			fmt.Println("error converting", v.Value, "to float64")
		} else {
			fmt.Println("float64", f)
		}

		fmt.Println("string ", v.Value.String())

		// print JSON output
		js, _ := json.Marshal(v)
		fmt.Println(string(js))

		fmt.Println()
	}
}
