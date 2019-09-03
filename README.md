nuu
=======

[![Build Status](https://travis-ci.org/Songmu/num.svg?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/Songmu/num/badge.svg?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/num?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/num
[coveralls]: https://coveralls.io/r/Songmu/num?branch=master
[license]: https://github.com/Songmu/num/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/num

The `num` command displays the argument given in binary, octal, decimal or hexadecimal number in each base.

## Synopsis

```console
$ num 0xff
bin: 0b11111111
oct: 0o377
dec: 255
hex: 0xff

$ num 255
bin: 0b11111111
oct: 0o377
dec: 255
hex: 0xff

$ num 0b11111111
bin: 0b11111111
oct: 0o377
dec: 255
hex: 0xff
```

## Description

## Installation

```console
% go get github.com/Songmu/num/cmd/num
```

## Author

[Songmu](https://github.com/Songmu)
