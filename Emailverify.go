package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/badoux/checkmail"
)

func main() {
	email := "samdam456@doceree.com"

	//  Check email syntax
	err := checkmail.ValidateFormat(email)
	if err != nil {
		log.Fatalf("Invalid email syntax: %s", err)
	}

	//  Extract the domain
	parts := strings.Split(email, "@")
	domain := parts[1]

	// Check domain MX (Mail Exchange) records
	_, mxErr := net.LookupMX(domain)
	if mxErr != nil {
		log.Fatalf("MX records lookup failed for domain %s: %s", domain, mxErr)
	}

	//  Try to send a test email (optional)
	// err = sendTestEmail(email)
	// if err != nil {
	// 	log.Fatalf("Failed to send test email: %s", err)
	// }

	fmt.Printf("Email %s is valid and can receive messages.\n", email)
}

// func sendTestEmail(email string) error {
// 	from := "samdam456@doceree.com"
// 	password := "S@niket567"

// 	smtpServer := "smtp.gmail.com"
// 	smtpPort := "587"

// 	// Connect to the SMTP server
// 	auth := smtp.PlainAuth("", from, password, smtpServer)

// 	// Set up the email message
// 	msg := "Subject: Email Verification\r\n\r\nThis is a test email to verify the email address."

// 	// Send the email
// 	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{email}, []byte(msg))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
