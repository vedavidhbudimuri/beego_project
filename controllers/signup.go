package controllers

import (
	"beego_project/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type SignUpController struct {
	beego.Controller
}

type UserDetails struct {
	Username   string
	Password   string
	ProfilePic string
}

func (controller *SignUpController) Post() {

	var userDetails UserDetails
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &userDetails)

	if err != nil {
		controller.Ctx.Output.SetStatus(400)
		controller.Ctx.Output.Body([]byte("Invalid Request Data"))
		return
	}

	storageImplementation := models.Storage{}
	_, err = signUp(userDetails, &storageImplementation)
	switch err.(type) {
	case *UserNameAlreadyTaken:
		controller.Ctx.Output.SetStatus(400)
		controller.Ctx.Output.Body([]byte("Username already taken"))
	case *ExpectationFailed:
		controller.Ctx.Output.SetStatus(417)
		controller.Ctx.Output.Body([]byte("Unexpected failure at database"))
	}
	controller.Data["json"] = map[string]bool{"success": true}
	controller.ServeJSON()
}

func signUp(userDetails UserDetails, storage StorageInterface) (bool, error) {

	isUserNameTaken := storage.IsUsernameTaken(userDetails.Username)
	fmt.Println("isUserNameTaken", isUserNameTaken)
	if isUserNameTaken {
		return false, &UserNameAlreadyTaken{}
	}

	err := storage.CreateUser(userDetails.Username, userDetails.Password, userDetails.ProfilePic)

	if err != nil {
		return false, &ExpectationFailed{"Failed Creating User"}
	}

	return true, nil

}
