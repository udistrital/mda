package controllers

import (
	"api/models"
	"github.com/astaxie/beego"
	"encoding/json"
	_  "gopkg.in/mgo.v2"
	"fmt"
)

// Operaciones Crud Votante
type VotanteController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Votante
// @Failure 403 :objectId is empty
// @router / [get]
func (j *VotanteController) GetAll() {
	session,_ := db.GetSession()
	obs := models.GetAllVotantes(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Votante by nombre
// @Param	nombre		path 	string	true		"El nombre de la Votante a consultar"
// @Success 200 {object} models.Votante
// @Failure 403 :uid is empty
// @router /:nombre [get]
func (j *VotanteController) Get() {
	name := j.GetString(":nombre")
	session, _ := db.GetSession()
	if name != "" {
		votante, err := models.GetVotanteByName(session,name)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = votante
		}
	}
	j.ServeJSON()
}

// @Title Borrar Votante
// @Description Borrar Votante
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *VotanteController) Delete() {
	session,_ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteVotanteById(session,objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Votante
// @Description Crear Votante
// @Param	body		body 	models.Votante	true		"Body para la creacion de Votante"
// @Success 200 {int} Votante.Id
// @Failure 403 body is empty
// @router / [post]
func (j *VotanteController) Post() {
	var votante models.Votante
	json.Unmarshal(j.Ctx.Input.RequestBody, &votante)
	fmt.Println(votante)
	session,_ := db.GetSession()
	models.InsertVotante(session,votante)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Votante
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *VotanteController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var votante models.Votante
	json.Unmarshal(j.Ctx.Input.RequestBody, &votante)
	session,_ := db.GetSession()

	err := models.UpdateVotante(session, votante,objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}