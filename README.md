# go-monitoring
A very basic monitoring server, written in Go. Why, because I needed a
monitoring server which, didn't use a lot of ram, and was easy to use with a
small setup.

What is not a lot of ram? Example setup uses less than 12 MB of RAM inside the
given Docker container.

Documentation of the config file can be found in the docs folder: [docs/config_format.md](docs/config_format.md)

## Build the server
```
docker build -t monitoringserver .
```

## Run the server after building
After running Gotify, registering an app in it, getting a token, and updating the config.yaml, run:
```
docker run -p 3000:3000 -v ./example/example.yaml:/app/example/config.yaml -e CONFIG_PATH=/app/example/config.yaml monitoringserver
```

## Run and build with docker compose
```
docker compose up gotify
```
Then register an application in Gotify and get the token, set this in the config and run:
```
docker compose up gomonitoring
```
