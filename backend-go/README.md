# Finite Vault - Go Backend

A go backend. I think this is a skill I want to work on but leaning into JS-heavy frameworks... so golang TBD.

## Contributing

Please see the parent directory README for more information here.

## Installation

You'll need to ensure you install the go libraries with `go install`. Additionally, for development at least, I would recommend installing [air](https://github.com/air-verse/air), which is effectively just live reload.

```bash
go install github.com/air-verse/air@latest
```

When running,

- Add this alias to your Bash profile: `alias air='$(go env GOPATH)/bin/air'`
- Make sure you have the `.env` file and it's in the root of `/backend-go`

### Docker

To run Docker, first build the container with

```bash
docker build -t finite-vault .
```

and then run the container with

```bash
docker run -t --rm finite-vault
```

TODO: Proper Docker compose to set up allowed frontend hosts and stuff for production

## Env Variables

Hit up Jellyfishers for the `.env` file(s). The backends should all use the same one and the frontend has one, so you'll need two.

## Documentation

See Postman for api documentation. Mark TODO: Write up documentation on the ideas and ethos of _Infinite Game_ as a whole.
