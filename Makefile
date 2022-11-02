# generate protos auth
gnpa:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protos/auth.proto

# generate protos user
gnpu:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protos/user.proto

# generate protos product
gnpp:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protos/product.proto

# generate protos image
gnpi:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative src/protos/images.proto

# run go
r:
	go run main.go

# run migration
migration:
	go run database/migration/migration.go

#docker
#build image
bimage:
	docker build -t server-alpha .

# build container
bcon:
	docker container create --name server-alpha --net grpc server-alpha

#run
rdoc:
	docker start server-alpha
