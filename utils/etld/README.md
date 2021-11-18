# effective top-level domains (eTLD)

eTLD: effective top-level domains.

a simple library to parse the subdomain, domain, and eTLD from host.
## Usage

```go
hostInfo := tldomains.Parse("www.lixueduan.com")
// hostInfo.Subdomain = "www"
// hostInfo.Domain = "lixueduan"
// hostInfo.Suffix = "com"
```

## Build

`make dist` will update the tldomains.dat file from https://publicsuffix.org/list/effective_tld_names.dat and rebuild the library.

The tld data file is automatically bundled in this library for distribution in tldomains.dat.gen.go.

## License

http://mozilla.org/MPL/2.0/
