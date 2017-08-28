package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Correo
type CorreoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Correo
// @Failure 403 :objectId is empty
// @router / [get]
func (j *CorreoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllCorreos(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Correo by nombre
// @Param	nombre		path 	string	true		"El nombre de la Correo a consultar"
// @Success 200 {object} models.Correo
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *CorreoController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		correo, err := models.GetCorreoByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = correo
		}
	}
	j.ServeJSON()
}

// @Title Borrar Correo
// @Description Borrar Correo
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *CorreoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteCorreoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Correo
// @Description Crear Correo
// @Param	body		body 	models.Correo	true		"Body para la creacion de Correo"
// @Success 200 {int} Correo.Id
// @Failure 403 body is empty
// @router / [post]
func (j *CorreoController) Post() {
	var correo models.Correo
	json.Unmarshal(j.Ctx.Input.RequestBody, &correo)
	fmt.Println(correo)
	session,_ := db.GetSession()
	models.InsertCorreo(session,correo)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Correo
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *CorreoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var correo models.Correo
	json.Unmarshal(j.Ctx.Input.RequestBody, &correo)
	session,_ := db.GetSession()

	err := models.UpdateCorreo(session, correo,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}