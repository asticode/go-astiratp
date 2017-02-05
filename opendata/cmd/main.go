package main

import (
	"flag"

	"github.com/asticode/go-astilog"
	"github.com/asticode/go-astiratp/opendata"
	"github.com/asticode/go-astitools/flag"
)

// Flags
var (
	id = flag.String("id", "", "the ID")
)

func main() {
	// Parse flags
	var s = astiflag.Subcommand()
	flag.Parse()

	// Init logger
	var l = astilog.New(astilog.FlagConfig())

	// Init client
	var c = astiopendata.New(astiopendata.FlagConfig())
	c.Logger = l

	// Switch on subcommand
	var err error
	switch s {
	case "geopoint":
		var g astiopendata.GeoPoint
		if g, err = c.GeoPoint(*id); err != nil {
			l.Fatal(err)
		}
		l.Infof("Geo point is %+v", g)
	case "lines":
		var ls []astiopendata.Line
		if ls, err = c.Lines(); err != nil {
			l.Fatal(err)
		}
		l.Infof("First line is %+v out of %d lines", ls[0], len(ls))
	case "stations":
		var ss []astiopendata.Station
		if ss, err = c.Stations(); err != nil {
			l.Fatal(err)
		}
		l.Infof("First station is %+v out of %d stations", ss[0], len(ss))
	case "version":
		var v astiopendata.Version
		if v, err = c.Version(); err != nil {
			l.Fatal(err)
		}
		l.Infof("Version is %s", v)
	}
}
