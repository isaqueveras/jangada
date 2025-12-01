package newapp

const tmplDockerfile = `FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o app ./cmd/app

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE {{ .DefaultHost }}

ENTRYPOINT ["./app"]`

const tmplPrometheus = `global:
  scrape_interval: 5s
  evaluation_interval: 5s

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          # - alertmanager:9093

rule_files:

scrape_configs:
  - job_name: {{ ToLower .AppName }}
    scrape_interval: 10s
    static_configs:
      - targets: ["{{ ToLower .AppName }}_app:{{ .DefaultHost }}"]

  - job_name: pushgateway
    scrape_interval: 10s
    static_configs:
      - targets: ["pushgateway:9091"]`

const tmplDockerCompose = `services:
  {{ ToLower .AppName }}_postgres:
    image: postgres:15
    restart: no
    container_name: {{ ToLower .AppName }}_postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB={{ ToLower .AppName }}
    ports:
      - "5432:5432"
    volumes:
      - {{ ToLower .AppName }}_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 5s
      timeout: 5s
      retries: 5
    networks:
      - monitor

  {{ ToLower .AppName }}_app:
    build:
      context: .
      dockerfile: ./Dockerfile.app
    container_name: {{ ToLower .AppName }}_app
    ports:
      - "{{ .DefaultHost }}:{{ .DefaultHost }}"
    environment:
      - CGO_ENABLED=0
      - {{ ToUpper .AppName }}_DATABASE_HOST={{ ToLower .AppName }}_postgres
      - APP_ADDRESS={{ .DefaultHost }}
    env_file:
      - .env
    restart: no
    depends_on:
      - {{ ToLower .AppName }}_postgres
    networks:
      - monitor

  {{ ToLower .AppName }}_prometheus:
    image: prom/prometheus:latest
    container_name: {{ ToLower .AppName }}_prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    restart: no
    networks:
      - monitor

  {{ ToLower .AppName }}_pushgateway:
    image: prom/pushgateway:latest
    container_name: {{ ToLower .AppName }}_pushgateway
    ports:
      - "9091:9091"
    restart: no
    networks:
      - monitor

  {{ ToLower .AppName }}_grafana:
    image: grafana/grafana:latest
    container_name: {{ ToLower .AppName }}_grafana
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_USERS_ALLOW_SIGN_UP=false
    restart: no
    networks:
      - monitor

networks:
  monitor:
    driver: bridge

volumes:
  {{ ToLower .AppName }}_data:`
