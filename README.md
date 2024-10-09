
A Go based service for parsing Kindle highlights and sending daily insights by email using SendGrid. This project is inspired by Readwise to help users manage and revisit their Kindle highlights simply

![](/img/pensieve.png)

### Getting started

Follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/HelixY2J/pensieve.git
```

```bash
cd pensieve
```

Install the necessary dependencies

```bash
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
go get github.com/sendgrid/sendgrid-go
```

Set up your MySQL database and SendGrid account to obtain an API key


```bash
go run main.go
```
