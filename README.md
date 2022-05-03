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
$ sasb "あ"
binary:      [11100011, 10000001, 10000010]
octal:       [343, 201, 202]
decimal:     [227, 129, 130]
hexadecimal: [e3, 81, 82]
```
