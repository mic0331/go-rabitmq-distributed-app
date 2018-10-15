# go-rabitmq-distributed-app
Experiment with Go and Rabitmq

This section details how to start the whole cluster using `docker-compose` and a YAML definition file

1. Create a network shared by all containers
```bash
docker network create rabbitmq-cluster
```

2. Start cluster:
```bash
docker-compose up -d
```

3. View logs for all containers
```bash
docker-compose logs -f
```
