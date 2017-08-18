package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const TipoCollection = "tipo"

type Tipo struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
			Nombre	string `json:"nombre"`
			Descripcion	string `json:"descripcion"`

}


func UpdateTipo(session *mgo.Session, j Tipo, id string) error{
	c := db.Cursor(session,TipoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertTipo(session *mgo.Session, j Tipo) {
	c := db.Cursor(session,TipoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllTipos(session *mgo.Session) []Tipo {
	c := db.Cursor(session,TipoCollection)
	defer session.Close()
    fmt.Println("Getting all tipos")
	var tipos []Tipo
	err := c.Find(bson.M{}).All(&tipos)
	if err != nil {
		fmt.Println(err)
	}
	return tipos
}

func GetTipoByName(session *mgo.Session,name string) ([]Tipo,error) {
	c := db.Cursor(session, TipoCollection)
	defer session.Close()
	var tipos []Tipo
	err := c.Find(bson.M{"nombre": name}).All(&tipos)
	return tipos,err
}

func DeleteTipoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, TipoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}