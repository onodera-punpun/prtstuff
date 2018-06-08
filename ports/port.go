package ports

import (
	"path"
)

// A Port describes a port. A port is a directory containing the files needed
// for building a package.
type Port struct {
	// Location specifies the location of the port, this is used as the "primary
	// key" of a port type.
	Location Location

	// TODO: Add signature, .nostrip, et cetera.
	Footprint Footprint
	Md5sum    Md5sum
	Pkgfile   Pkgfile
}

// New return a Port with the Location field populated. Use the various `Parse*`
// functions to populate the other fields.
func New(location string) Port {
	var p Port
	p.Location = Location{path.Dir(path.Dir(location)), path.Base(path.Dir(
		location)), path.Base(location)}
	p.Footprint = Footprint{Port: &p}
	p.Md5sum = Md5sum{Port: &p}
	p.Pkgfile = Pkgfile{Port: &p}

	return p
}