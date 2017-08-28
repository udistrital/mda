package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Telefono
type TelefonoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Telefono
// @Failure 403 :objectId is empty
// @router / [get]
func (j *TelefonoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllTelefonos(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Telefono by nombre
// @Param	nombre		path 	string	true		"El nombre de la Telefono a consultar"
// @Success 200 {object} models.Telefono
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *TelefonoController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		telefono, err := models.GetTelefonoByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = telefono
		}
	}
	j.ServeJSON()
}

// @Title Borrar Telefono
// @Description Borrar Telefono
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *TelefonoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteTelefonoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Telefono
// @Description Crear Telefono
// @Param	body		body 	models.Telefono	true		"Body para la creacion de Telefono"
// @Success 200 {int} Telefono.Id
// @Failure 403 body is empty
// @router / [post]
func (j *TelefonoController) Post() {
	var telefono models.Telefono
	json.Unmarshal(j.Ctx.Input.RequestBody, &telefono)
	fmt.Println(telefono)
	session,_ := db.GetSession()
	models.InsertTelefono(session,telefono)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Telefono
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *TelefonoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var telefono models.Telefono
	json.Unmarshal(j.Ctx.Input.RequestBody, &telefono)
	session,_ := db.GetSession()

	err := models.UpdateTelefono(session, telefono,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}