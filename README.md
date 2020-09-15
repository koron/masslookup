# koron/masslookup

[![GoDoc](https://godoc.org/github.com/koron/masslookup?status.svg)](https://godoc.org/github.com/koron/masslookup)
[![Actions/Go](https://github.com/koron/masslookup/workflows/Go/badge.svg)](https://github.com/koron/masslookup/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/masslookup)](https://goreportcard.com/report/github.com/koron/masslookup)

mass DNS lookup

```console
$ masslookup < ip.txt > name.txt
```

input format:

```
1.2.3.4
192.168.0.1
111.112.113.114
```

output format:

* Found

        {IP} \t FOUND \t {WHITESPACE_SEPARATED_NAMES}

* Error:

        {IP} \t ERROR \t {ERROR_MESSAGE}

* No entries:

        {IP} \t NOENTRY

## Usage

```
  -concurrency int
        concurrency of DNS query (default 500)
  -hideerror
        hide query ERROR
  -hidenoentry
        hide NOENTRY
```
