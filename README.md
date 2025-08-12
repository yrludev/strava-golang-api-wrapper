
# Strava Golang API Wrapper

This project is a Go (Golang) client library and demo for interacting with the Strava API. It provides convenient methods to authenticate, fetch athlete data, activities, segments, clubs, gear, routes, uploads, and more, using idiomatic Go code.

## Features
- Authenticate with Strava using an access token
- Fetch athlete profile and stats
- List and get activities, comments, laps, and kudoers
- Explore and get segments
- List starred segments
- Get club and gear details
- Fetch routes, uploads, and activity streams
- Example usage in `main.go`

## Setup
1. **Clone the repository:**
	```sh
	git clone https://github.com/yrludev/strava-golang-api-wrapper.git
	cd strava-golang-api-wrapper
	```
2. **Install dependencies:**
	This project uses Go modules. Dependencies will be installed automatically when you build or run the project.

3. **Get a Strava API Access Token:**
	- Register an application at https://www.strava.com/settings/api
	- Follow Strava's OAuth process to obtain an access token

4. **Set the Access Token as an Environment Variable:**
	- On Windows PowerShell:
	  ```powershell
	  $env:STRAVA_ACCESS_TOKEN = "your_access_token_here"
	  ```
	- On Linux/macOS:
	  ```sh
	  export STRAVA_ACCESS_TOKEN=your_access_token_here
	  ```

## Running the Demo
Run the example program to see the API wrapper in action:

```sh
# In the project root directory
# Make sure STRAVA_ACCESS_TOKEN is set as described above

go run main.go
```

The demo will print your athlete info, recent activities, segments, clubs, gear, and more, using the Strava API.

## Notes
- The wrapper is a work in progress and may not cover every Strava API endpoint.
- You need a valid Strava access token for most API calls.
- See `main.go` for example usage and how to call each method.

## License
MIT License
