package tools

import (
	"RandomMusic/dataSource"
	"RandomMusic/model"
	"crypto/md5"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// AddMD5 用户密码MD5加密
func AddMD5(text string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	return md5Str
}

// SwitchTimeStampToData 将传入的时间戳转为时间
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

// VerifyEmailFormat 正则验证邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// SendEmailCode 发送邮箱验证码
func SendEmailCode(em string) {
	//定义收件人
	mailTo := []string{em}

	//读取配置文件
	emailConf := dataSource.LoadEmailConf()
	//邮件主题为"Hello"
	subject := emailConf.Title
	// 邮件正文
	body := emailConf.Body

	err := SendMail(mailTo, subject, body, em)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}
}

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, body string, em string) error {
	//读取配置文件
	emailConf := dataSource.LoadEmailConf()
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": emailConf.Username,
		"pass": emailConf.Password,
		"host": emailConf.Host,
		"port": emailConf.Port,
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], emailConf.Sender)) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	//设置随机数
	num := rand.Intn(999999)
	m.SetBody("text/html", body+strconv.Itoa(num)) //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err == nil {
		//验证码写入数据库
		var emailCode model.UserEmailCode
		emailCode.Email = em
		emailCode.Code = strconv.Itoa(num)
		t := time.Now().Unix()
		t += 7200
		emailCode.EndTime = int(t)
		dataSource.DB.Create(&emailCode)
	}
	return err
}
