package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const TelefonoCollection = "telefono"

type Telefono struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Indicador	string `json:"indicador"`
  Numero	string `json:"numero"`
}


func UpdateTelefono(session *mgo.Session, j Telefono, id string) error{
	c := db.Cursor(session,TelefonoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertTelefono(session *mgo.Session, j Telefono) {
	c := db.Cursor(session,TelefonoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllTelefonos(session *mgo.Session) []Telefono {
	c := db.Cursor(session,TelefonoCollection)
	defer session.Close()
    fmt.Println("Getting all telefonos")
	var telefonos []Telefono
	err := c.Find(bson.M{}).All(&telefonos)
	if err != nil {
		fmt.Println(err)
	}
	return telefonos
}

func GetTelefonoByName(session *mgo.Session,name string) ([]Telefono,error) {
	c := db.Cursor(session, TelefonoCollection)
	defer session.Close()
	var telefonos []Telefono
	err := c.Find(bson.M{"nombre": name}).All(&telefonos)
	return telefonos,err
}

func DeleteTelefonoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, TelefonoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}