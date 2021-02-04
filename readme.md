# Go REST API for Beer Ingredients
Add ingredients to create your favourite beer! The rest-API supports GET, POST, and DELETE requests.

## Run

First, install Mux:
```go
go get -u github.com/gorilla/mux
```

Then, run the API:
```go
go run main.go
```

## Testing

Use an API testing tool like Postman or https://reqbin.com/ and send a POST request like adding instrument:

```json
{
	"ID" : "0",
	"Type" : "Cryptocurrency",
	"Name" : "BTC",
	"Price" : 27303.96,
	"Quantity" : 1,
}
```


##  Input:

```go
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
```