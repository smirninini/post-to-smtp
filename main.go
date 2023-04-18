package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var smtpConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	From           string
	To             []string
	DefaultSubject string
}

var corsAllowOrigins []string

func main() {
	r := gin.Default()
	r.POST("/send", sendHandler)
	r.Run()
}

func init() {
	//check .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	smtpConfig.Host = os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	smtpConfig.Port = int(smtpPort)
	smtpConfig.User = os.Getenv("SMTP_USER")
	smtpConfig.Password = os.Getenv("SMTP_PASSWORD")
	smtpConfig.From = os.Getenv("SMTP_FROM")
	if !isValidEmail(smtpConfig.From) {
		log.Fatal("Not valid SMTP 'From' email address")
	}
	smtpConfig.To = strings.Split(os.Getenv("SMTP_TO"), ",")
	for _, s := range smtpConfig.To {
		if !isValidEmail(s) {
			log.Fatal("Not valid SMTP 'To' email address or addreses")
		}
	}
	smtpConfig.DefaultSubject = os.Getenv("SMTP_DEFAULT_SUBJECT")
	log.SetFormatter(&log.JSONFormatter{})
}

// validEmail
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func sendHandler(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	body := string(bytes)
	log.Info("Processing " + body)
	m := gomail.NewMessage()
	m.SetHeader("From", smtpConfig.From)
	m.SetHeader("To", smtpConfig.To...)
	m.SetHeader("Subject", smtpConfig.DefaultSubject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpConfig.Host, int(smtpConfig.Port), smtpConfig.User, smtpConfig.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
