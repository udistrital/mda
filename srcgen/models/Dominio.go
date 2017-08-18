package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const DominioCollection = "dominio"

type Dominio struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
			Nombre	string `json:"nombre"`
			Descripcion	string `json:"descripcion"`

}


func UpdateDominio(session *mgo.Session, j Dominio, id string) error{
	c := db.Cursor(session,DominioCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertDominio(session *mgo.Session, j Dominio) {
	c := db.Cursor(session,DominioCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllDominios(session *mgo.Session) []Dominio {
	c := db.Cursor(session,DominioCollection)
	defer session.Close()
    fmt.Println("Getting all dominios")
	var dominios []Dominio
	err := c.Find(bson.M{}).All(&dominios)
	if err != nil {
		fmt.Println(err)
	}
	return dominios
}

func GetDominioByName(session *mgo.Session,name string) ([]Dominio,error) {
	c := db.Cursor(session, DominioCollection)
	defer session.Close()
	var dominios []Dominio
	err := c.Find(bson.M{"nombre": name}).All(&dominios)
	return dominios,err
}

func DeleteDominioById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, DominioCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}