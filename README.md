# Postal Code API

> An API that allows searching in Switzerlands Postal Codes

## Features

* Caches Responses for 30 Minutes
* Compresses Responses with GZIP
* Fast JSON-Encoding

## Run Locally

```bash
air
```

## Generate Documentation (Swagger)

Run this command

```bash
swag init
```

### Problems in ASDF

RUn this first

```bash
asdf reshim
```

## Build Docker Image Locally

```bash
docker build -t plz-api .
```

### Run the Image

Database needs to be running!

```bash
docker run --rm -p 3000:3000 --name plz_api -e DB_CONN="<DB_CONN>" plz-api
```

## Todo

- [ ] Consider upping the time data stays in cache, because source only updates every month... Maybe to 1 Month?
- [ ] ...