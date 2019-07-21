# goa with out grpc 

go build ./cmd/calc && go build ./cmd/calc-cli

# Run the server

./calc
[calcapi] 21:35:36 HTTP "Add" mounted on GET /add/{a}/{b}
[calcapi] 21:35:36 HTTP "./gen/http/openapi.json" mounted on GET /openapi.json
[calcapi] 21:35:36 serving gRPC method calc.Calc/Add
[calcapi] 21:35:36 HTTP server listening on "localhost:8000"
[calcapi] 21:35:36 gRPC server listening on "localhost:8080"

# Run the client

# Contact HTTP server
$ ./calc-cli --url="http://localhost:8000" calc add --a 1 --b 2
3

# Contact gRPC server
$ ./calc-cli --url="grpc://localhost:8080" calc add --message '{"a": 1, "b": 2}'
3
