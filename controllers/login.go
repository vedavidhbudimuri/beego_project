package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type UserLoginDetails struct {
	Username string
	Password string
}

type AuthCreds struct {
	AccessToken  string
	RefreshToken string
}

func (controller *LoginController) Post() {

	var userLoginDetails UserLoginDetails
	fmt.Println(controller.Ctx.Input.RequestBody)
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &userLoginDetails)
	if err != nil {
		controller.Ctx.Output.SetStatus(400)
		controller.Ctx.Output.Body([]byte("Invalid Request Data"))
		return
	}
	fmt.Println(userLoginDetails.Username)
	fmt.Println(userLoginDetails.Password)
	authCreds := AuthCreds{AccessToken: "Access Token", RefreshToken: "Refresh Token"}

	fmt.Println("At Login Controller")
	fmt.Println(authCreds)

	controller.Data["json"] = authCreds
	controller.ServeJSON()
}
