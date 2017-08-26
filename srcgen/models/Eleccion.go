package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const EleccionCollection = "eleccion"

type Eleccion struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Nombre	string `json:"nombre"`
		Votantes	Votante `json:"votantes"`
		Fechainicio	Time.time `json:"fechainicio"`
		Fechafinal	Time.time `json:"fechafinal"`
		Ponderaciones	Time.time `json:"ponderaciones"`
		Responsables	Time.time `json:"responsables"`
		Habilitado	Time.time `json:"habilitado"`

}


func UpdateEleccion(session *mgo.Session, j Eleccion, id string) error{
	c := db.Cursor(session,EleccionCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertEleccion(session *mgo.Session, j Eleccion) {
	c := db.Cursor(session,EleccionCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllEleccions(session *mgo.Session) []Eleccion {
	c := db.Cursor(session,EleccionCollection)
	defer session.Close()
    fmt.Println("Getting all eleccions")
	var eleccions []Eleccion
	err := c.Find(bson.M{}).All(&eleccions)
	if err != nil {
		fmt.Println(err)
	}
	return eleccions
}

func GetEleccionByName(session *mgo.Session,name string) ([]Eleccion,error) {
	c := db.Cursor(session, EleccionCollection)
	defer session.Close()
	var eleccions []Eleccion
	err := c.Find(bson.M{"nombre": name}).All(&eleccions)
	return eleccions,err
}

func DeleteEleccionById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, EleccionCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}