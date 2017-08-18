package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Tipo
type TipoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Tipo
// @Failure 403 :objectId is empty
// @router / [get]
func (j *TipoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllTipos(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Tipo by nombre
// @Param	nombre		path 	string	true		"El nombre de la Tipo a consultar"
// @Success 200 {object} models.Tipo
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *TipoController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		tipo, err := models.GetTipoByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = tipo
		}
	}
	j.ServeJSON()
}

// @Title Borrar Tipo
// @Description Borrar Tipo
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *TipoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteTipoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Tipo
// @Description Crear Tipo
// @Param	body		body 	models.Tipo	true		"Body para la creacion de Tipo"
// @Success 200 {int} Tipo.Id
// @Failure 403 body is empty
// @router / [post]
func (j *TipoController) Post() {
	var tipo models.Tipo
	json.Unmarshal(j.Ctx.Input.RequestBody, &tipo)
	fmt.Println(tipo)
	session,_ := db.GetSession()
	models.InsertTipo(session,tipo)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Tipo
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *TipoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var tipo models.Tipo
	json.Unmarshal(j.Ctx.Input.RequestBody, &tipo)
	session,_ := db.GetSession()

	err := models.UpdateTipo(session, tipo,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}