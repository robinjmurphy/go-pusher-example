# go-pusher-example

> A simple example of using Pusher in a Go server

## Starting the server

```
export PUSHER_APP_ID=...
export PUSHER_KEY=...
export PUSHER_SECRET=...
export PUSHER_CLUSTER=...
go get
go run main.go
```

You can also run the server in live-reload mode using `gin`:

```
go get github.com/codegangsta/gin
gin run main.go
```

Visit [http://127.0.0.1:8000](http://127.0.0.1:8000) in your browser.

