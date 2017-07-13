package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const DireccionCollection = "direccion"

type Direccion struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
			calle	string `json:"calle"`
			ciudad	string `json:"ciudad"`
			pais	string `json:"pais"`

}


func UpdateDireccion(session *mgo.Session, j Direccion, id string) error{
	c := Cursor(session,DireccionCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertDireccion(session *mgo.Session, j Direccion) {
	c := Cursor(session,DireccionCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllDireccions(session *mgo.Session) []Direccion {
	c := Cursor(session,DireccionCollection)
	defer session.Close()
    fmt.Println("Getting all direccions")
	var direccions []Direccion
	err := c.Find(bson.M{}).All(&direccions)
	if err != nil {
		fmt.Println(err)
	}
	return direccions
}

func GetDireccionByName(session *mgo.Session,name string) ([]Direccion,error) {
	c := Cursor(session, DireccionCollection)
	defer session.Close()
	var direccions []Direccion
	err := c.Find(bson.M{"nombre": name}).All(&direccions)
	return direccions,err
}

func DeleteDireccionById(session *mgo.Session,id string) (string,error) {
	c:= Cursor(session, JornadaCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}