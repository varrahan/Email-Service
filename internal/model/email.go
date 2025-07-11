package model

type Email struct {
    Name    string `json:"name"`    // Sender's name
    Email   string `json:"email"`   // Sender's email address
    Subject string `json:"subject"`
    Message string `json:"message"`
}