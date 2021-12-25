protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./greet/greetpb/greet.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./calculator/calculatorpb/calculator.proto 

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative blog/blogpb/blog.proto 

/Users/ysimeonov/mongodb-macos-x86_64-5.0.5/bin/mongod --dbpath data/db/

