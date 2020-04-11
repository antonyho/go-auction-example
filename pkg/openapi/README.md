# Go API Server for openapi

This is an example server for auction bid tracker.

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.  
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0.0
- Build date: 2020-04-11T21:14:03.879Z[GMT]
For more information, please visit [https://antonyho.net/](https://antonyho.net/)


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi 
```

