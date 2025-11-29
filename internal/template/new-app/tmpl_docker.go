package newapp

const tmplDockerfile = `FROM golang:1.24 AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY ../ ./

RUN go build -o server ./cmd/app

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

ENTRYPOINT ["./server"]`

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
      - targets: ["{{ ToLower .AppName }}_app:8080"]

  - job_name: pushgateway
    scrape_interval: 10s
    static_configs:
      - targets: ["pushgateway:9091"]`

const tmplDockerCompose = `services:
  pg_{{ ToLower .AppName }}:
    image: postgres:15
    restart: no
    container_name: pg_{{ ToLower .AppName }}
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: {{ ToLower .AppName }}
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

  app_{{ ToLower .AppName }}:
    build:
      context: ..
      dockerfile: ./docker/Dockerfile
    container_name: app_{{ ToLower .AppName }}
    ports:
      - "8080:8080"
    environment:
      - CGO_ENABLED=0
    env_file:
      - ../.env
    restart: no
    depends_on:
      - prometheus
      - pg_{{ ToLower .AppName }}
    networks:
      - monitor

  prometheus:
    image: prom/prometheus:latest
    container_name: prometheus_{{ ToLower .AppName }}
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    restart: no
    networks:
      - monitor
    depends_on:
      - pushgateway

  pushgateway:
    image: prom/pushgateway:latest
    container_name: pushgateway_{{ ToLower .AppName }}
    ports:
      - "9091:9091"
    restart: no
    networks:
      - monitor

  grafana:
    image: grafana/grafana:latest
    container_name: grafana_{{ ToLower .AppName }}
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_USERS_ALLOW_SIGN_UP=false
    depends_on:
      - prometheus
    restart: no
    networks:
      - monitor

networks:
  monitor:
    driver: bridge

volumes:
  {{ ToLower .AppName }}_data:`
