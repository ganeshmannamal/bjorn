# bjorn

## Image comparison tool for Bjorn

**bjorn allow the users (like Bjorn) to compare images that are provided as a list in a csv file**

# Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
    * [Project Structure](#project-structure)
    * [Image Comparison Algorithm](#image-comparison-algorithm)
    * [Further Development](#further-development)
    * [Running Tests](#tests)
    
## Installation
### [LATEST RELEASE](https://github.com/ganeshmannamal/bjorn/releases/latest)
Download the appropriate release package as per you operating system. Supported operating systems are:
 * MacOS (Darwin)
 * Windows
 * Linux
 
Unpack the downloaded archive in you local directory to install the `bjorn` cli tool.

## Usage
```
Usage:
  bjorn [command]

Available Commands:
  diff        compare images listed in a csv file
  help        Help about any command
```

### bjorn diff

compare images listed in a csv file.

```
bjorn diff [flags]
```

### Options

```
  -f, --file string   CSV file to read image list
  -h, --help          help for diff
  -o, --out string    Output file location
      --alsologtostderr                  log to standard error as well as files
      --config string                    config file (default is $HOME/.bjorn.yaml)
  -h, --help                             help for bjorn
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

## Development

### Project structure
We are using Cobra (a library for creating CLI applications using Go) to build this project. This allows to quickly prototype the solution as a user friendly CLI application and makes future extension with more commands easier.

For release management we are using GoReleaser along with TravisCI when a new tag is pushed. It generates binaries for each of the supported operating systems (MacOs, Windows, Linux).

### Image Comparison Algorithm
Bjorn uses a simple pixel based approach to image comparison. Each pixel of the 2 images being compared using the RGB values of each pixel.
This is a relatively simple approach to image comparison, suitable for basic percentage difference comparison.

The Score is calculated by [summing up the difference in RGB values for each pixel](https://github.com/ganeshmannamal/bjorn/blob/master/pkg/pair/pair.go#L48) and using this value to get the ratio of different pixels to total pixels in the images.
To calculate the time taken for each comparison we called a deferred function [elapsed()](https://github.com/ganeshmannamal/bjorn/blob/master/pkg/pair/pair.go#L76) at start of each `Compare()` call.
 
Limitations of current approach:
 * It assumes the images are of the same size. Scaled versions of the same image will be considered different.
 * Comparing same images in different formats (png, jpg, gif) can produce a score of ~0.01 to 0.02. The comparison algorithm uses a threshold of 0.01, below which images are considered similar.
 
### Further Development
The comparison algorithm is defined as the `Compare()` function in the [pair](https://github.com/ganeshmannamal/bjorn/blob/master/pkg/pair/pair.go) package. This function may be extended to improve the comparison algorithm.
Possible improvements include:
 * Improve comparison algorithm to consider scaled images.
 * Add ability to define multiple comparison algorithm and flags to select each.
 * Add command to compare 2 images separately (eg `bjorn diff images img1 img2`).
 
### Tests
Run tests using `go test ./..`

