# Top-Level Domain Parser

Tiny library to parse the subdomain, domain, and tld extension from a host string.
https://github.com/goware/tldomains
## Usage

```go
hostInfo := tldomains.Parse("mmmm.jello.co.uk")
// hostInfo.Subdomain = "mmmm"
// hostInfo.Domain = "jello"
// hostInfo.Suffix = "co.uk"
```

## Build

`make dist` will update the tldomains.dat file from https://publicsuffix.org/list/effective_tld_names.dat and rebuild the library.

The tld data file is automatically bundled in this library for distribution in tldomains.dat.gen.go.

## License

http://mozilla.org/MPL/2.0/
