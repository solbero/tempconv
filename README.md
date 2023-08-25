<p align="center">
    <img src="https://raw.githubusercontent.com/solbero/tempconv/main/logo.png" alt="Logo" />
</p>
<p align="center">
    <em>Temperature conversion tool for the command line written in Go.</em>
</p>
<p align="center">
    <a href="https://github.com/solbero/tempconv/releases">
        <img src="https://img.shields.io/github/v/release/solbero/tempconv" alt="Release" />
    </a>
    <a href="https://github.com/solbero/tempconv/actions/workflows/release.yml">
        <img src="https://img.shields.io/github/actions/workflow/status/solbero/tempconv/release.yml" alt="Build" />
    </a>
    <a href="https://github.com/solbero/tempconv/actions/workflows/test.yml">
        <img src="https://img.shields.io/github/actions/workflow/status/solbero/tempconv/test.yml?label=tests" alt="Tests" />
    </a>
    <a href="https://codecov.io/gh/solbero/tempconv" >
        <img alt="Codecov" src="https://img.shields.io/codecov/c/github/solbero/tempconv" />
    </a>
    <a href="https://github.com/solbero/tempconv/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/solbero/tempconv" alt="License" />
    </a>
</p>

## Demo

<img src="https://raw.githubusercontent.com/solbero/tempconv/main/demo.gif" alt="Demo" style="zoom:90%;" />

## About

Tempconv supports conversion between the following temperature scales:
 - Kelvin
 - Celsius
 - Fahrenheit
 - Rankine
 - Delisle
 - Newton
 - Réaumur
 - Rømer

## Installation

### Binary

Download the latest release for your platform from the [releases page](https://github.com/solbero/tempconv/releases) and place it somewhere in your `PATH`.

### Go

If you have Go installed, you can install tempconv with the following command:


```sh
go install https://github.com/solbero/tempconv
```

## Usage

```sh
tempconv [-d -u <int> | -v | -h] temp from_scale to_scale
```

**Arguments**

* `temp`: Temperature to convert
* `from_scale`: Scale to convert temperature from
* `to_scale`: Scale to convert temperature to

If temperature is negative, it must be prefixed with '--' to avoid being interpreted as a flag.

**Options**

* `-d <int>`: Number of decimal places [default: 2, min: 0, max: 12]
* `-h`: Show help and exit
* `-u`: Include temperature unit
* `-v`: Show version and exit