# Genut

Genut is a tool designed to simplify development, security, and operations tasks for the DG project.

## Overview

This tool provides functionalities for generating automatically generated mocks using `mockgen` and wrapping files as needed. It is built using the Go programming language and the [Cobra](https://github.com/spf13/cobra) library for creating command-line interfaces.

## Installation

To install `genut`, you need to have Go installed on your machine. You can then use the following command:

```bash
go get -u github.com/RanguraGIT/genut
```

# Usage

The tool comes with a command-line interface to execute its functionalities.

# Generate Command

The generate command allows you to trigger various code generation actions, such as generating mocks and wrappers.

```bash
genut generate [flags]
```

Available flags:

- `--mocking`: Generate mockgen mocks.
- `--wrapper`: Generate wrappers.

Example:

```bash
genut generate --mocking
```
This will generate mockgen mocks based on the configurations interface provided.

# Contributing
Contributions are welcome! If you have suggestions, bug reports, or feature requests, please open an issue or a pull request in this repository.

# License
This project is licensed under the MIT License.
