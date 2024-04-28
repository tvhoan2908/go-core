## Descriptions
- Project for ELearning

## Requirements
- golang 1.22 or higher.
- postgresql
- redis
- Nodemon

## Getting started
```sh
$ go get -u && go mod tidy
$ make run-watch
```

## Migrate database
- Using with goose
```sh
$ make migration create_demo_table
$ make db-up
$ make db-down
```