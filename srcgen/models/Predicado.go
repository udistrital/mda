package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const PredicadoCollection = "predicado"

type Predicado struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
			Nombre	string `json:"nombre"`
			Descripcion	string `json:"descripcion"`
	    Tipo []Tipo `json:"tipos"`
	    Dominio []Dominio `json:"dominios"`

}


func UpdatePredicado(session *mgo.Session, j Predicado, id string) error{
	c := db.Cursor(session,PredicadoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertPredicado(session *mgo.Session, j Predicado) {
	c := db.Cursor(session,PredicadoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllPredicados(session *mgo.Session) []Predicado {
	c := db.Cursor(session,PredicadoCollection)
	defer session.Close()
    fmt.Println("Getting all predicados")
	var predicados []Predicado
	err := c.Find(bson.M{}).All(&predicados)
	if err != nil {
		fmt.Println(err)
	}
	return predicados
}

func GetPredicadoByName(session *mgo.Session,name string) ([]Predicado,error) {
	c := db.Cursor(session, PredicadoCollection)
	defer session.Close()
	var predicados []Predicado
	err := c.Find(bson.M{"nombre": name}).All(&predicados)
	return predicados,err
}

func DeletePredicadoById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, PredicadoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}