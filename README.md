# Sample EmailService

package mian
import(
   "gopkg.in/gomail.v2"
)
package main(){
  m:=gomail.NewMessage()
	m.SetHeader("From", "xxx@gmail.com")
	m.SetHeader("To", "xxx@gmail.com")
	m.SetHeader("Subject", "xxxx")
	m.SetHeader("text/plain", "xxxxx")
	push:=gomail.NewDialer("smtp.gmail.com",25,"xxx@gmail.com","emailpwd")
	push.DialAndSend(m)
}

