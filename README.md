
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


## Running Examples

### 1. Demo: Fetch Athlete Info, Activities, Segments, etc.

```sh
# In the project root directory
# Make sure STRAVA_ACCESS_TOKEN is set as described above
go run main.go
```

The demo will print your athlete info, recent activities, segments, clubs, gear, and more, using the Strava API.

### 2. Create an Activity (like a direct API call)

You can create a new activity using command-line flags, similar to a direct API POST:

```powershell
$env:STRAVA_ACCESS_TOKEN="your_token_here"; go run main.go --create-activity --activity-name="Morning Ride" --activity-type="Ride" --sport-type="Cycling" --start-date-local="2025-08-12T07:00:00Z" --elapsed-time=3600 --description="Test ride" --distance=20000 --trainer=0 --commute=0 --json
```

Or on Linux/macOS:

```sh
export STRAVA_ACCESS_TOKEN=your_token_here
go run main.go --create-activity --activity-name="Morning Ride" --activity-type="Ride" --sport-type="Cycling" --start-date-local="2025-08-12T07:00:00Z" --elapsed-time=3600 --description="Test ride" --distance=20000 --trainer=0 --commute=0 --json
```

This will create the activity and print the result as JSON. All required fields must be provided.

## Notes
- The wrapper is a work in progress and may not cover every Strava API endpoint.
- You need a valid Strava access token for most API calls.
- See `main.go` for example usage and how to call each method.

## License
MIT License
