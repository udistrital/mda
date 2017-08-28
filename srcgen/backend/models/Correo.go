package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const CorreoCollection = "correo"

type Correo struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Url	string `json:"url"`
  Dominio	string `json:"dominio"`
  Institucional	boolean `json:"institucional"`
}


func UpdateCorreo(session *mgo.Session, j Correo, id string) error{
	c := db.Cursor(session,CorreoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertCorreo(session *mgo.Session, j Correo) {
	c := db.Cursor(session,CorreoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllCorreos(session *mgo.Session) []Correo {
	c := db.Cursor(session,CorreoCollection)
	defer session.Close()
    fmt.Println("Getting all correos")
	var correos []Correo
	err := c.Find(bson.M{}).All(&correos)
	if err != nil {
		fmt.Println(err)
	}
	return correos
}

func GetCorreoByName(session *mgo.Session,name string) ([]Correo,error) {
	c := db.Cursor(session, CorreoCollection)
	defer session.Close()
	var correos []Correo
	err := c.Find(bson.M{"nombre": name}).All(&correos)
	return correos,err
}

func DeleteCorreoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, CorreoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}