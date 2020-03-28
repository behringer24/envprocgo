# envproc
Easy environment variable preprocessor for configuration files

## Why
When building docker containers you usually rely on environment variables as a source for configuration setting and exposing single settings to the user of the ready made containers. Unfortunately not all softwares allow the substitution of environment variables in their config files.

Here envproc comes in handy. Ultra lightweight (apart from needing python 2.7 to run) and easy to use while setting up your containers during build or even on startup time.

## Installation

### Dependencies
envproc is written in python and tested with python 2.7 so you will nee a configured python interpreter

### Get the file
Download envproc from github or clone the entire repository. Only the `envproc` file is needed. You might need to add the extension `.py` on windows based systems. 

## Usage

### Getting help
Call `envproc -h` or `envproc --help` to get the standard commandline help

``` /bin/bash
usage: envproc [-h] [-c CHAR] [-f] [-v] [infile] [outfile]

Preprocess configuration files and fille with environment variables.

positional arguments:
  infile                the input file to preprocess. Or use pipe for stdin
  outfile               the output file to write the parsed result to. Or use
                        pipe for stdout

optional arguments:
  -h, --help            show this help message and exit
  -c CHAR, --char CHAR  character that is used as an variable indicator
  -f, --force           do not stop execution if environment variable is not
                        found
  -v, --version         show program's version number and exit
```

### Adding variables to config files
Before you can parse your configuration files you need to add markers that can be replaced by envproc.

```
[...]
variable_name = ${env:ENVNAME}
[...]
```

In this example the placeholder `${env:ENVNAME}` will be replaced with the value of the environment variable `ENVNAME`. If the variable is not set, envproc will stop execution and display an error.

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
envproc your_config_template.conf your_config.conf
```

This will read `your_config_template.conf` as an input file template and write the processed results to `your_config.conf`.

#### Using pipes for in and out
```
envproc < your_config_template.conf > your_config.conf
```

### Changing the prefix character
As default envproc uses `$` as the prefix character in the pattern it searches your configs for, like `${env:PATH}`. In some cases you might want to change it to another character. For this you can use the `--char` or `-c` option.

```
envproc -c% infile.conf outfile.conf
```

This changes the pattern it looks for to `%{env:PATH}`.