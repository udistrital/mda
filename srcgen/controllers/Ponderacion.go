package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Ponderacion
type PonderacionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Ponderacion
// @Failure 403 :objectId is empty
// @router / [get]
func (j *PonderacionController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllPonderacions(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Ponderacion by nombre
// @Param	nombre		path 	string	true		"El nombre de la Ponderacion a consultar"
// @Success 200 {object} models.Ponderacion
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *PonderacionController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		ponderacion, err := models.GetPonderacionByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = ponderacion
		}
	}
	j.ServeJSON()
}

// @Title Borrar Ponderacion
// @Description Borrar Ponderacion
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *PonderacionController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeletePonderacionById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Ponderacion
// @Description Crear Ponderacion
// @Param	body		body 	models.Ponderacion	true		"Body para la creacion de Ponderacion"
// @Success 200 {int} Ponderacion.Id
// @Failure 403 body is empty
// @router / [post]
func (j *PonderacionController) Post() {
	var ponderacion models.Ponderacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &ponderacion)
	fmt.Println(ponderacion)
	session,_ := db.GetSession()
	models.InsertPonderacion(session,ponderacion)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Ponderacion
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *PonderacionController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var ponderacion models.Ponderacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &ponderacion)
	session,_ := db.GetSession()

	err := models.UpdatePonderacion(session, ponderacion,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}