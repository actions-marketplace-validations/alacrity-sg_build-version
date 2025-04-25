# Build Version

`build-version` is a nifty tool to help generate semver compliant version strings that can be used with your builds whether it is in a CICD workflow or just locally.
`build-version` works by analyzing tags and branches in your repository to figure out 3 types of versioning: build version, release candidate version and release versions.


## How to Use
`build-version` accepts the following configuration as inputs to customise how the tool runs.
| Input Name     | Input Type | Required | Default Value       | Description                                                                                   |
|----------------|------------|----------|---------------------|-----------------------------------------------------------------------------------------------|
| repo-path      | string     | False    | ./                  | Path to your code repository                                                                  |
| token          | string     | False    | Empty               | PAT or API Token to communicate with your git service. Not required if offlineMode is enabled |
| offline        | boolean    | False    | False               | If enabled, build-version will determine build versions without using remote APIs             |
| increment-type | string     | False    | patch               | Accepts 'patch', 'minor' or 'major'. If provided, value will be used to increment the version |
| output-file    | string     | False    | ./build-version.env | Output file containing the generated version. Generated content uses .env file format         |

### Method 1: Using GitHub actions
The easiest way to use build-version is by using the published action in your GitHub Actions Workflow. Copy the configuration below and add it to your workflow file.

```yaml
permissions:
  pull-requests: read # Must be provided! build-version will throw an error during validation if this is not provided.
  contents: read # Must be provided! build-version will throw an error during validation if this is not provided.
jobs:
  my-job:
    runs-on: ubuntu-latest
    outputs:
      BUILD_VERSION: ${{ steps.build-version.outputs.BUILD_VERSION }}
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0 # Important! build-version requires full git log
          fetch-tags: true # Important! build-version requires tags to determine versions
      - uses: alacrity-sg/build-version@v1
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
```
### Method 2: Binary
Download the latest binaries from [releases page](https://github.com/alacrity-sg/build-version/releases). You may use the following script to download the latest version.

```bash
curl -Lo build-version https://github.com/alacrity-sg/build-version/releases/download/v1.0.1/build-version_linux_amd64
chmod +x ./build-version
mv build-version /usr/local/bin
export PATH="${PATH}:/usr/local/bin" # Add to path if this has not been done yet.
```

With the download complete, you can now run the tool using the following command.
```bash
# If you want to run using defaults
export GITHUB_TOKEN="my-pat-token"
build-version -token="${GITHUB_TOKEN}"

# Or run using offline mode if you don't need remote api support
build-version -offline

# Or run with customised inputs
build-version -token="${GITHUB_TOKEN}" -repo-path=/my/repo/path -output-file=./my-build-version.env
```

### Method 3: Source
You can also run build-version tool using the source code. You may use the script below to download and run build-version.
```bash
brew install go # Download Go if you do not have it already
git clone https://github.com/alacrity-sg/build-version.git build-version-source
cd build-version-source
go mod download
go run main.go -repo-path=/my/repo/path -token="${GITHUB_TOKEN}"
```
