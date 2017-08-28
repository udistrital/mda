package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Direccion
type DireccionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Direccion
// @Failure 403 :objectId is empty
// @router / [get]
func (j *DireccionController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllDireccions(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Direccion by nombre
// @Param	nombre		path 	string	true		"El nombre de la Direccion a consultar"
// @Success 200 {object} models.Direccion
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *DireccionController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		direccion, err := models.GetDireccionByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = direccion
		}
	}
	j.ServeJSON()
}

// @Title Borrar Direccion
// @Description Borrar Direccion
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *DireccionController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteDireccionById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Direccion
// @Description Crear Direccion
// @Param	body		body 	models.Direccion	true		"Body para la creacion de Direccion"
// @Success 200 {int} Direccion.Id
// @Failure 403 body is empty
// @router / [post]
func (j *DireccionController) Post() {
	var direccion models.Direccion
	json.Unmarshal(j.Ctx.Input.RequestBody, &direccion)
	fmt.Println(direccion)
	session,_ := db.GetSession()
	models.InsertDireccion(session,direccion)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Direccion
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *DireccionController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var direccion models.Direccion
	json.Unmarshal(j.Ctx.Input.RequestBody, &direccion)
	session,_ := db.GetSession()

	err := models.UpdateDireccion(session, direccion,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}