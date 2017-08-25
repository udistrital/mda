package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const PonderacionCollection = "ponderacion"

type Ponderacion struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
			Estamento	string `json:"estamento"`
			Porcentaje	string `json:"porcentaje"`
			Nombre	string `json:"nombre"`

}


func UpdatePonderacion(session *mgo.Session, j Ponderacion, id string) error{
	c := db.Cursor(session,PonderacionCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertPonderacion(session *mgo.Session, j Ponderacion) {
	c := db.Cursor(session,PonderacionCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllPonderacions(session *mgo.Session) []Ponderacion {
	c := db.Cursor(session,PonderacionCollection)
	defer session.Close()
    fmt.Println("Getting all ponderacions")
	var ponderacions []Ponderacion
	err := c.Find(bson.M{}).All(&ponderacions)
	if err != nil {
		fmt.Println(err)
	}
	return ponderacions
}

func GetPonderacionByName(session *mgo.Session,name string) ([]Ponderacion,error) {
	c := db.Cursor(session, PonderacionCollection)
	defer session.Close()
	var ponderacions []Ponderacion
	err := c.Find(bson.M{"nombre": name}).All(&ponderacions)
	return ponderacions,err
}

func DeletePonderacionById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, PonderacionCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}