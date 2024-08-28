# Finite Vault - Go Backend

A go backend. I think this is a skill I want to work on but leaning into JS-heavy frameworks... so golang TBD.

## Setup

You'll need to ensure you install the go libraries with `go install`. Additionally, for development at least, I would recommend installing [air](https://github.com/air-verse/air), which is effectively just live reload.

```bash
go install github.com/air-verse/air@latest
```

## OS

If you are on Mac or Linux, in the [air configuration file](./.air.toml), make sure the `bin` path is set to

```bash
bin = "tmp\main.exe"
```

whereas on Windows you'll need

```bash
bin = "tmp\\main.exe"
```

Until we get Docker set up, that's what will need to be done for now because of the path differences.

## Env Variables

Hit up Mark or someone else for the `.env` file(s). The backends should all use the same one and the frontend has one, so you'll need two.

## Documentation

See Postman for api documentation. Mark TODO: Write up documentation on the ideas and ethos of _Infinite Game_ as a whole.
