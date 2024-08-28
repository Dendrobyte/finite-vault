
## Installation:

- Run `go install github.com/air-verse/air@latest`
- Depending on your local machine, you need to make changes to the `.air.toml` file-- specifically the `bin = "tmp\main.exe"`
	- Mac OS: it should be `bin = "tmp\main.exe"`
	- Windows: `bin = "tmp\\main.exe"`

  
  

## Running:
- Add this alias to your Bash profile: `alias air='$(go env GOPATH)/bin/air'`
- Make sure you have the `.env` file and it's in the root of `backend-go`

  
  

Documentation:
- https://github.com/air-verse/air