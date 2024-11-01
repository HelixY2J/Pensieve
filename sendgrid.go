package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var FromName = "Pensieve"

type SendGridMailer struct {
	FromMail  string
	Clientele *sendgrid.Client
}

func NewSendGridMailer(apikey, fromMail string) *SendGridMailer {
	client := sendgrid.NewSendClient(apikey)

	return &SendGridMailer{
		FromMail:  fromMail,
		Clientele: client,
	}
}

func (m *SendGridMailer) SendInsights(insights []*DailyInsight, u *User) error {
	if u.Email == "" {
		return fmt.Errorf("user has no email")
	}

	from := mail.NewEmail(FromName, m.FromMail)
	subject := "Daily Insight(s)"
	userName := fmt.Sprintf("%v %v", u.FirstName, u.LastName)

	to := mail.NewEmail(userName, u.Email)

	html := BuildInsightsMailTemplate(u, insights)

	message := mail.NewSingleEmail(from, subject, to, "", html)
	_, err := m.Clientele.Send(message)
	if err != nil {
		return err
	}
	return nil
}

func BuildInsightsMailTemplate(u *User, ins []*DailyInsight) string {
	templ, err := template.ParseFiles("daily.templ")
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
