//
//  main.go
//  golang
//
//  Created by zhangzhe on 2014-3-25
//
package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
)

const (
	URL   = "mongodb://192.168.2.73:27017"
	title = "id,longitude,latitude,xoffset,yoffset\r\n"
)

type Ponint struct {
	Type        string
	Coordinates []float64
}
type GoogleGIS struct {
	Id      int32
	Loc     Ponint
	XOffset float64
	YOffset float64
}
type MainController struct {
	beego.Controller
}

func getGoogleData(longitude float64, latitude float64, radius int64) (result string) {

	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	dbNew := session.DB("GISNew")
	collectionNew := dbNew.C("googlegis")

	var id int
	var google GoogleGIS
	var coordinates []float64
	coordinates = append(coordinates, longitude, latitude)
	iter := collectionNew.Find(bson.M{"loc": bson.M{"$near": bson.M{"$geometry": bson.M{"type": "Point", "coordinates": coordinates}, "$maxDistance": radius}}}).Iter()

	result = result + title //标题  第一行
	id = 0

	for iter.Next(&google) {

		longitude01 := int(google.Loc.Coordinates[0] * 1000)
		latitude01 := int(google.Loc.Coordinates[1] * 1000)

		xoffset := google.XOffset
		yoffset := google.YOffset

		var longitude int = int(longitude01 / 10)
		var latitude int = int(latitude01 / 10)

		temp := fmt.Sprintf("%d,%d,%d,%.20f,%.20f\r\n", id, longitude, latitude, xoffset, yoffset)
		result = result + temp
		id++
	}
	fmt.Println("longitude:", longitude, "latitude:", latitude, "Total Num:", id)
	return

}
func (this *MainController) Get() {

	longitudestr := this.Ctx.Params[":longitude"]
	latitudestr := this.Ctx.Params[":latitude"]
	radiusstr := this.Ctx.Params[":radius"]

	longitude, _ := strconv.ParseFloat(longitudestr, 10)
	latitude, _ := strconv.ParseFloat(latitudestr, 10)
	radius, _ := strconv.ParseInt(radiusstr, 10, 64)

	reply := getGoogleData(longitude, latitude, radius)
	this.Ctx.WriteString(reply)
}

func main() {
	beego.Router("/google/:longitude/:latitude/:radius", &MainController{})
	beego.Run()
}
