# About this repo
- Learning about how to use redis by using golang

# Service in This Project
- Redis: Using for cach data and getting high performance because redis use in memmory database for getting data
- InfluxDB: Database for save load test data from K6
- Grafana: Using for creating dashboard load test data from K6
- K6: Using load test API

# Start image Redis, InfluxDB, Grafana
- docker compose up redis influxdb grafana

# Stop and Remove container Redis, InfluxDB, Grafana
- docker compose down redis influxdb grafana

# Run Script load test with K6
- docker compose run --rm k6 run /scripts/test.js 