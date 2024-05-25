# Custom Prometheus Metrics with Go, Docker, and Grafana

This repository contains code to create custom gauge metrics using Go and exposes them to Prometheus. The setup includes a Docker Compose configuration that sets up Prometheus to scrape these custom metrics and Grafana to visualize them.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

Follow these steps to set up and run the architecture using Docker Compose.

### 1. Clone the Repository

```bash
git clone https://github.com/bhavishya-8/Prometheus_custom_metrics_go.git
```

### 2. Build and Run with Docker Compose

Use the following command to build the Docker images and start the containers:

```bash
docker-compose up --build
```
This command will start three containers one of the go code container, a prometheus container at 9090 port and a Grafana container at 3000 port.

### 3. Access Prometheus and Grafana

- Prometheus will be available at [http://localhost:9090](http://localhost:9090)
- Grafana will be available at [http://localhost:3000](http://localhost:3000)

### 4. Configure Grafana

1. Open your browser and navigate to [http://localhost:3000](http://localhost:3000).
2. Log in with the default credentials (`admin`/`devops123`). You will be prompted to change your password.
3. Add Prometheus as a data source:
    - Go to **Configuration** > **Data Sources**.
    - Click **Add data source** and select **Prometheus**.
    - Set the URL to `http://prometheus:9090`.
    - Click **Save & Test** to verify the connection.
4. Import the provided Grafana dashboard (if any) or create a new dashboard to visualize your custom metrics.
