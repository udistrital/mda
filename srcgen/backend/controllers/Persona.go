package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Persona
type PersonaController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Persona
// @Failure 403 :objectId is empty
// @router / [get]
func (j *PersonaController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllPersonas(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Persona by nombre
// @Param	nombre		path 	string	true		"El nombre de la Persona a consultar"
// @Success 200 {object} models.Persona
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *PersonaController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		persona, err := models.GetPersonaByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = persona
		}
	}
	j.ServeJSON()
}

// @Title Borrar Persona
// @Description Borrar Persona
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *PersonaController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeletePersonaById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Persona
// @Description Crear Persona
// @Param	body		body 	models.Persona	true		"Body para la creacion de Persona"
// @Success 200 {int} Persona.Id
// @Failure 403 body is empty
// @router / [post]
func (j *PersonaController) Post() {
	var persona models.Persona
	json.Unmarshal(j.Ctx.Input.RequestBody, &persona)
	fmt.Println(persona)
	session,_ := db.GetSession()
	models.InsertPersona(session,persona)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Persona
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *PersonaController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var persona models.Persona
	json.Unmarshal(j.Ctx.Input.RequestBody, &persona)
	session,_ := db.GetSession()

	err := models.UpdatePersona(session, persona,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}