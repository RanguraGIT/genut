# Genut

Genut is a tool designed to simplify development, security, and operations tasks for the DG project.

## Overview

This tool provides functionalities for generating automatically generated mocks using `mockgen` and wrapping files as needed. It is built using the Go programming language.

## Installation

To install `genut`, you need to have Go installed on your machine. You can then use the following command:

```bash
go install github.com/RanguraGIT/genut@latest

```

# Usage

The tool comes with a command-line interface to execute its functionalities.

## Generate Command

The generate command allows you to trigger various code generation actions, such as generating mocks and wrappers.

```bash
genut [command/flags]

```
Available `genut` *flags*:
- --config or -c       
*Generate config file to root project*
- --version or -v      
*Print the version number of Genut*

Available `genut` *command*:
- create  *"project name"*    
*Generate new project*
- install *"service name"*  
*Add new service to project*
- mocks  
*Generate mocks from available interface*
- pre-commit  
*Installing pre-commit configuration*


Example:
```bash
# For generating genut config to root project
genut --config
# For Generate mocks from available interface
genut mocks
# For Installing pre-commit configuration
genut pre-commit
```

# Contributing

Contributions are welcome! If you have suggestions, bug reports, or feature requests, please open an issue or a pull request in this repository.

# License

This project is licensed under the MIT License.
