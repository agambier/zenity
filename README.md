# Zenity dialogs for Golang, Windows and macOS

This repo is a fork of [go-zenity](https://github.com/agambier/zenity/v23) which implements some features of zenity (only those portable on every platform). 
This fork implements the missing features even i fthey can not be fully supported by each platform.

[![Go Reference](https://pkg.go.dev/badge/image)](https://pkg.go.dev/github.com/agambier/zenity/v23)
[![Go Report](https://goreportcard.com/badge/github.com/agambier/zenity/v23)](https://goreportcard.com/report/github.com/agambier/zenity/v23)
[![Go Coverage](https://github.com/agambier/zenity/v23/wiki/coverage.svg)](https://raw.githack.com/wiki/agambier/zenity/v23/coverage.html)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

This repo includes:
- a cross-platform [Go](https://go.dev/) package providing
[Zenity](https://help.gnome.org/users/zenity/stable/)-like dialogs
(simple dialogs that interact graphically with the user)
- a *“port”* of the `zenity` command to both Windows and macOS based on that library.

Implemented dialogs:
* [message](https://github.com/agambier/zenity/v23/wiki/Message-dialog) (error, info, question, warning)
* [text entry](https://github.com/agambier/zenity/v23/wiki/Text-entry-dialog)
* [list](https://github.com/agambier/zenity/v23/wiki/List-dialog) (simple)
* [password](https://github.com/agambier/zenity/v23/wiki/Password-dialog)
* [file selection](https://github.com/agambier/zenity/v23/wiki/File-selection-dialog)
* [color selection](https://github.com/agambier/zenity/v23/wiki/Color-selection-dialog)
* [calendar](https://github.com/agambier/zenity/v23/wiki/Calendar-dialog)
* [progress](https://github.com/agambier/zenity/v23/wiki/Progress-dialog)
* [notification](https://github.com/agambier/zenity/v23/wiki/Notification)

Behavior on Windows, macOS and other Unixes might differ slightly.
Some of that is intended (reflecting platform differences),
other bits are unfortunate limitations.

## Installing

The Go package:

    go get github.com/agambier/zenity/v23@latest

The `zenity` command on macOS/WSL using [Homebrew](https://brew.sh/) 🍺:

    brew install ncruces/tap/zenity

The `zenity` command on Windows using [Scoop](https://scoop.sh/) 🍨:

    scoop install https://ncruces.github.io/scoop/zenity.json

The `zenity` command on macOS/Windows, if you have [Go](https://go.dev/):

    go install github.com/agambier/zenity/v23/cmd/zenity@latest

Or download the [latest release](https://github.com/agambier/zenity/v23/releases/latest).

## Using

For the Go package, consult the [documentation](https://pkg.go.dev/github.com/agambier/zenity/v23#section-documentation)
and [examples](https://pkg.go.dev/github.com/agambier/zenity/v23#pkg-examples).

The `zenity` command does its best to be compatible with the GNOME version.\
Consult the [documentation](https://help.gnome.org/users/zenity/stable/)
and [man page](https://linux.die.net/man/1/zenity) of that command.

## Why?

#### Benefits of the Go package:

* no `cgo` (see [benefits](https://dave.cheney.net/2016/01/18/cgo-is-not-go), mostly cross-compilation)
* no main loop (or any other threading or initialization requirements)
* cancelation through [`context`](https://golang.org/pkg/context/)
* on Windows:
  * no additional dependencies
    * Explorer shell not required
    * works in Server Core
  * Unicode support
  * High DPI (no manifest required)
  * Visual Styles (no manifest required)
  * WSL/Cygwin/MSYS2 [support](https://github.com/agambier/zenity/v23/wiki/Zenity-for-WSL,-Cygwin,-MSYS2)
* on macOS:
  * only dependency is `osascript`
* on other Unixes:
  * wraps either one of `zenity`, `matedialog`, [`qarma`](https://github.com/luebking/qarma)

## Credits

I'd like to thank all [contributors](https://github.com/agambier/zenity/v23/graphs/contributors),
but [@gen2brain](https://github.com/gen2brain) in particular
for [`dlgs`](https://github.com/gen2brain/dlgs),
which was instrumental to the Windows port of `zenity`.
