// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tGorpController struct {}
var GorpController tGorpController


func (_ tGorpController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Begin", args).Url
}

func (_ tGorpController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Commit", args).Url
}

func (_ tGorpController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Rollback", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tApplication struct {}
var Application tApplication


func (_ tApplication) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.CheckUser", args).Url
}

func (_ tApplication) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.AddUser", args).Url
}

func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).Url
}

func (_ tApplication) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Register", args).Url
}

func (_ tApplication) SaveUser(
		user interface{},
		verifyPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "verifyPassword", verifyPassword)
	return revel.MainRouter.Reverse("Application.SaveUser", args).Url
}

func (_ tApplication) LoginIndex(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.LoginIndex", args).Url
}

func (_ tApplication) Login(
		username string,
		password string,
		remember bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "remember", remember)
	return revel.MainRouter.Reverse("Application.Login", args).Url
}

func (_ tApplication) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Logout", args).Url
}


type tWxApp struct {}
var WxApp tWxApp


func (_ tWxApp) Wx(
		signature string,
		timestamp string,
		nonce string,
		echostr string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "signature", signature)
	revel.Unbind(args, "timestamp", timestamp)
	revel.Unbind(args, "nonce", nonce)
	revel.Unbind(args, "echostr", echostr)
	return revel.MainRouter.Reverse("WxApp.Wx", args).Url
}

func (_ tWxApp) WxP(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("WxApp.WxP", args).Url
}

func (_ tWxApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("WxApp.Index", args).Url
}


type tProduct struct {}
var Product tProduct


func (_ tProduct) Detail(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Product.Detail", args).Url
}


type tSellers struct {}
var Sellers tSellers


func (_ tSellers) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sellers.Index", args).Url
}


