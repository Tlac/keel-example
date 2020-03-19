# Push workflow example with Keel and Webhook Relay

## Build image

```
docker build -t docker.io/keelhq/push-workflow-example:latest .
docker push docker.io/keelhq/push-workflow-example:latest
```

To try it locally:

```
docker run -p 8500:8500 keelhq/push-workflow-example
```

## Create Kubernetes deployment

