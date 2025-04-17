# Build Version

This tool helps to generate specification compliant version tags that is used mainly with Trunk-based Branching strategy.

## Requirements
build-version communicates with the specified Git Service (eg. GitHub, GitLab) to fetch labels to decide whether to increment the major, minor or patch version. As such, the following requirements must be met for a successful run:
- Server running build-version must allow HTTPS egress to the specified Git Service
- Token must be provided to communicate with the specified Git Service. build-version by default will try to use known default environment variables if they are not provided.
- Full repo checkout is required for build-version to introspect the git history. By default, checkouts in CI such as GitHub Actions (using actions/checkout@v4) only checks out a single commit.

To ensure the requirements are met, please follow the setup guide in the next section closely.

## How to Use

### Method 1: Docker Container
You can run build-version by using the published docker container at `ghcr.io/alacrity-sg/build-version:latest`.
An example run would be as shown below. Replace the variables accordingly.

```bash
docker run -v /my/repo/path:/my-repo -e TOKEN=my-gh-token -e REPO_PATH=/my-repo ghcr.io/alacrity-sg/build-version:latest
```

### Method 2: Binary
build-version publishes complied binaries for multiple platforms on the [releases](https://github.com/alacrity-sg/build-version/releases). You can use this tool by downloading the right binary for your platform/os and run it as shown below.

```bash
curl -Lo build-version <release url>
chmod +x ./build-version
./build-version -repo-path=/my/repo/path -token="${GITHUB_TOKEN}"
```

### Method 3: Source
If you would like to run using the source, follow the steps shown below.
```bash
brew install go # Download Go if you do not have it already
git clone https://github.com/alacrity-sg/build-version.git build-version-source
cd build-version-source
go mod download
go run main.go -repo-path=/my/repo/path -token="${GITHUB_TOKEN}"
```

### Method 4: Using GitHub actions
If you are planning to run build-version from GitHub Actions, you can use our published action script as shown below.

```yaml
permissions:
  pull-requests: read
  contents: write # Set to read if you do not want build-version to create release/tags for you
jobs:
  my-job:
    runs-on: ubuntu-latest
    outputs:
      BUILD_VERSION: ${{ steps.build-version.outputs.BUILD_VERSION }}
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0 # Important! build-version requires full git log
      - uses: alacrity-sg/build-version
        id: build-version
        with:
          repo-path: ./ # Optional Input, defaults to ./
          token: ${{ secrets.GITHUB_TOKEN }} # Required Input, can be a PAT, actions token or installation token
  subsequent-job: # This portion is an example to show how you can use the generated version in subsequent jobs
    runs-on: ubuntu-latest
    needs: [my-job] # Required to set as needs to get output
    env:
      BUILD_VERISON: ${{ needs.my-job.outputs.BUILD_VERSION }} # Set as env for easier access
    steps:
      - uses: actions/checkout@v4
      - name: Check token
        runs: echo "My Build Version is ${BUILD_VERSION}"
