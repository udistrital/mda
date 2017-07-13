package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const PersonaCollection = "persona"

type Persona struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
			nombre	string `json:"nombre"`
			apellido	string `json:"apellido"`
	    direccion []Direccion `json:"direccions"`
			edad	integer `json:"edad"`
	    telefono []Telefono `json:"telefonos"`
			correo	string `json:"correo"`

}


func UpdatePersona(session *mgo.Session, j Persona, id string) error{
	c := Cursor(session,PersonaCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertPersona(session *mgo.Session, j Persona) {
	c := Cursor(session,PersonaCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllPersonas(session *mgo.Session) []Persona {
	c := Cursor(session,PersonaCollection)
	defer session.Close()
    fmt.Println("Getting all personas")
	var personas []Persona
	err := c.Find(bson.M{}).All(&personas)
	if err != nil {
		fmt.Println(err)
	}
	return personas
}

func GetPersonaByName(session *mgo.Session,name string) ([]Persona,error) {
	c := Cursor(session, PersonaCollection)
	defer session.Close()
	var personas []Persona
	err := c.Find(bson.M{"nombre": name}).All(&personas)
	return personas,err
}

func DeletePersonaById(session *mgo.Session,id string) (string,error) {
	c:= Cursor(session, JornadaCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}