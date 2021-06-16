# GitHub Stats

A simple tools to get GitHub pull requests statistics written in Go.

## Requirements

It needs [GitHub Personal Access Token](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token) to access the repo and organisations.

## Build

```shell
go build -o github-stats
```

It will create `github-stats` command.

## Usage

```shell
./github-stats
```

### Get Pull Requests Statistics

```shell
./github-stats pr --token {yourtoken}
```

For getting more info, use `--help` in the command.

## Contributing

Pull requests are welcome.

## License

MIT License, please refer to [LICENSE](LICENSE.md) file.