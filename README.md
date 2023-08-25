<p align="center"><img src="https://github.com/solbero/tempconv/blob/main/logo.png" alt="Logo" /></p>
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
    <a href="https://app.codecov.io/gh/solbero/tempconv">
        <img src="https://img.shields.io/github/actions/workflow/status/solbero/tempconv/test.yml?label=tests" alt="Tests" />
    </a>
    <a href="https://github.com/solbero/tempconv/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/solbero/tempconv" alt="License" />
    </a>
</p>

## Demo

<img src="https://github.com/solbero/tempconv/blob/main/demo.gif" alt="Demo" style="zoom:90%;" />

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