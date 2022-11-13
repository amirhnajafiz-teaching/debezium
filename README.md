<h1 align="center">
    Henry
</h1>

Using _Debezium_ to publish _PostgresQL_ updates over _Kafka_.
For testing _Debezium_ I create a _PostgresQL_ client which gets 
http requests from users and inserts data into database. Meanwhile,
the _Debezium_ server listens over database and publishes the transactions
over _Kafka_ message broker.

## Start

First build the _golang_ application:

```shell
go build . henry
```

```shell
docker compose up -d
```

### Migration

```shell
./henry migrate
```

Executes the migration queries.

### Kafka

```shell
./henry kafka
```

Listens over kafka cluster.

### HTTP

```shell
./henry http
```

Accepts http requests:

```json
[
  {
    "url": "localhost:7490/api",
    "method": "post",
    "request": {
      "type": "application/json",
      "body": {
        "name": "[string] type",
        "email": "[string] type"
      }
    },
    "response": {
      "status code": "200",
      "status message": "OK",
      "message": "response message"
    }
  },
  {
    "url": "localhost:7490/api",
    "method": "get",
    "response": {
      "status code": "200",
      "status message": "OK",
      "body": {
        "id": "int",
        "name": "string",
        "email": "string"
      }
    }
  }
]
```

### Debezium

```shell
./setup.sh
```

Now you can consume events over _kafka_ client.

