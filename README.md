# Go Packet Decoder

A Go program that decodes a fixed-size packet according to a specific decoding format and returns the decoded data in a struct.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This Go program is designed to decode a fixed-size packet based on a specific decoding format. It unpacks the data in the packet and returns the decoded information in the form of a struct. The decoding map is as follows:

1. First 2 bytes represent a short data type.
2. The next 12 bytes represent 12 characters.
3. The next 1 byte represents a single byte.
4. The next 8 bytes represent 8 characters.
5. The next 2 bytes represent a short data type.
6. The next 15 bytes represent 15 characters.
7. The next 4 bytes represent a long data type.

## Installation

To use this package, you can install it via `go get`:

```sh
go get github.com/BrijDev/essentially-ai
