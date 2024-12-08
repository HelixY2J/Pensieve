package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/smtp2go-oss/smtp2go-go"
)

type Smtp2goMailer struct {
	FromMail string
	//Clientele *sendgrid.Client
}

func NewSmtp2goMailer(fromMail string) *Smtp2goMailer {

	return &Smtp2goMailer{
		FromMail: fromMail,
		//Clientele: client,
	}
}

func (m *Smtp2goMailer) SendInsights(insights []*DailyInsight, u *User) error {
	if u.Email == "" {
		return fmt.Errorf("user has no email")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
	key := os.Getenv("SMTP2GO_API_KEY")
	if key == "" {
		log.Fatal("SMTP2GO_API_KEY must be set")
	}

	email := smtp2go.Email{
		From: m.FromMail,
		To: []string{
			u.Email,
		},
		Subject:  "Dive into your Daily Insight(s)",
		TextBody: "Pensive.",
		HtmlBody: BuildInsightsMailTemplate(u, insights),
	}

	res, err := smtp2go.Send(&email)
	if err != nil {
		fmt.Printf("Something broke : %s", err)
	}
	fmt.Printf("Mail sent: %s", res)

	return nil
}

func BuildInsightsMailTemplate(u *User, ins []*DailyInsight) string {
	templ, err := template.ParseFiles("daily.tmpl")
	if err != nil {
		panic(err)
	}

	payload := struct {
		User     *User
		Insights []*DailyInsight
	}{
		User:     u,
		Insights: ins,
	}

	var out bytes.Buffer
	err = templ.Execute(&out, payload)
	if err != nil {
		panic(err)
	}

	return out.String()
}
