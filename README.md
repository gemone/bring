# Bring

Go client library for [Apache Guacamole](http://guacamole.apache.org) Protocol.

Fork from [`deluan/bring`](https://github.com/deluan/bring)

## Feature

TODO

## TODO

- [ ] Improve performance
  - [ ] Optimize communication process
  - [ ] Import log Library [uber-go/zap][2]
- [ ] Perfect realization
  - [ ] Websocket channel
  - [ ] Perfect stream processing
- [ ] Newer
  - [ ] Close to new standards
  - [ ] Use new treatment scheme
  - [ ] More specific documentation

## Quick start (tl;dr)

1. Install the library in your project:
   `go get github.com/gemone/bring`
2. Create a [Client](https://pkg.go.dev/github.com/gemone/bring#Client)
   with the [NewClient()](https://pkg.go.dev/github.com/gemone/bring#NewClient) function.
   This creates a session with the specified `guacd` server
3. Start the client with `go client.Start()`
4. Get screen updates with `client.Screen()`
5. Send keystrokes with `client.SendKey()`
6. Send mouse updates with `client.SendMouse()`

See the [sample app](sample/main.go) for a working example

## Documentation

The API is provided
by the [Client](https://pkg.go.dev/github.com/gemone/bring#Client) struct.
The [documentation](https://pkg.go.dev/github.com/gemone/bring) is a work in progress,
but the API is very simple and you can take a look at all features available in the
[sample app](sample) provided. Here are the steps to run the app:

1. You'll need a working `guacd` server in your machine.
   The easiest way is using docker
   and docker-compose. Just call `docker-compose up -d` in the root of this project.
   It starts the `guacd` server and a sample headless linux with a VNC server
2. Run the sample app with `make run`. It will connect to the linux container
   started by docker.

Take a look at the Makefile to learn how to run it in different scenarios.

## Why?

Apache Guacamole was created with the goal of making a dedicated client unnecessary.
So why create a client?!

The idea is that if you need to control a remote machine from your Go code, you can
leverage the Guacamole protocol and the `guacd` server as a bridge. This way you
can use any protocol supported by Guacamole (currently RDP and VNC,
with X11 coming in the future) to do screen capture and remote control of
networked servers/desktop machines from within your Go app.

It seems that the [project `deluan/bring`][1] is no longer maintained
and there is no relevant plan.
Therefore, try to improve the protocol to expand its use.

## References

- [The Guacamole protocol](http://guacamole.apache.org/doc/gug/guacamole-protocol.html)
- [Guacamole protocol reference](http://guacamole.apache.org/doc/gug/protocol-reference.html#rect-instruction)
- [Apache Guacamole Client implementation](https://github.com/apache/guacamole-client/tree/master/guacamole-common-js)
- [deluan/bring][1]
- [dushixiang/next-terminal](https://github.com/dushixiang/next-terminal)

[1]: https://github.com/deluan/bring
[2]: https://github.com/uber-go/zap
