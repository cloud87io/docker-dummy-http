# Dummy HTTP Server

This will help test server deployments and networking setups.

## Usage
```shell
docker run --rm -e PORT=8080 -p 8080:8080 ghcr.io/cloud87io/docker-dummy-http
```

## Environment Variables
* `PORT` - set the port of the HTTP server. Defaults to `8080`

## Endpoints

| Path  | Description |
| ------------- | ------------- |
| `/health`  | Health check.  |
| `/headers`  | Returns the headers that were provided in the request  |
| `/error/###`  | Responds with the HTTP error code that was requested  |
| `/exit` | Will exit the server |
| `/panic` | Will panic print error logs |
| `/`  | Home page  |
| _catchall_ | Same as `/` |
