# FIFA World Cup Winners

This project exposes a Web API for accessing historic data from
the FIFA World Cup championship.

## Running tests

A proper Go environment is required in order to run this project.
Once setup, tests can be run with the following command:

`go test -v ./handlers/`

## Running the server

Once all tests are passing, the server can be started with
the `go run server.go` command.

## Testing the API manually

Start the server with `go run server.go` and then
use the example commands printed to the console to
test the program.
like so
```
go run server.go


        GETting:

        curl -i http://localhost:8000/
        curl -i http://localhost:8000/winners
        curl -i http://localhost:8000/winners?year=1970
        curl -i http://localhost:8000/winners?year=banana

        POSTing with NO access token:

        curl -i -X POST -d "{\"country\":\"Croatia\", \"year\": 2030}" http://localhost:8000/winners

        POSTing with valid access token:

        curl -i -X POST -H "X-ACCESS-TOKEN: 5577006791947779410" -d "{\"country\":\"Croatia\", \"year\": 2030}" http://localhost:8000/winners

        Then check for the newly added winner

        curl -i http://localhost:8000/winners

        POSTing with invalid data:

        curl -i -X POST -H "X-ACCESS-TOKEN: 5577006791947779410" -d "{\"country\":\"Russia\", \"year\": 1984}" http://localhost:8000/winners

        POSTing with invalid method:

        curl -i -X PUT -d "{\"country\":\"Russia\", \"year\": 2030}" http://localhost:8000/winners
```

### Running with Docker

To build the image from the Dockerfile, run:

`docker build -t project-fifa-world-cup .`

To start an interactive shell, run:

`docker run -it --rm --name run-fifa project-fifa-world-cup`

From inside the shell, run the tests with:

`go test handlers/*`
