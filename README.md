# vindinium-starter-go [![Build Status](https://drone.io/github.com/geetarista/vindinium-starter-go/status.png)](https://drone.io/github.com/geetarista/vindinium-starter-go/latest)

Go starter bot for [Vindinium](http://vindinium.org), based on the [Python starter](https://github.com/ornicar/vindinium-starter-python).

## Prerequisites

Before using this starter, you need to [install Go](http://golang.org/doc/install). It is outside of the scope of this documentation to get your environment, but there are plenty of resources available online if you need help.

Then just get this repository or create your own fork and get it:

```bash
go get github.com/geetarista/vindinium-starter-go
```

## Getting Started

This starter gives you everything you need to start playing with Vindinium except for a bot. There are a couple [example bots](vindinium/bot.go), but these are just stubs. Your goal is to build a bot or bots that you can use to compete on Vindinium. Either flesh out the FighterBot or create a new bot with your logic for competing. When adding a new bot, ensure it is properly linked in the `Setup` method in [client.go](vindinium/client.go).

Once you've added a bot, all you need to do is build run `make` to create a client executable. If you are adding tests for your bot, you can run them with `make test`.

## Usage

To use the vindinium client, just call the executable and pass in the flags you need. The only required flag is `-k`, which is your API key:

```bash
./client -k abc123
```

The remaining flags are optional:

`-m` This is the game mode. Choices are `arena` and `training`. Default: `arena`

`-b` The name of your bot. This name is defined in [client.go](vindinium/client.go) and must set up the associated bot. Default: `fighter`

`-c` The count of games (arena) or turns (training) to play. Default: `1`

`-r` Use a random map in training mode. Setting to `false` can be useful to ensure consistency in training. Default: `true`

`-s` The server to play against. Default: `http://vindinium.org`

`-d` Debug mode&mdash;verbose output for debugging and testing. Default: `false`

## API Documentation

You can [view the documentation on Godoc](http://godoc.org/github.com/geetarista/vindinium-starter-go) to get a better feel for the package internals.

## Contributing

If you want to help improve this starter or report issues, please read the [contributing guidelines](CONTRIBUTING.md) first.

## License

MIT. See [LICENSE](LICENSE).
