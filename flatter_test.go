package flatter

import (
	"fmt"
	"testing"
)

var jsonStr = `{
   	"cart_id":27325477,
	"name": "Gopher",
	"total_amount": 520.79,
	"promo": true,
   	"delivery_address":{
      "contact_number": "08123456789",
	  "city": "Jakarta",
	  "zip_code": "11440"
   	},
	"books": [
		{"title": "AAA"},
		{"title": "BBB"},
		{"title": "CCC"}
	],
	"opts": [ "Monday", "Tuesday", "Wednesday" ]
}`

func TestFlatter(t *testing.T) {
	result, err := Flatter(jsonStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v", result)
}

func BenchmarkFlatter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Flatter(jsonStr)
	}
}