// Package controllers is for mapping controller functions to endpoints in routers.
package controllers

import (
	groot "beego-standard-layout"
	"beego-standard-layout/inmemory"
	"beego-standard-layout/json"
	"beego-standard-layout/service"
	"github.com/astaxie/beego"
)

// Operations about struct
type StructController struct {
	beego.Controller
}

// @Description Create Struct
// @Param	body		body 	groot.StructRequest		true	"The struct content, e.q {"value": 1}"
// @Success 200 {result} groot.Struct
// @Failure 400 bad request
// @Failure 500 internal server error
// @router / [post]
func (s *StructController) Create() {
	var req groot.StructRequest
	err := json.ConvertRequest(s.Ctx.Input.RequestBody, &req)
	if err != nil {
		s.CustomAbort(400, "bad request")
	}

	//TODO must be able to switch inmemory / postgres impl (possibly using conf)
	repo := inmemory.NewStructRepository()
	svc := service.NewStructService(repo)
	resp := svc.CreateStruct(req)

	if resp.Status != 200 {
		body := resp.Result.(string)
		s.CustomAbort(resp.Status, body)
	}

	s.Data["json"] = resp.Result
	s.ServeJSON()
}
