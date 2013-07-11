spc2geo
=======

Converts state plane coordinates to WGS84 geodetic.

Requirements
------------

* Proj4 4.7.0-2 (dev package)
* Go 1.1

Install
-------

First, check the cgo parameters are correctly pointing to the Proj4 header. After that, build it:

    go build spc2geo.go

Usage
-----

    spc2geo --nad83 <zone> <Easting> <Northing>
    zone := 4-digit code based on NAD83(97) datum; defaults to 1201
    Easting := U.S. survey feet
    Northing := U.S. survey feet

Example
-------

    $ ./spc2geo --nad83 1201 1014500 1255455
    lon=-88.2251778219	lat=40.1139980020


