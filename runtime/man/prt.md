# prt 8 "2017-01-02" prt "General Commands Manual"

## NAME

prt - CRUX port utility written in Go, aiming to replace prt-get, ports, and some pkgutils (on my machine)


## SYNOPSIS

prt command [arguments]


## DESCRIPTION

prt is like `prt-get(8)` a port/package management utility which provides additional functionality to the
CRUX pkgutils. It works with the local ports tree and is therefore fully compatible with `ports(8)`, `pkgmk(8)`,
`pkgadd(8)` and of course `prt-get(8)`. It offers the following features:

* listing dependencies of ports recursively, with an optional flag to print using tree view
* listing outdated package, by comparing port versions with the installed version
* easily printing port information such as the maintainer, version, release, et cetera
* install ports and their dependencies with a single command
* list ports, with optional flags to also only list installed ports, print with repo information or to print
  with additional version information
* print the location of a port
* searching through files ports provide, with an optional flag to only search through installed ports
* pull in ports using git(1)
* update outdated packages
* uninstall installed packages

like `prt-get(8)`, prt is basically a wrapper around `pkgmk(8)`/`pkgadd(8)` and provides some nice functionality such as
listing and installing dependencies, getting the location of a port, aliasing ports (for example `core/openssl`
to `6c37-dropin/libressl`), and ordering ports with the same name depending on how "important" the repo is the port resides in.

There are a few differences though, for example, unlike `prt-get(8)` you need to  be in the port's directory for most
commands to work, like how `pkgmk(8)` works. This has a few advantages, for example you can quickly download a port
anywhere on the filesystem, and install it and its dependencies using `prt install`. Because `prt-get depinst` needs
a port name, you can *only* install ports that are located in a predefined `prtdir`.

Another difference with `prt-get(8)` is that prt does not use a cache file, while still being nearly as fast or faster
in some cases.

Aliasing is also handeled a bit different. `prt-get(8)` aliases ports based on name, but prt on name and repo.
This makes it possible to alias `foo/bar` to `baz/bar`.


## COMMANDS

The prt syntax is inspired by `prt-get(8)`, `git(8)` and `go(8)`, and thus uses so called commands which always have to be the first
non-option argument passed. Each command is documented in its own manpage. For example, the install command is documented in
`prt-install(8)`. The commands are:

`depends`   list dependencies recursively,

`diff`      list outdated packages

`info`      print port information

`install`   build and install ports and their dependencies

`list`      list porst and packages

`loc`       print port locations

`prov`      search ports for files

`pull`      pull in ports

`sysup`     update outdated packages

`uninstall` uninstall packages

`help`      print help and exit


## CONFIGURATION

See man `prt.toml(5)`


## AUTHORS

Camille Scholtz


## SEE ALSO

`cdp(1)`, `prt.toml(5)`, `prt-depends(8)`, `prt-diff(8)`, `prt-info(8)`, `prt-install(8)`, `prt-list(8)`,
`prt-loc(8)`, `prt-prov(8)`, `prt-pull(8)`, `prt-sysup(8)`, `prt-uninstall(8)`, `prt-get(8)`,
`pkgmk(8)`, `pkgrm(8)`, `pkgadd(8)`, `ports(8)`, `pkginfo(8)`, `prt-utils(1)`
