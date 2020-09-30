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

## Configure DNS server

To change DNS server (resolver), you have to change OS specified resolver.

For Linux, please edit `/etc/resolv.conf`.

*   [CoreOS](https://coreos.com/rkt/docs/latest/networking/dns.html)
*   [Ubuntu](https://ubuntu.com/server/docs/network-configuration#heading--dns-client-configuration)
*   [Configure Google PUblic DNS for Linux](https://developers.google.com/speed/public-dns/docs/using#linux)

For Windows, please configure network adapter's TCP/IP settings.

*   [Configure Google Public DNS for Windows](https://developers.google.com/speed/public-dns/docs/using#windows)
