# go-pusher-example

> A simple example of using Pusher in a Go server

* [server](server) - a Go HTTP server that triggers events using Pusher
* [client](client) - a simple web page renders events in a web UI

## Starting the server

```
export PUSHER_APP_ID=...
export PUSHER_KEY=...
export PUSHER_SECRET=...
export PUSHER_CLUSTER=...
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
open index.html
```

## Sending an event

```bash
curl -X POST http://127.0.0.1:8000/events -H 'Content-Type: application/json' -d '{"foo": "bar"}'
```
