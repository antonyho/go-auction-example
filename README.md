# Go API Server Example as Auction Bid Tracker

This is an example server for auction bid tracker.

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.  

### Running the server
To run the server, follow these simple steps:

```
make run
```

To run the server in a docker container
```
make image
```

Once image is built use
```
docker run --rm -it -p 127.0.0.1:8080:8080 antonyho-auction-api-example
```