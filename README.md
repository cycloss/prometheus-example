# Prometheus Demo

This repo demonstrates how the following technologies can be used together to create a web app with monitoring and alerts:

- Nginx
- Prometheus
- Prometheus Alert Manager
- Prometheus Black Box Exporter
- Go
- Prometheus exporter for Go
- Grafana
- Docker
- Docker Compose

The aim is to make it as simple as possible from a beginners point of view.

## Usage

- Run with `docker compose up -d`
- Make a load of curl requests to endpoints like: `curl 'localhost:8080/sneed/feed'` and `curl 'localhost:8080/sneed/feed'`
- Visit `localhost:3000` to see the metric dashboards in Grafana
- If you visit `localhost:8080`, nginx will serve up the static `index.html` page, and at `localhost:8080/test` it will serve a test page

## Other

- The nginx config is configured to only allow curl, postman and mozilla user agents to show how bots can be fobbed off
