package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const VotanteCollection = "votante"

type Votante struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Identificacion	string `json:"identificacion"`
		Nombre	string `json:"nombre"`
		Apellidos	string `json:"apellidos"`
		Codigo	string `json:"codigo"`
		Estamento	string `json:"estamento"`
		Edad	int `json:"edad"`

}


func UpdateVotante(session *mgo.Session, j Votante, id string) error{
	c := db.Cursor(session,VotanteCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertVotante(session *mgo.Session, j Votante) {
	c := db.Cursor(session,VotanteCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllVotantes(session *mgo.Session) []Votante {
	c := db.Cursor(session,VotanteCollection)
	defer session.Close()
    fmt.Println("Getting all votantes")
	var votantes []Votante
	err := c.Find(bson.M{}).All(&votantes)
	if err != nil {
		fmt.Println(err)
	}
	return votantes
}

func GetVotanteByName(session *mgo.Session,name string) ([]Votante,error) {
	c := db.Cursor(session, VotanteCollection)
	defer session.Close()
	var votantes []Votante
	err := c.Find(bson.M{"nombre": name}).All(&votantes)
	return votantes,err
}

func DeleteVotanteById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, VotanteCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}