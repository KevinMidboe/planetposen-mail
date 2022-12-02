# Planetposen email sender

App for sending emails related to planetposen webshop. Currently only supports sending order confirmations emails after purchase.

## Setup
From source:
1. `git clone https://github.com/kevinmidboe/planetposen-mail`
2. `cp .env.example .env`
3. Update variables in `.env` file
4. `make install`

## Run
Run api from command line with:
```bash
make server
```

Run as docker container using:

```bash
(sudo) docker run -d --name planetposen-mail -p 8000:8000 \
    -e PORT=8000
    -e SEND_GRID_API_ENDPOINT="https://api.sendgrid.com"
    -e SEND_GRID_API_KEY="sg.nE3..."
```
or
```bash
(sudo) docker run -d --name planetposen-mail -p 8000:8000 --env-file .env planetposen-mail
```

## Preview
Use view template as preview during local development run:
```bash
make preview
```