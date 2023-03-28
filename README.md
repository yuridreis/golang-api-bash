# Golang API that executes bash commands.

This application expects a POST request with either a query param or a json and executes the command sent.

The application expects to be run in a linux environment. (or in a linux running bash)

## Information about the route:

1- The endpoint of the application is "/api".

2- The port of the application is "4444"

http://localhost:4444/api

## Example of a request with json:

The body of the json expects a key value pair with the key beeing "data" and the value the command.

```sh
{
    "data" : "mkdir new_folder"
}
```

## Example of a request with query param:

POST: http://localhost:4444/api?data=mkdir%20new_folder

## Running locally

Cloning the repository and compile/running the program.

```sh
git clone git@github.com:yuridreis/golang-bash-api.git
go run main.go
```


