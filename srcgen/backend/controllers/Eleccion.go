package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Eleccion
type EleccionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Eleccion
// @Failure 403 :objectId is empty
// @router / [get]
func (j *EleccionController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllEleccions(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Eleccion by nombre
// @Param	nombre		path 	string	true		"El nombre de la Eleccion a consultar"
// @Success 200 {object} models.Eleccion
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *EleccionController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		eleccion, err := models.GetEleccionByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = eleccion
		}
	}
	j.ServeJSON()
}

// @Title Borrar Eleccion
// @Description Borrar Eleccion
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *EleccionController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteEleccionById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Eleccion
// @Description Crear Eleccion
// @Param	body		body 	models.Eleccion	true		"Body para la creacion de Eleccion"
// @Success 200 {int} Eleccion.Id
// @Failure 403 body is empty
// @router / [post]
func (j *EleccionController) Post() {
	var eleccion models.Eleccion
	json.Unmarshal(j.Ctx.Input.RequestBody, &eleccion)
	fmt.Println(eleccion)
	session,_ := db.GetSession()
	models.InsertEleccion(session,eleccion)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Eleccion
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *EleccionController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var eleccion models.Eleccion
	json.Unmarshal(j.Ctx.Input.RequestBody, &eleccion)
	session,_ := db.GetSession()

	err := models.UpdateEleccion(session, eleccion,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}