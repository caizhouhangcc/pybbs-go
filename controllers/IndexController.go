package controllers

import (
	"pybbs-go/filters"
	"pybbs-go/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/sluu99/uuid"
	"golang.org/x/crypto/bcrypt"
)

// verification code
var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

type IndexController struct {
	beego.Controller
}

//首页
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页 | WhereSmile"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	size, _ := beego.AppConfig.Int("page.size")
	s, _ := strconv.Atoi(c.Ctx.Input.Query("s"))
	c.Data["S"] = s
	section := models.Section{Id: s}
	c.Data["Page"] = models.PageTopic(p, size, &section)
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
}

//登录页
func (c *IndexController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		u := models.FindPermissionByUser(1)
		beego.Debug(u)
		c.Data["PageTitle"] = "登录 | WhereSmile"
		c.Layout = "layout/layout.tpl"
		c.TplName = "login.tpl"
	}
}

//验证登录
func (c *IndexController) Login() {
	flash := beego.NewFlash()

	// verify code
	verify := cpt.VerifyReq(c.Ctx.Request)
	logs.Debug("verify: ", verify)
	if !verify {
		c.Data["Success"] = verify
		flash.Error("验证码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	username, password := c.Input().Get("username"), c.Input().Get("password")
	if flag, user := models.Login(username, password); flag {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

//注册页
func (c *IndexController) RegisterPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["PageTitle"] = "注册 | WhereSmile"
		c.Layout = "layout/layout.tpl"
		c.TplName = "register.tpl"
	}
}

//验证注册
func (c *IndexController) Register() {
	flash := beego.NewFlash()

	// verify code
	verify := cpt.VerifyReq(c.Ctx.Request)
	logs.Debug("verify: ", verify)
	if !verify {
		c.Data["Success"] = verify
		flash.Error("验证码错误")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
		return
	}

	username := c.Input().Get("username")
	nickname := c.Input().Get("nickname")
	password := c.Input().Get("password")

	if len(username) == 0 || len(nickname) == 0 || len(password) == 0 {
		flash.Error("用户名、昵称或密码不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if flag, _ := models.FindUserByUserName(username); flag {
		flash.Error("用户名已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		// password
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			flash.Error("密码有毒，请重试一遍=。=")
			flash.Store(&c.Controller)
			c.Redirect("/register", 302)
			return
		}

		var token = uuid.Rand().Hex()
		user := models.User{
			Username: username,
			Nickname: nickname,
			Password: string(hash),
			Avatar:   "/static/imgs/avatar.png",
			Token:    token}
		models.SaveUser(&user)
		// others are ordered as cookie's max age time, path,domain, secure and httponly.
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
	}
}

//登出
func (c *IndexController) Logout() {
	c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	c.Redirect("/", 302)
}

//关于
func (c *IndexController) About() {
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "关于 | WhereSmile"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}
