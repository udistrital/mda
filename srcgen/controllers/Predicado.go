package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Predicado
type PredicadoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Predicado
// @Failure 403 :objectId is empty
// @router / [get]
func (j *PredicadoController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllPredicados(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Predicado by nombre
// @Param	nombre		path 	string	true		"El nombre de la Predicado a consultar"
// @Success 200 {object} models.Predicado
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *PredicadoController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		predicado, err := models.GetPredicadoByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = predicado
		}
	}
	j.ServeJSON()
}

// @Title Borrar Predicado
// @Description Borrar Predicado
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *PredicadoController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeletePredicadoById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Predicado
// @Description Crear Predicado
// @Param	body		body 	models.Predicado	true		"Body para la creacion de Predicado"
// @Success 200 {int} Predicado.Id
// @Failure 403 body is empty
// @router / [post]
func (j *PredicadoController) Post() {
	var predicado models.Predicado
	json.Unmarshal(j.Ctx.Input.RequestBody, &predicado)
	fmt.Println(predicado)
	session,_ := db.GetSession()
	models.InsertPredicado(session,predicado)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Predicado
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *PredicadoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var predicado models.Predicado
	json.Unmarshal(j.Ctx.Input.RequestBody, &predicado)
	session,_ := db.GetSession()

	err := models.UpdatePredicado(session, predicado,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}