package service

import(
	"github.com/kataras/iris"
	"gopkg.in/gomail.v2"
	"go.uber.org/zap"
	"net/http"
	
)

//EmailService emailService
type (
	EmailService struct{}
	Response struct{
		Code int `json:"code"`
		Result string `json:"result"`
	}

)

//NewEmailService newEmailService
func NewEmailService() *EmailService{
	return &EmailService{}
}

//Start start
func (e *EmailService)Start(){
	iris.Get("/email", emailSer)
	iris.Post("/email", emailSer)
	iris.Listen(Conf.IrisC.Addr)
}

//Close close
func(e *EmailService)Close(){}


func emailSer(ctx *iris.Context){
	name:=ctx.FormValueString("name")
	if name==""{
		Logger.Error("ctx.FromValueString", zap.String("@ctx.FromValueString","name faild is null"))
		ctx.JSON(http.StatusOK,&Response{
			Code:http.StatusOK,
			Result:"name faild is null",
		})
		return
	}
	email:=ctx.FormValueString("email")
	if email ==""{
		Logger.Error("ctx.FromValueString", zap.String("@ctx.FromValueString","email faild is null"))
		ctx.JSON(http.StatusOK,&Response{
			Code:http.StatusOK,
			Result:"email faild is null",
		})
		return
	}
	txt:=ctx.FormValueString("txt")
	if txt ==""{
		Logger.Error("ctx.FromValueString", zap.String("@ctx.FromValueString","txt faild is null"))
		ctx.JSON(http.StatusOK,&Response{
			Code:http.StatusOK,
			Result:"the txt faild is null",
		})
		return
	}
	m:=gomail.NewMessage()
	m.SetHeader("From", Conf.IrisC.Fromaddr)
	m.SetHeader("To", email)
	m.SetHeader(Conf.IrisC.Sub, name)
	m.SetHeader(Conf.IrisC.Bodytitle, txt)
	push:=gomail.NewDialer(Conf.IrisC.Serviceaddr,Conf.IrisC.Port,Conf.IrisC.Fromaddr,Conf.IrisC.Pwd)
	err:=push.DialAndSend(m)
	if err!=nil{
		Logger.Error("push.DialAndSend", zap.String("@push.DialAndSend",err.Error()))
		ctx.JSON(http.StatusOK, &Response{
			Code:http.StatusOK,
			Result:err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,&Response{
		Code:http.StatusOK,
		Result:"send success!",
	})
}


