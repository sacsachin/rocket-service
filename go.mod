module example.com/microservice/rocket-service

go 1.16

require (
	example.com/microservice/rocket-proto v0.0.0
	github.com/go-kit/kit v0.11.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/mock v1.6.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.2
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210828152312-66f60bf46e71 // indirect
	google.golang.org/grpc v1.40.0
)

replace example.com/microservice/rocket-proto => ../rocket-proto
