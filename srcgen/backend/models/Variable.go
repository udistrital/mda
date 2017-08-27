package models

import (
  "api/db"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

const VariableCollection = "variable"

type Variable struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Tipo	string `json:"tipo"`
  Valor	string `json:"valor"`
  Nombre	string `json:"nombre"`
}


func UpdateVariable(session *mgo.Session, j Variable, id string) error{
	c := db.Cursor(session,VariableCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}


func InsertVariable(session *mgo.Session, j Variable) {
	c := db.Cursor(session,VariableCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllVariables(session *mgo.Session) []Variable {
	c := db.Cursor(session,VariableCollection)
	defer session.Close()
    fmt.Println("Getting all variables")
	var variables []Variable
	err := c.Find(bson.M{}).All(&variables)
	if err != nil {
		fmt.Println(err)
	}
	return variables
}

func GetVariableByName(session *mgo.Session,name string) ([]Variable,error) {
	c := db.Cursor(session, VariableCollection)
	defer session.Close()
	var variables []Variable
	err := c.Find(bson.M{"nombre": name}).All(&variables)
	return variables,err
}

func DeleteVariableById(session *mgo.Session,id string) (string,error) {
	c:= db.Cursor(session, VariableCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok",err
}