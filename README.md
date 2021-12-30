![Build Status](https://github.com/dh1tw/hid/workflows/Cross%20Platform%20Tests/badge.svg?branch=master)
[![GoDoc](https://godoc.org/github.com/dh1tw/hid?status.svg)](https://godoc.org/github.com/dh1tw/hid)

This package has been forked from [karalabe/hid](https://github.com/karalabe/hid) where the development
has unfortunately stopped in 2019. The package came with vendored versions of 
[`libusb`](https://github.com/libusb/libusb) and [`hidapi`](https://github.com/libusb/hidapi) which
haven't been updated in more than 5 years. I updated and simplified the cgo build commands,
so that the package builds again on Linux, Windows and MacOS. As part of the simplification, I removed
the vendored version of `libusb`. It's only needed for Linux, and all major distros provide more
recent versions of `libusb`. You just have to make sure that `libusb-1.0-0-dev` is installed on your 
Linux system. On MacOS and Windows `libusb` is not needed. Both OS have direct `hid` APIs which are
accessed by `hidapi`.

The vendored version of `hidapi` is `hidapi-0.11.0-7-gaf6c601`.

# Gopher Interface Devices (USB HID)

Supported platforms at the moment are Linux, macOS and Windows (exclude constraints are also specified
for Android and iOS to allow smoother vendoring into cross platform projects).

# Linux support

Make sure you have `libusb-1.0-0-dev` installed on your system.

```
$ sudo apt install libusb-1.0-0-dev
```

## Cross-compiling

Using `go get` the embedded C library is compiled into the binary format of your host OS. Cross compiling to a different platform or architecture entails disabling CGO by default in Go, causing device enumeration `hid.Enumerate()` to yield no results.

To cross compile a functional version of this library, you'll need to enable CGO during cross compilation via `CGO_ENABLED=1` and you'll need to install and set a cross compilation enabled C toolkit via `CC=your-cross-gcc`.

## Acknowledgements

Although the `hid` package is an implementation from scratch, it was heavily inspired by the existing
[`go.hid`](https://github.com/GeertJohan/go.hid) library, which seems abandoned since 2015; is incompatible
with Go 1.6+; and has various external dependencies. Given its inspirational roots, I thought it important
to give credit to the author of said package too.

Wide character support in the `hid` package is done via the [`gowchar`](https://github.com/orofarne/gowchar)
library, unmaintained since 2013; non buildable with a modern Go release and failing `go vet` checks. As
such, `gowchar` was also vendored in inline (copyright headers and origins preserved).

## License

The components of `hid` are licensed as such:

 * `hidapi` is released under the [3-clause BSD](https://github.com/signal11/hidapi/blob/master/LICENSE-bsd.txt) license.
 * `libusb` is released under the [GNU LGPL 2.1](https://github.com/libusb/libusb/blob/master/COPYING)license.
 * `go.hid` is released under the [2-clause BSD](https://github.com/GeertJohan/go.hid/blob/master/LICENSE) license.
 * `gowchar` is released under the [3-clause BSD](https://github.com/orofarne/gowchar/blob/master/LICENSE) license.

Given the above, `hid` is licensed under GNU LGPL 2.1 or later on Linux and 3-clause BSD on other platforms.
