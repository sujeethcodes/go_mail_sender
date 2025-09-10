# 📧 Golang Mail Sender

This project allows sending emails using Golang and viewing email logs

---

## **1️⃣ Setup & Start the Server**

1. Clone the repository:

```bash
git clone https://github.com/sujeethcodes/go_mail_sender.git
cd project folder

go mod tidy

go run cmd/main.go
POST http://localhost:7000/send
Content-Type: application/json

{
  "to": "recipient@example.com",
  "subject": "Test Email",
  "body": "Hello! This is a test email."
}
Response
{
    "data": {
        "status": "sent_attempted",
        "delivered": true,
        "error": null
    }
}
GET http://localhost:9000/logs/mysql
{
    "code": 202,
    "message": "mail fetch Successfully",
    "data": [
        {
            "id": 1,
            "to_email": "recipient@example.com",
            "subject": "Hello",
            "body": "This is a test email",
            "status": "sent_attempted",
            "delivered": true,
            "error_text": null,
            "created_at": "2025-09-09 10:28:39"
        },
        {
            "id": 2,
            "to_email": "recipient@example.com",
            "subject": "Hello",
            "body": "This is a test email",
            "status": "sent_attempted",
            "delivered": true,
            "error_text": null,
            "created_at": "2025-09-09 10:33:40"
        },
        ]
