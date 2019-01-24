package router

import (
	"html/template"
	"net/http"
	"github.com/renosyah/CustomAuth/Golang/model"
	"fmt"
	"github.com/spf13/viper"
	"github.com/renosyah/CustomAuth/Golang/auth"
)

type Auth_Holder struct {
	Auth *auth.Auth_Server
}

func RegisterUserPage(res http.ResponseWriter,req *http.Request){
	page, err := template.ParseFiles(viper.GetString("template.customer_register"),
		viper.GetString("template.customer_header"),
		viper.GetString("template.customer_footer"))
	if err != nil {
		fmt.Println(err.Error())
	}
	page.Execute(res, nil)
}

func LoginUserPage(res http.ResponseWriter,req *http.Request){

	id_callback := "" + req.FormValue("id_callback")

	page, err := template.ParseFiles(viper.GetString("template.customer_login"),
		viper.GetString("template.customer_header"),
		viper.GetString("template.customer_footer"))

	if err != nil {
		fmt.Println(err.Error())
	}
	page.Execute(res, map[string]string{"id_callback" : id_callback})
}

func RegisterUser(res http.ResponseWriter,req *http.Request){

	user := &model.User{
		Id : RandStringRunes(10),
		Name:req.FormValue("name"),
		Email:req.FormValue("email"),
		Username: req.FormValue("username"),
		Password:req.FormValue("password"),
	}

	err := user.AddUser(req.Context(),dbPool)
	if err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(res,req,"/",301)
}

func (a *Auth_Holder) LoginUser(res http.ResponseWriter,req *http.Request){

	id_callback := req.FormValue("id_callback")

	user := &model.User{
		Email:req.FormValue("username"),
		Username: req.FormValue("username"),
		Password:req.FormValue("password"),
	}

	result,err := user.LoginUser(req.Context(),dbPool)
	if err != nil {
		fmt.Println(err.Error())
	}


	if err == nil && len(id_callback) > 0 {

		a.Auth.Broadcast <- auth.CallbackData{
			IdCallback:id_callback,
			User:&auth.UserData{
				Id:result.Id,
				Name: result.Name,
				Email:result.Email,
				Password:"",
				Username:result.Username,
			},
		}

		http.Redirect(res,req,fmt.Sprintf("/?id_callback=%s",id_callback),301)

	} else if err == nil && len(id_callback) < 0 {
		SetCookie(costumer_session, result.Id, res)
		http.Redirect(res, req, "/", 301)

	} else if err != nil && len(id_callback) > 0 {
		http.Redirect(res,req,fmt.Sprintf("/?id_callback=%s",id_callback),301)

	} else if err != nil && len(id_callback) < 0 {
		http.Redirect(res, req, "/", 301)

	}
}