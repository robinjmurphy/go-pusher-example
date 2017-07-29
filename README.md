# pusher-example

> A simple example of using Pusher

* [server](server) - a Go HTTP server that triggers events using Pusher
* [client](client) - a React application that renders events in a web UI

## Starting the server

```
cd server
go get
go run main.go
```

You can also run the server in live-reload mode using `gin`:

```
go get github.com/codegangsta/gin
gin run main.go
```

## Starting the client

```
cd client
npm install
npm start
```

## Sending an event

```bash
curl -X POST http://127.0.0.1:8000/events -H 'Content-Type: application/json' -d '{"foo": "bar"}'
```
