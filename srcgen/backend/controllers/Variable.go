package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Variable
type VariableController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Variable
// @Failure 403 :objectId is empty
// @router / [get]
func (j *VariableController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllVariables(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Variable by nombre
// @Param	nombre		path 	string	true		"El nombre de la Variable a consultar"
// @Success 200 {object} models.Variable
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *VariableController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		variable, err := models.GetVariableByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = variable
		}
	}
	j.ServeJSON()
}

// @Title Borrar Variable
// @Description Borrar Variable
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *VariableController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteVariableById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Variable
// @Description Crear Variable
// @Param	body		body 	models.Variable	true		"Body para la creacion de Variable"
// @Success 200 {int} Variable.Id
// @Failure 403 body is empty
// @router / [post]
func (j *VariableController) Post() {
	var variable models.Variable
	json.Unmarshal(j.Ctx.Input.RequestBody, &variable)
	fmt.Println(variable)
	session,_ := db.GetSession()
	models.InsertVariable(session,variable)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Variable
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *VariableController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var variable models.Variable
	json.Unmarshal(j.Ctx.Input.RequestBody, &variable)
	session,_ := db.GetSession()

	err := models.UpdateVariable(session, variable,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}