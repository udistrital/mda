package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Dominio
type DominioController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Dominio
// @Failure 403 :objectId is empty
// @router / [get]
func (j *DominioController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllDominios(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Dominio by nombre
// @Param	nombre		path 	string	true		"El nombre de la Dominio a consultar"
// @Success 200 {object} models.Dominio
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *DominioController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		dominio, err := models.GetDominioByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = dominio
		}
	}
	j.ServeJSON()
}

// @Title Borrar Dominio
// @Description Borrar Dominio
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *DominioController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteDominioById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Dominio
// @Description Crear Dominio
// @Param	body		body 	models.Dominio	true		"Body para la creacion de Dominio"
// @Success 200 {int} Dominio.Id
// @Failure 403 body is empty
// @router / [post]
func (j *DominioController) Post() {
	var dominio models.Dominio
	json.Unmarshal(j.Ctx.Input.RequestBody, &dominio)
	fmt.Println(dominio)
	session,_ := db.GetSession()
	models.InsertDominio(session,dominio)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Dominio
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *DominioController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var dominio models.Dominio
	json.Unmarshal(j.Ctx.Input.RequestBody, &dominio)
	session,_ := db.GetSession()

	err := models.UpdateDominio(session, dominio,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}