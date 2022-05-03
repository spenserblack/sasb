# String as Bytes

[![CI](https://github.com/spenserblack/sasb/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/sasb/actions/workflows/ci.yml)

A utility executable for viewing strings as bytes.

I wrote this because I found myself frequently checking the byte values of characters.

## Install

```shell
go install github.com/spenserblack/sasb
```

## Usage

```console
$ sasb "@"
binary:      [01000000]
octal:       [100]
decimal:     [64]
hexadecimal: [40]
```
