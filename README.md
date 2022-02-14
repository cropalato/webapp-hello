# webapp-hello

## running the container from dockerhub
```bash
docker run --rm -p 3334:3000 -v $(pwd)/index.html:/app/index.html cropalato/webapp-hello:latest
```
