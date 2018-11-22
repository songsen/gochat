package controllers

import "fmt"
import (
        "github.com/astaxie/beego"
        "github.com/gislu/gochat/src/models"
        "encoding/xml"
        "time"
        "io/ioutil"
        "net/http"
        "strings"
)

type PubController struct  {
	beego.Controller
}


//TODO 这里负责回调模式的验证
func (c * PubController) Get() {
	var signature models.Signature
	if err := c.ParseForm(&signature) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}

	fmt.Println(signature.Echostr)
	c.Ctx.WriteString(signature.Echostr)

}




func (c * PubController) Post(){
	var msgIn models.PubTextMsg
	err := xml.Unmarshal(c.Ctx.Input.RequestBody,&msgIn)
	if err != nil {
		Lg(err)
		c.Abort("400")
		return
	}
	msgback := "这是自动回复"
	if(msgIn.MsgType == "event"){
		msgback = "感谢您的关注(O w O)～～"
		_ = c.PubSendBack(msgback,msgIn)
		return 
	}else if(strings.HasPrefix(msgIn.Content,"查询") == true){
		sentence := strings.Replace(msgIn.Content,"查询","",1)
		sentence = strings.TrimSpace(sentence)
		msgback =  RobotApi(sentence + "的做法")
	}else if(strings.HasPrefix(msgIn.Content,"温度") == true){
		msgback =  models.Readmm("mqtt_shm")
	}
	_ = c.PubSendBack(msgback,msgIn)
}

func RobotApi(keymsg string) string{
	url := "http://api.douqq.com/?key=PUVLKzdjeDduTWNHUFVXQUU3PWhRTytOekFrQUFBPT0&msg=" + keymsg
	resp,err :=http.Get(url)
	if err!=nil{
		fmt.Print(err)
	}
	fetchrs,err :=ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(fetchrs)
}

func (this * PubController)PubSendBack(backMsg string,msgIn models.PubTextMsg)error{
	msgOut := models.PubTextOut{
		ToUserName:msgIn.FromUserName,
		FromUserName:msgIn.ToUserName,
		CreateTime:time.Now().Unix(),
		MsgType:"text",
		Content:fmt.Sprint(backMsg),
	}

	xmlData ,err := msgOut.ToXml()
	this.Ctx.WriteString(string(xmlData))
	return err
}



