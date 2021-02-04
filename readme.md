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

Use an API testing tool like Postman and send a ```POST``` request and add an instrument like:

```json
{
	"ID" : "3",
	"Type" : "Cryptocurrency",
	"Name" : "BTC",
	"Price" : 27303.96,
	"Quantity" : 1,
}
```