package controllers

import (
	"dummyV1/models"
	"encoding/json"
	"errors"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type PlayerController struct {
	beego.Controller
}

// @Title GetPLayers
// @Description Get List Of Players From a CSV File
// @Param	body		body 	models.File	true		"body with name of File	"
// @Success 200 {object} []models.Player
// @Failure 400 bad request body
// @Failure 404 file not found
// @Failure 500 internal server error
// @router / [get]
func (t *PlayerController) Get() {
	var file models.File
	err := json.Unmarshal(t.Ctx.Input.CopyBody(1024), &file)
	if (err != nil || file == models.File{}) {
		t.BadRequestError(errors.New("bad request"))
		return
	}

	// open file
	f, err := os.Open(file.FilePath)
	if err != nil {
		t.NoFoundError(errors.New("not found"))
		return
	}
	// remember to close the file at the end of the program
	defer f.Close()

	Response, err := models.GetPlayer(f)
	if err != nil {
		t.InternalServerError(err)
		return
	}

	t.Data["json"] = Response
	t.ServeJSON()
}

// File Not FoundError
func (t *PlayerController) NoFoundError(err error) {
	t.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	t.Ctx.Output.SetStatus(404)
	t.ServeJSON()
}

// Bad Request Error
func (t *PlayerController) BadRequestError(err error) {
	t.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	t.Ctx.Output.SetStatus(400)
	t.ServeJSON()
}

// Internal Server Errors
func (t *PlayerController) InternalServerError(err error) {
	t.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	t.Ctx.Output.SetStatus(500)
	t.ServeJSON()
}

// objectId := o.Ctx.Input.Param(":objectId")
// /:objectId [get]
