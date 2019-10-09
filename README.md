# flatter
Transform json string to golang map in flat way

This is simple repo works just like https://github.com/fatih/structs, which transforming to golang map,
unless the source is directly from json string. This one also supports for nested struct and array.
- Nested Example
```go
jsonStr := {
  "delivery_address": {
    "contact_number": "08123456789",  // Output: map["delivery_address.contact_number"]
    "city": "Jakarta",                // Output: map["delivery_address.city"]
    "zip_code": "11440"               // Output: map["delivery_address.zip_code"]
  }
}
```
- Array Example
```go
jsonStr := {
  "opts": [ "Monday", "Tuesday", "Wednesday" ] // Output: map["opts[0]"] = Monday opts[1]:Tuesday opts[2]:Wednesday
}
```


### Use Case
Sometimes we need data structure to be able automatically manipulate things from dynamic input (such as from template).

### Example
You have json to be transformed to map:
```json
{
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
}
```
Call the function:
```go
    result, err := Flatter(jsonStr)
    if err != nil {
    	fmt.Println(err.Error())
    	return
    }
    
    fmt.Printf("%+v", result)
    // Output: map[books[0].title:AAA books[1].title:BBB books[2].title:CCC cart_id:27325477 delivery_address.city:Jakarta delivery_address.contact_number:08123456789 delivery_address.zip_code:11440 name:Gopher opts[0]:Monday opts[1]:Tuesday opts[2]:Wednesday promo:true total_amount:520.79]
```

### Benchmark
```text
goos: darwin
goarch: amd64
pkg: github.com/napazca/flatter
BenchmarkFlatter-8        200000             10986 ns/op
PASS
ok      github.com/napazca/flatter      2.332s
```