# Go-Proxy

It offers an interface for configuring an HTTP server that mounts domains as they were it's middleware.

This helps for example when we are developing an application that has it's api on a different domain and we don't want to enable CORS by default.

# Configuration

We need to create a ```servers.json``` file. (by default)

or we can specify the file with ```--config my_servers.json```

```json
{
  "port": 8080,
  "servers": [
    {
      "mount": "/app",
      "host": "my-app",
      "port": 3000
    },
    {
      "mount": "/*",
      "host": "my-api",
      "port": 7003
    }
  ]
}
```

```go-proxy --config server_config.json``` will start a server on port 8080 displaying the following information
```
Proxy: /app  ->  http://localhost:3000
Proxy: /google  ->  http://www.google.com:80
Proxy: /*  ->  http://soa-server:7003
```

It means that any request sent over /app will be redirected to localhost:3000 for example.
