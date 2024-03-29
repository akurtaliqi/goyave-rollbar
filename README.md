# How to use Rollbar inside a Goyave project to get error monitoring in real time

See the article I wrote here [Medium](https://ajtene-kurtaliqi.medium.com/go-programming-how-to-use-rollbar-inside-a-goyave-project-to-get-error-monitoring-in-real-time-667d65259b6c).

# Goyave Template

A template project to get started with the [Goyave](https://github.com/System-Glitch/goyave) framework.

## Getting Started

### Requirements

- Go 1.13+
- Go modules

### Running the project

First, make your own configuration for your local environment. You can copy `config.example.json` to `config.json`.

Run `go run main.go` in your project's directory to start the server, then try to request the `hello` route.
```
$ curl http://localhost:8080/hello
Hi!
```

There is also an `echo` route, with a basic body validation.
```
$ curl -H "Content-Type: application/json" -X POST -d '{"text":"abc 123"}' http://localhost:8080/echo
abc 123
```

## Learning Goyave

The Goyave framework has an extensive documentation covering in-depth subjects and teaching you how to run a project using Goyave from setup to deployment.

<a href="https://goyave.dev/guide/installation"><h3 align="center">Read the documentation</h3></a>

<a href="https://pkg.go.dev/goyave.dev/goyave/v3"><h3 align="center">pkg.go.dev</h3></a>

## Contributing

Thank you for considering contributing to the Goyave framework! You can find the contribution guide in the [documentation](https://goyave.dev/guide/contribution-guide.html).

I have many ideas for the future of Goyave. I would be infinitely grateful to whoever want to support me and let me continue working on Goyave and making it better and better.

You can support me on Github Sponsor or Patreon.

<a href="https://github.com/sponsors/System-Glitch">❤ Sponsor me!</a>

<a href="https://www.patreon.com/bePatron?u=25997573">
    <img src="https://c5.patreon.com/external/logo/become_a_patron_button@2x.png" width="160">
</a>

## License

The Goyave framework is MIT Licensed. Copyright © 2019 Jérémy LAMBERT (SystemGlitch)
