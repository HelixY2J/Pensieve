
A Go based service for parsing Kindle highlights and sending daily insights by email using SMPT2GO. This project is inspired by Readwise to help users manage and revisit their Kindle highlights simply

![](/img/insights.png)

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
go mody tidy
```

Set up your MySQL database and SMPT2GO account to obtain an API key

Make sure to have GO 1.12+ installed

```bash
make run
```
Also check setup the env variables as listed in .envrc.example file