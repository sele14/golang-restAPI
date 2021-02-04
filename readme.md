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

Use an API testing tool like Postman or https://reqbin.com/ and send a POST request like adding another beer ingredient:

``
{
  	"UID" : "2",
	"Name" : "Barley",
	"Desc" : "Many brewers see barley as beer's soul. Barley has distinctive characteristics that make it a favored ingredient over other grains.",
	"Quantity" : 20
}
``