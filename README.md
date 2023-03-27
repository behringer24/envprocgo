[![Build + Test](https://github.com/behringer24/envprocgo/actions/workflows/go.yml/badge.svg)](https://github.com/behringer24/envprocgo/actions/workflows/go.yml)
[![Release build](https://github.com/behringer24/envprocgo/actions/workflows/release.yml/badge.svg)](https://github.com/behringer24/envprocgo/actions/workflows/release.yml)
# envproc
Easy environment variable preprocessor for configuration files

## Why
When building docker containers you usually rely on environment variables as a source for configuration setting and exposing single settings to the user of the ready made containers. Unfortunately not all softwares allow the substitution of environment variables in their config files.

Here envproc comes in handy. Ultra lightweight and easy to use while setting up your containers during build or even on startup time.

## Installation

### Dependencies
envproc is written in Go (Golang) and compiles to a single binary file.

### Install the binary
Download the binary from envproc from github or clone the entire repository. Only the `envproc` file is needed. 

``` /bin/bash
> wget https://github.com/behringer24/envprocgo/releases/download/v1.0.3/envproc-v1.0.2-linux-amd64.tar.gz
> tar -xvzf envproc-v1.0.3-linux-amd64.tar.gz
```

or

``` /bin/bash
> wget -qO- https://github.com/behringer24/envprocgo/releases/download/v1.0.3/envproc-v1.0.2-linux-amd64.tar.gz | tar xvz
```

Above is an example for version v1.0.3. Check for newer versions here https://github.com/behringer24/envprocgo/releases

### Install from source
Make sure you have Go installed. Find more information here https://go.dev/dl/

Checkout the sources of the main branch or a specific (latest) release tag.

``` /bin/bash
> git clone git@github.com:behringer24/envprocgo.git
```

to install

``` /bin/bash
> go install
```

## Usage

### Getting help
Call `envproc -h` or `envproc --help` to get the standard commandline help

``` /bin/bash
> envproc -h
envproc
Config file preprocessor, inject environment variables into static config files

Usage: envproc [-f] [-h] [-v] [-c] infile [outfile]

Flags:
-h, --help               Show this help text
-v, --version            Show version information
-f, --force              Show version information

Options:
-c, --char               Another description (Default: $)

Positional arguments:
infile                   File to read from
outfile                  File to write to
```

### Getting the current version
To check the current version of the executable use

``` /bin/bash
> envproc -v
envproc version v0.0.1
```

### Adding variables to config files
Before you can parse your configuration files you need to add markers that can be replaced by envproc.

```
[...]
variable_name = ${env:ENVNAME}
[...]
```

In this example the placeholder `${env:ENVNAME}` will be replaced with the value of the environment variable `ENVNAME`. If the variable is not set, envproc will stop execution and display an error. You can override the stopping with the `-f|--force` parameter.

So if you would like to put the value of the common `PATH` variable into your configuration file you would write:
```
[...]
app_path_var = ${env:PATH}
[...]
```

### Parsing your config files
envproc can be used in different ways, with the simple use of input- and output file or with streams/pipes stdin and stdout.

#### Simple filenames for in and out
```
>envproc your_config_template.conf your_final_config.conf
```

This will read `your_config_template.conf` as an input file template and write the processed results to `your_final_config.conf`.

#### Using pipe for output
```
> envproc your_config_template.conf > your_config.conf
```

### Changing the prefix character
As default envproc uses `$` as the prefix character in the pattern it searches your configs for, like `${env:PATH}`. In some cases you might want to change it to another character. For this you can use the `--char` or `-c` option.

```
> envproc -c % infile.conf outfile.conf
```

The pattern envproc now looks for is `%{env:PATH}`.

## License

envproc and anvprocgo are released under the GNU GENERAL PUBLIC LICENSE Version 3. See [LICENSE](https://github.com/behringer24/envprocgo/blob/main/LICENSE)