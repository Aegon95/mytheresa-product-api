# Mytheresa Product API



## Libraries used

Name  | Description
------------- | -------------
Chi Router  | For routing and middleware
Zap  | For logging
sqlx  | For ORM
viper  | For config management
migrate  | For Running DB scripts automatically
testify, httpexpect, sqlmock | For Unit tests


## How to run Test cases 

`go test ./...`

## How to run Application

### With Docker

1. `docker-compose up -d`


### without Docker

1. download golang 1.17
2. checkout github.com/Aegon95/mytheresa-product-api
3. run `go mod download`
4. install latest postgres and create new database called fashionstore
5. update postgres details in config/config.yml
6. go to cmd/api folder
7. run `go run .`
8. it will start the server at port 3000
9. Browse the api with the query params by going to http://localhost:3000/api/v1/products?category=boots&priceLessThan=90000

Decisions Taken

1. The code is structured considering SOLID principles
2. Didn't use any web frameworks like gin,echo etc because the usecase is simple
3. Chose postgres DB to store data, to simulate actual production scenario
4. Used Migrate library for handling sql scripts, its a flexible library can be moved into a cli as well
5. Used chi router for better route handling and middleware management
6. Stored discounts in a table, so that it can be easily modified at runtime without redeployment of app
7. Ideally it's better to also use redis, because we have to frequently hit discounts table
8. Used testify library, because it has good features and mock support
9. Used sqlx ORM for Database Queries, its a lightweight ORM which compliments the database/sql package
10. Added Docker support, so that its easily deployable
