package main

import (
	"fmt"

	"github.com/ip2location/ip2location-go"
)

func main() {
	// ip2location.Open("./IPV6-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE.BIN")

	// t1 := time.Now()
	ip := "60.213.47.147"
	ip2location.Open("/home/ubuntu/go/IP2LOCATION-LITE-DB1.BIN")
	fmt.Println(ip2location.Get_country_short(ip).Country_short)

	results := ip2location.Get_all(ip)

	fmt.Printf("country_short: %s\n", results.Country_short)
	// fmt.Printf("country_long: %s\n", results.Country_long)
	// fmt.Printf("region: %s\n", results.Region)
	// fmt.Printf("city: %s\n", results.City)
	// fmt.Printf("isp: %s\n", results.Isp)
	// fmt.Printf("latitude: %f\n", results.Latitude)
	// fmt.Printf("longitude: %f\n", results.Longitude)
	// fmt.Printf("domain: %s\n", results.Domain)
	// fmt.Printf("zipcode: %s\n", results.Zipcode)
	// fmt.Printf("timezone: %s\n", results.Timezone)
	// fmt.Printf("netspeed: %s\n", results.Netspeed)
	// fmt.Printf("iddcode: %s\n", results.Iddcode)
	// fmt.Printf("areacode: %s\n", results.Areacode)
	// fmt.Printf("weatherstationcode: %s\n", results.Weatherstationcode)
	// fmt.Printf("weatherstationname: %s\n", results.Weatherstationname)
	// fmt.Printf("mcc: %s\n", results.Mcc)
	// fmt.Printf("mnc: %s\n", results.Mnc)
	// fmt.Printf("mobilebrand: %s\n", results.Mobilebrand)
	// fmt.Printf("elevation: %f\n", results.Elevation)
	// fmt.Printf("usagetype: %s\n", results.Usagetype)
	// fmt.Printf("api version: %s\n", ip2location.Api_version())

	ip2location.Close()
	// t2 := time.Now()
	// fmt.Println("", t2.Sub(t1))
}
