// code snippet for finding the distance between two latitube and longitude points
package main

import (
	"fmt"
	"math"
	"sort"
)
 
func distance(long1 float64,long2 float64, latt1 float64,latt2 float64,unit string) float64{
	if latt1 == latt2 && long1 == long2{ // checks if the lattitude and longitiude points are the same
		fmt.Println(0) // if they are the same then the distance is 0
	}
	radlat1 := math.Pi * latt1/180	// converts the lattitudes to radians
	radlat2 := math.Pi * latt2/180
	theta := long1 - long2
	radtheta := math.Pi * theta/180
	dist := math.Sin(radlat1) * math.Sin(radlat2) * math.Cos(radlat1) * math.Cos(radtheta)
	// sin(lt1)*sin(lt2)*cos(lt1)*cos(theta)
	if dist > 1{
		dist = 1
	}
	dist = math.Acos(dist)
	dist = dist * 180/math.Pi // changing from radians to degrees
	dist = dist * 60 * 1.1515	// converting to miles
	if unit == "k"{
		dist = dist * 1.609344 // converting to kilometers
	}
	if unit == "N"{
		dist = dist * 0.8684 // converting to nautical miles
	}
	return dist
}
func main(){
	// a simple test
	// creates a type dictionary
	// creates a list of type dictionary which would hold maps
	long2 := 345.06
	latt2 := 12.99
	type Dictionary map[string]interface{}
	data := []Dictionary{}
	dict1 := Dictionary{"name":"zed","class":1,"long":123.45,"lat":26.765,"id":1}
	dict2 := Dictionary{"name":"zod","class":1,"long":13.45,"lat":6.765,"id":2}
	data = append(data,dict1,dict2)
	unit := "k"
	dist := []float64{}
	// loop through the list of maps and call the distance function
	for _,i := range data{
		long1 := i["long"].(float64)
		latt1 := i["lat"].(float64)
		_distance := distance(long1,long2,latt1,latt2,unit)
		dist = append(dist,_distance)
	}
	// sort the list to get the lowest value
	sort.Sort(sort.Float64Slice(dist))
	fmt.Println(dist)
	fmt.Println(dist[0],"is closest to you, should i book an appointment?")
	// speed = distance / time
}