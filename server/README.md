# Golang temp by zipcode

A tool for get weather of location by zipcode

## Config

```bash
go mod tidy
```

### Running

You must configure your [Weather API](https://weatherapi.com/) key on `.env` file like the example `.env.example`

```bash
go run cmd/main.go
```
#### Or
```bash
docker compose up -d
```

The server will be running on port `:8080`

## Usage

### Local

```bash
curl http://localhost:8080/01153000/temperature
```

### GCP
```bash
curl https://golang-temp-by-zipcode-ea5pytw5bq-uc.a.run.app/01153000/temperature
```

The path parameter zipcode (01153000 on example), is mandatory