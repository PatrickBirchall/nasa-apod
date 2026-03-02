# NASA APOD

A small Go CLI that downloads the latest NASA Astronomy Picture of the Day (APOD) and saves the image locally.

## Requirements

- **Go**: 1.26 or later (as specified in `go.mod`)
- **NASA API key**: the `NASA_KEY` environment variable must be set
- **Network access**: to reach the NASA APOD API

## Quick start

From the project root:

```bash
export NASA_KEY=your_api_key_here

make build           # builds ./bin/nasa-apod
./bin/nasa-apod      # downloads the latest APOD into ./images
```

If you prefer not to use `make`, you can build directly with:

```bash
go build -o bin/nasa-apod
./bin/nasa-apod
```

## Configuration

The CLI is configured via a mix of environment variables and flags:

- **Environment**
  - **`NASA_KEY`** (required): your NASA API key. If this is not set, the program exits with an error.

- **Flags**
  - **`-api-url`**: NASA APOD API base URL  
    - Default: `https://api.nasa.gov/planetary/apod`
  - **`-out`**: output directory for downloaded images  
    - Default: `images`

Examples:

```bash
# Use a custom output directory
./bin/nasa-apod -out downloads

# Override the API base URL (e.g. for testing or a proxy)
./bin/nasa-apod -api-url http://localhost:8080/apod
```

When run successfully, the program:

- Fetches metadata for the latest APOD from the NASA API
- Prints the image title to stdout (e.g. `Downloading image: The Milky Way`)
- Downloads the high‑resolution image to the output directory (default `images/`)
- Saves the file as `<title>.jpg`, with spaces in the title replaced by underscores

## Development

Useful commands while developing:

- **Format code**: `make fmt`
- **Run vet**: `make vet`
- **Build binary**: `make build`
- **Run tests**: `go test ./...`

The repository also includes GitHub Actions workflows for linting and testing/building to help keep the project healthy.