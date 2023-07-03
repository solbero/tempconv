# Tempconv

![Release](https://img.shields.io/github/v/release/solbero/tempconv)
![Build](https://img.shields.io/github/actions/workflow/status/solbero/tempconv/release.yml)
![Tests](https://img.shields.io/github/actions/workflow/status/solbero/tempconv/test.yml?label=tests)
![Codecov](https://img.shields.io/codecov/c/github/solbero/tempconv)
![License](https://img.shields.io/github/license/solbero/tempconv)


Temperature conversion tool for the command line written in Go.

## Description

The tool supports conversion between the following temperature scales:
 - Kelvin
 - Celsius
 - Fahrenheit
 - Rankine
 - Delisle
 - Newton
 - Réaumur
 - Rømer

## Installation

Download the latest release for your platform from the [releases page](https://github.com/solbero/tempconv/releases) and place it somewhere in your `PATH`.

## Usage

`tempconv [-u | -d <int>] <temp> <from scale> <to scale>` — Converts a temperature from one scale to another.

`tempconv -h` — Shows help.

`tempconv -v` — Shows version.