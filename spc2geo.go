package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

/*
#cgo LDFLAGS: -L/usr/lib -lproj
#include <proj_api.h>
*/
import "C"

var zone = flag.String("nad83", "1201", "NAD83(97) coordinate zone")
var ErrHelp = errors.New("flag: help requested")

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "spc2geo --nad83=<zone> <Easting> <Northing>\n")
	flag.PrintDefaults()
}

func main() {

	zone2params := make(map[string]string)

	zone2params["0101"] = "+proj=tmerc +lat_0=30.5 +lon_0=-85.83333333333333 +k=0.99996 +x_0=200000 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0102"] = "+proj=tmerc +lat_0=30 +lon_0=-87.5 +k=0.999933333 +x_0=600000 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0201"] = "+proj=tmerc +lat_0=31 +lon_0=-110.1666666666667 +k=0.9999 +x_0=213360 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["0202"] = "+proj=tmerc +lat_0=31 +lon_0=-111.9166666666667 +k=0.9999 +x_0=213360 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["0203"] = "+proj=tmerc +lat_0=31 +lon_0=-113.75 +k=0.999933333 +x_0=213360 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["0301"] = "+proj=lcc +lat_1=36.23333333333333 +lat_2=34.93333333333333 +lat_0=34.33333333333334 +lon_0=-92 +x_0=399999.99998984 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0302"] = "+proj=lcc +lat_1=34.76666666666667 +lat_2=33.3 +lat_0=32.66666666666666 +lon_0=-92 +x_0=399999.99998984 +y_0=399999.99998984 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0401"] = "+proj=lcc +lat_1=41.66666666666666 +lat_2=40 +lat_0=39.33333333333334 +lon_0=-122 +x_0=2000000.0001016 +y_0=500000.0001016001 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0402"] = "+proj=lcc +lat_1=39.83333333333334 +lat_2=38.33333333333334 +lat_0=37.66666666666666 +lon_0=-122 +x_0=2000000.0001016 +y_0=500000.0001016001 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0403"] = "+proj=lcc +lat_1=38.43333333333333 +lat_2=37.06666666666667 +lat_0=36.5 +lon_0=-120.5 +x_0=2000000.0001016 +y_0=500000.0001016001 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0404"] = "+proj=lcc +lat_1=37.25 +lat_2=36 +lat_0=35.33333333333334 +lon_0=-119 +x_0=2000000.0001016 +y_0=500000.0001016001 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0405"] = "+proj=lcc +lat_1=35.46666666666667 +lat_2=34.03333333333333 +lat_0=33.5 +lon_0=-118 +x_0=2000000.0001016 +y_0=500000.0001016001 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0406"] = "+proj=lcc +lat_1=33.88333333333333 +lat_2=32.78333333333333 +lat_0=32.16666666666666 +lon_0=-116.25 +x_0=2000000.0001016 +y_0=500000.0001016001 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0501"] = "+proj=lcc +lat_1=40.78333333333333 +lat_2=39.71666666666667 +lat_0=39.33333333333334 +lon_0=-105.5 +x_0=914401.8288036576 +y_0=304800.6096012192 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0502"] = "+proj=lcc +lat_1=39.75 +lat_2=38.45 +lat_0=37.83333333333334 +lon_0=-105.5 +x_0=914401.8288036576 +y_0=304800.6096012192 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0503"] = "+proj=lcc +lat_1=38.43333333333333 +lat_2=37.23333333333333 +lat_0=36.66666666666666 +lon_0=-105.5 +x_0=914401.8288036576 +y_0=304800.6096012192 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0600"] = "+proj=lcc +lat_1=41.86666666666667 +lat_2=41.2 +lat_0=40.83333333333334 +lon_0=-72.75 +x_0=304800.6096012192 +y_0=152400.3048006096 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0700"] = "+proj=tmerc +lat_0=38 +lon_0=-75.41666666666667 +k=0.999995 +x_0=200000.0001016002 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0901"] = "+proj=tmerc +lat_0=24.33333333333333 +lon_0=-81 +k=0.999941177 +x_0=200000.0001016002 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0902"] = "+proj=tmerc +lat_0=24.33333333333333 +lon_0=-82 +k=0.999941177 +x_0=200000.0001016002 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["0903"] = "+proj=lcc +lat_1=30.75 +lat_2=29.58333333333333 +lat_0=29 +lon_0=-84.5 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1001"] = "+proj=tmerc +lat_0=30 +lon_0=-82.16666666666667 +k=0.9999 +x_0=200000.0001016002 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1002"] = "+proj=tmerc +lat_0=30 +lon_0=-84.16666666666667 +k=0.9999 +x_0=699999.9998983998 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1101"] = "+proj=tmerc +lat_0=41.66666666666666 +lon_0=-112.1666666666667 +k=0.9999473679999999 +x_0=200000.0001016002 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1102"] = "+proj=tmerc +lat_0=41.66666666666666 +lon_0=-114 +k=0.9999473679999999 +x_0=500000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1103"] = "+proj=tmerc +lat_0=41.66666666666666 +lon_0=-115.75 +k=0.999933333 +x_0=800000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1201"] = "+proj=tmerc +lat_0=36.66666666666666 +lon_0=-88.33333333333333 +k=0.9999749999999999 +x_0=300000.0000000001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1202"] = "+proj=tmerc +lat_0=36.66666666666666 +lon_0=-90.16666666666667 +k=0.999941177 +x_0=699999.9999898402 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1301"] = "+proj=tmerc +lat_0=37.5 +lon_0=-85.66666666666667 +k=0.999966667 +x_0=99999.99989839978 +y_0=249999.9998983998 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1302"] = "+proj=tmerc +lat_0=37.5 +lon_0=-87.08333333333333 +k=0.999966667 +x_0=900000 +y_0=249999.9998983998 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1401"] = "+proj=lcc +lat_1=43.26666666666667 +lat_2=42.06666666666667 +lat_0=41.5 +lon_0=-93.5 +x_0=1500000 +y_0=999999.9999898402 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1402"] = "+proj=lcc +lat_1=41.78333333333333 +lat_2=40.61666666666667 +lat_0=40 +lon_0=-93.5 +x_0=500000.00001016 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1501"] = "+proj=lcc +lat_1=39.78333333333333 +lat_2=38.71666666666667 +lat_0=38.33333333333334 +lon_0=-98 +x_0=399999.99998984 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1502"] = "+proj=lcc +lat_1=38.56666666666667 +lat_2=37.26666666666667 +lat_0=36.66666666666666 +lon_0=-98.5 +x_0=399999.99998984 +y_0=399999.99998984 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1601"] = "+proj=lcc +lat_1=37.96666666666667 +lat_2=38.96666666666667 +lat_0=37.5 +lon_0=-84.25 +x_0=500000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1602"] = "+proj=lcc +lat_1=37.93333333333333 +lat_2=36.73333333333333 +lat_0=36.33333333333334 +lon_0=-85.75 +x_0=500000.0001016001 +y_0=500000.0001016001 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1701"] = "+proj=lcc +lat_1=32.66666666666666 +lat_2=31.16666666666667 +lat_0=30.5 +lon_0=-92.5 +x_0=999999.9999898402 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1702"] = "+proj=lcc +lat_1=30.7 +lat_2=29.3 +lat_0=28.5 +lon_0=-91.33333333333333 +x_0=999999.9999898402 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1801"] = "+proj=tmerc +lat_0=43.66666666666666 +lon_0=-68.5 +k=0.9999 +x_0=300000.0000000001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1802"] = "+proj=tmerc +lat_0=42.83333333333334 +lon_0=-70.16666666666667 +k=0.999966667 +x_0=900000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["1900"] = "+proj=lcc +lat_1=39.45 +lat_2=38.3 +lat_0=37.66666666666666 +lon_0=-77 +x_0=399999.9998983998 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2001"] = "+proj=lcc +lat_1=42.68333333333333 +lat_2=41.71666666666667 +lat_0=41 +lon_0=-71.5 +x_0=200000.0001016002 +y_0=750000 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2002"] = "+proj=lcc +lat_1=41.48333333333333 +lat_2=41.28333333333333 +lat_0=41 +lon_0=-70.5 +x_0=500000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2111"] = "+proj=lcc +lat_1=47.08333333333334 +lat_2=45.48333333333333 +lat_0=44.78333333333333 +lon_0=-87 +x_0=7999999.999968001 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["2112"] = "+proj=lcc +lat_1=45.7 +lat_2=44.18333333333333 +lat_0=43.31666666666667 +lon_0=-84.36666666666666 +x_0=5999999.999976001 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["2113"] = "+proj=lcc +lat_1=43.66666666666666 +lat_2=42.1 +lat_0=41.5 +lon_0=-84.36666666666666 +x_0=3999999.999984 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["2201"] = "+proj=lcc +lat_1=48.63333333333333 +lat_2=47.03333333333333 +lat_0=46.5 +lon_0=-93.09999999999999 +x_0=800000.0000101599 +y_0=99999.99998983997 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2202"] = "+proj=lcc +lat_1=47.05 +lat_2=45.61666666666667 +lat_0=45 +lon_0=-94.25 +x_0=800000.0000101599 +y_0=99999.99998983997 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2203"] = "+proj=lcc +lat_1=45.21666666666667 +lat_2=43.78333333333333 +lat_0=43 +lon_0=-94 +x_0=800000.0000101599 +y_0=99999.99998983997 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2301"] = "+proj=tmerc +lat_0=29.5 +lon_0=-88.83333333333333 +k=0.99995 +x_0=300000.0000000001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2302"] = "+proj=tmerc +lat_0=29.5 +lon_0=-90.33333333333333 +k=0.99995 +x_0=699999.9998983998 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2401"] = "+proj=tmerc +lat_0=35.83333333333334 +lon_0=-90.5 +k=0.9999333333333333 +x_0=250000 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2402"] = "+proj=tmerc +lat_0=35.83333333333334 +lon_0=-92.5 +k=0.9999333333333333 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2403"] = "+proj=tmerc +lat_0=36.16666666666666 +lon_0=-94.5 +k=0.9999411764705882 +x_0=850000 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2500"] = "+proj=lcc +lat_1=49 +lat_2=45 +lat_0=44.25 +lon_0=-109.5 +x_0=599999.9999976 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["2600"] = "+proj=lcc +lat_1=43 +lat_2=40 +lat_0=39.83333333333334 +lon_0=-100 +x_0=500000.00001016 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2701"] = "+proj=tmerc +lat_0=34.75 +lon_0=-115.5833333333333 +k=0.9999 +x_0=200000.00001016 +y_0=8000000.000010163 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2702"] = "+proj=tmerc +lat_0=34.75 +lon_0=-116.6666666666667 +k=0.9999 +x_0=500000.00001016 +y_0=6000000 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2703"] = "+proj=tmerc +lat_0=34.75 +lon_0=-118.5833333333333 +k=0.9999 +x_0=800000.0000101599 +y_0=3999999.99998984 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2800"] = "+proj=tmerc +lat_0=42.5 +lon_0=-71.66666666666667 +k=0.999966667 +x_0=300000.0000000001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["2900"] = "+proj=tmerc +lat_0=38.83333333333334 +lon_0=-74.5 +k=0.9999 +x_0=150000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3001"] = "+proj=tmerc +lat_0=31 +lon_0=-104.3333333333333 +k=0.999909091 +x_0=165000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3002"] = "+proj=tmerc +lat_0=31 +lon_0=-106.25 +k=0.9999 +x_0=500000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3003"] = "+proj=tmerc +lat_0=31 +lon_0=-107.8333333333333 +k=0.999916667 +x_0=830000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3101"] = "+proj=tmerc +lat_0=38.83333333333334 +lon_0=-74.5 +k=0.9999 +x_0=150000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3102"] = "+proj=tmerc +lat_0=40 +lon_0=-76.58333333333333 +k=0.9999375 +x_0=249999.9998983998 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3103"] = "+proj=tmerc +lat_0=40 +lon_0=-78.58333333333333 +k=0.9999375 +x_0=350000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3104"] = "+proj=lcc +lat_1=41.03333333333333 +lat_2=40.66666666666666 +lat_0=40.16666666666666 +lon_0=-74 +x_0=300000.0000000001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3200"] = "+proj=lcc +lat_1=36.16666666666666 +lat_2=34.33333333333334 +lat_0=33.75 +lon_0=-79 +x_0=609601.2192024384 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3301"] = "+proj=lcc +lat_1=48.73333333333333 +lat_2=47.43333333333333 +lat_0=47 +lon_0=-100.5 +x_0=599999.9999976 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["3302"] = "+proj=lcc +lat_1=47.48333333333333 +lat_2=46.18333333333333 +lat_0=45.66666666666666 +lon_0=-100.5 +x_0=599999.9999976 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["3401"] = "+proj=lcc +lat_1=41.7 +lat_2=40.43333333333333 +lat_0=39.66666666666666 +lon_0=-82.5 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3402"] = "+proj=lcc +lat_1=40.03333333333333 +lat_2=38.73333333333333 +lat_0=38 +lon_0=-82.5 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3501"] = "+proj=lcc +lat_1=36.76666666666667 +lat_2=35.56666666666667 +lat_0=35 +lon_0=-98 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3502"] = "+proj=lcc +lat_1=35.23333333333333 +lat_2=33.93333333333333 +lat_0=33.33333333333334 +lon_0=-98 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3601"] = "+proj=lcc +lat_1=46 +lat_2=44.33333333333334 +lat_0=43.66666666666666 +lon_0=-120.5 +x_0=2500000.0001424 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["3602"] = "+proj=lcc +lat_1=44 +lat_2=42.33333333333334 +lat_0=41.66666666666666 +lon_0=-120.5 +x_0=1500000.0001464 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["3701"] = "+proj=lcc +lat_1=41.95 +lat_2=40.88333333333333 +lat_0=40.16666666666666 +lon_0=-77.75 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3702"] = "+proj=lcc +lat_1=40.96666666666667 +lat_2=39.93333333333333 +lat_0=39.33333333333334 +lon_0=-77.75 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3800"] = "+proj=tmerc +lat_0=41.08333333333334 +lon_0=-71.5 +k=0.99999375 +x_0=99999.99998983997 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["3900"] = "+proj=lcc +lat_1=34.83333333333334 +lat_2=32.5 +lat_0=31.83333333333333 +lon_0=-81 +x_0=609600 +y_0=0 +ellps=GRS80 +to_meter=0.3048 +no_defs"
	zone2params["4001"] = "+proj=lcc +lat_1=45.68333333333333 +lat_2=44.41666666666666 +lat_0=43.83333333333334 +lon_0=-100 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4002"] = "+proj=lcc +lat_1=44.4 +lat_2=42.83333333333334 +lat_0=42.33333333333334 +lon_0=-100.3333333333333 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4100"] = "+proj=lcc +lat_1=36.41666666666666 +lat_2=35.25 +lat_0=34.33333333333334 +lon_0=-86 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4201"] = "+proj=lcc +lat_1=36.18333333333333 +lat_2=34.65 +lat_0=34 +lon_0=-101.5 +x_0=200000.0001016002 +y_0=999999.9998983998 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4202"] = "+proj=lcc +lat_1=33.96666666666667 +lat_2=32.13333333333333 +lat_0=31.66666666666667 +lon_0=-98.5 +x_0=600000 +y_0=2000000.0001016 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4203"] = "+proj=lcc +lat_1=31.88333333333333 +lat_2=30.11666666666667 +lat_0=29.66666666666667 +lon_0=-100.3333333333333 +x_0=699999.9998983998 +y_0=3000000 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4204"] = "+proj=lcc +lat_1=30.28333333333333 +lat_2=28.38333333333333 +lat_0=27.83333333333333 +lon_0=-99 +x_0=600000 +y_0=3999999.9998984 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4205"] = "+proj=lcc +lat_1=27.83333333333333 +lat_2=26.16666666666667 +lat_0=25.66666666666667 +lon_0=-98.5 +x_0=300000.0000000001 +y_0=5000000.0001016 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4301"] = "+proj=lcc +lat_1=41.78333333333333 +lat_2=40.71666666666667 +lat_0=40.33333333333334 +lon_0=-111.5 +x_0=500000.00001016 +y_0=999999.9999898402 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4302"] = "+proj=lcc +lat_1=40.65 +lat_2=39.01666666666667 +lat_0=38.33333333333334 +lon_0=-111.5 +x_0=500000.00001016 +y_0=2000000.00001016 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4303"] = "+proj=lcc +lat_1=38.35 +lat_2=37.21666666666667 +lat_0=36.66666666666666 +lon_0=-111.5 +x_0=500000.00001016 +y_0=3000000 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4400"] = "+proj=tmerc +lat_0=42.5 +lon_0=-72.5 +k=0.9999642857142857 +x_0=500000 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4501"] = "+proj=lcc +lat_1=39.2 +lat_2=38.03333333333333 +lat_0=37.66666666666666 +lon_0=-78.5 +x_0=3500000.0001016 +y_0=2000000.0001016 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4502"] = "+proj=lcc +lat_1=37.96666666666667 +lat_2=36.76666666666667 +lat_0=36.33333333333334 +lon_0=-78.5 +x_0=3500000.0001016 +y_0=999999.9998983998 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4601"] = "+proj=lcc +lat_1=48.73333333333333 +lat_2=47.5 +lat_0=47 +lon_0=-120.8333333333333 +x_0=500000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4602"] = "+proj=lcc +lat_1=47.33333333333334 +lat_2=45.83333333333334 +lat_0=45.33333333333334 +lon_0=-120.5 +x_0=500000.0001016001 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4701"] = "+proj=lcc +lat_1=40.25 +lat_2=39 +lat_0=38.5 +lon_0=-79.5 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4702"] = "+proj=lcc +lat_1=38.88333333333333 +lat_2=37.48333333333333 +lat_0=37 +lon_0=-81 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4801"] = "+proj=lcc +lat_1=46.76666666666667 +lat_2=45.56666666666667 +lat_0=45.16666666666666 +lon_0=-90 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4802"] = "+proj=lcc +lat_1=45.5 +lat_2=44.25 +lat_0=43.83333333333334 +lon_0=-90 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4803"] = "+proj=lcc +lat_1=44.06666666666667 +lat_2=42.73333333333333 +lat_0=42 +lon_0=-90 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4901"] = "+proj=tmerc +lat_0=40.5 +lon_0=-105.1666666666667 +k=0.9999375 +x_0=200000.00001016 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4902"] = "+proj=tmerc +lat_0=40.5 +lon_0=-107.3333333333333 +k=0.9999375 +x_0=399999.99998984 +y_0=99999.99998983997 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4903"] = "+proj=tmerc +lat_0=40.5 +lon_0=-108.75 +k=0.9999375 +x_0=600000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["4904"] = "+proj=tmerc +lat_0=40.5 +lon_0=-110.0833333333333 +k=0.9999375 +x_0=800000.0000101599 +y_0=99999.99998983997 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5001"] = "+proj=omerc +lat_0=57 +lonc=-133.6666666666667 +alpha=-36.86989764583333 +k=0.9999 +x_0=4999999.999999999 +y_0=-4999999.999999999 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5002"] = "+proj=tmerc +lat_0=54 +lon_0=-142 +k=0.9999 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5003"] = "+proj=tmerc +lat_0=54 +lon_0=-146 +k=0.9999 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5004"] = "+proj=tmerc +lat_0=54 +lon_0=-150 +k=0.9999 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5005"] = "+proj=tmerc +lat_0=54 +lon_0=-154 +k=0.9999 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5006"] = "+proj=tmerc +lat_0=54 +lon_0=-158 +k=0.9999 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5007"] = "+proj=tmerc +lat_0=54 +lon_0=-162 +k=0.9999 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5008"] = "+proj=tmerc +lat_0=54 +lon_0=-166 +k=0.9999 +x_0=500000.0000000002 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5010"] = "+proj=lcc +lat_1=51.83333333333334 +lat_2=53.83333333333334 +lat_0=51 +lon_0=-176 +x_0=1000000 +y_0=0 +ellps=GRS80 +datum=NAD83 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5101"] = "+proj=tmerc +lat_0=18.83333333333333 +lon_0=-155.5 +k=0.999966667 +x_0=500000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5102"] = "+proj=tmerc +lat_0=20.33333333333333 +lon_0=-156.6666666666667 +k=0.999966667 +x_0=500000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5103"] = "+proj=tmerc +lat_0=21.16666666666667 +lon_0=-158 +k=0.99999 +x_0=500000.00001016 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"
	zone2params["5104"] = "+proj=tmerc +lat_0=21.83333333333333 +lon_0=-159.5 +k=0.99999 +x_0=500000 +y_0=0 +ellps=GRS80 +to_meter=0.3048006096012192 +no_defs"

	flag.Parse()

	from_coords := C.CString(zone2params[*zone])
	from_pj := C.pj_init_plus(from_coords)

	to_coords := C.CString("+proj=latlong +ellps=WGS84 +datum=WGS84 +no_defs")
	to_pj := C.pj_init_plus(to_coords)

	deg2rad := C.double(0.0174532925)

	if flag.NArg() == 2 {
		easting, _ := strconv.ParseFloat(flag.Arg(0), 64)
		northing, _ := strconv.ParseFloat(flag.Arg(1), 64)
		x := C.double(easting)
		y := C.double(northing)
		C.pj_transform(from_pj, to_pj, 1, 1, &x, &y, nil)
		fmt.Printf("lon=%.10f\tlat=%.10f\n", x/deg2rad, y/deg2rad)
	} else {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "spc2geo --nad83=<zone> <Easting> <Northing>\n")
		flag.PrintDefaults()
	}
}
