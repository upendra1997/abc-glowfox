# ABC Glofox

## Setup

1. ensure you have [golang](https://go.dev/dl/) setup and accessible through your command line.
2. run `go get .` to get all the dependencies of the project.
3. run `go run .` to start the server on port `8080`, you can change the port in `config/config.go`

## Test
1. Run tests using `go test`

## Sample APIs

### Classes

#### Create

```bash
curl --location --request POST 'localhost:8080/api/classes' \
--header 'Content-Type: application/json' \
-d '
{
  "name": "pilates",
  "start_date": "2024-12-01",
  "end_date": "2024-12-05",
  "capacity": 20
}
'
```

#### Get

```bash
curl --location --request GET 'localhost:8080/api/classes' \
--header 'Content-Type: application/json'
```

```json
[
  {
    "name": "pilates",
    "date": "2024-12-01",
    "capacity": 20
  },
  {
    "name": "pilates",
    "date": "2024-12-02",
    "capacity": 20
  },
  {
    "name": "pilates",
    "date": "2024-12-03",
    "capacity": 20
  },
  {
    "name": "pilates",
    "date": "2024-12-04",
    "capacity": 20
  },
  {
    "name": "pilates",
    "date": "2024-12-05",
    "capacity": 20
  }
]
```

### Bookings

#### Book

```bash
curl --location --request POST 'localhost:8080/api/booking/upendra' \
--header 'Content-Type: application/json' \
-d '
{
  "name": "pilates",
  "date": "2024-12-03"
}
'
```

#### Get

```bash
curl --location --request GET 'localhost:8080/api/booking' \
--header 'Content-Type: application/json'
```

```json
[
  {
    "user_name": "upendra",
    "class_name": "pilates",
    "date": "2024-12-03"
  }
]
```

#### Side Effect

one of the classes' capacity got reduced

```json
[
  {
    "name": "pilates",
    "date": "2024-12-01",
    "capacity": 20
  },
  {
    "name": "pilates",
    "date": "2024-12-02",
    "capacity": 20
  },
  {
    "name": "pilates",
    "date": "2024-12-03",
    "capacity": 19
  },
  {
    "name": "pilates",
    "date": "2024-12-04",
    "capacity": 20
  },
  {
    "name": "pilates",
    "date": "2024-12-05",
    "capacity": 20
  }
]
```
