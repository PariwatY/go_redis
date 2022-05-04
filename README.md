# About this repo
- Learning about how to use redis by using golang

# Service in This Project
- Redis: Using for cach data and getting high performance because redis use in memmory database for getting data
- InfluxDB: Database for save load test data from K6
- Grafana: Using for creating dashboard load test data from K6
- K6: Using load test API

# Preparing Docker image
```
1. docker pull redis
2. docker pull loadimpact/k6
3. docker pull influxdb:1.8.10
4. docker pull grafana/grafana
```

# Preparing CLI
```
brew install redis
```
# Step to start project
1. Start Docker Image for Redis, InfluxDB, Grafana with the following command
```
 docker compose up redis influxdb grafana
```

2. Start Go Service 
```
go run main.go
```

3. Run Script load test with K6
```
docker compose run --rm k6 run /scripts/test.js 
```

4. If you want to stop services Stop  Redis, InfluxDB, Grafana, you can use the following command
```
docker compose down redis influxdb grafana
```


# Checklist URL Grafana and InfluxDB
  - http://localhost:3000/    -> Grafana
  - http://localhost:8086/    -> InfluxDB


# Note
 - host.docker.internal will use when you want to call service that run with docker image if you use localhost will cannot call service in docker image 