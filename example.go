package main

import (
	"fmt"
	"github.com/loudbund/go-convey/convey_v1"
)

func main() {
	d := convey_v1.NewConvey()

	// ShouldEqual
	if err := d.ShouldEqual(123, 321); err != nil {
		fmt.Println(err)
	}
	// ShouldEqual
	if err := d.ShouldEqual(123, 123); err != nil {
		fmt.Println(err)
	}
	// JShouldEqual
	if _, err := d.JShouldEqual(map[string]interface{}{
		"abc": []interface{}{"a", 3},
	}, []interface{}{"abc", 1}, 3); err != nil {
		fmt.Println(err)
	}
	// JShouldEqual
	if _, err := d.JShouldEqual(map[string]interface{}{
		"abc": []interface{}{"a", 3},
	}, []interface{}{"abc", 0}, 3); err != nil {
		fmt.Println(err)
	}

	fmt.Println(d.GetNum())
}
