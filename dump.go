package _dump

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/dump"
)

func Print(a interface{}) {
	dump.P(a)
}

func Json(a interface{}) {
	d, _ := json.MarshalIndent(a, "", "\t")
	fmt.Println(string(d))
}
