package main

import (
	"crypto/x509"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/sevlyar/go-daemon"
)

const (
	certDirectory = "/etc/ssl/certs/"  // Ugh, the stupid certificate directory
	emailFrom     = "your_email@example.com"
	emailTo       = "recipient@example.com"
	smtpServer    = "smtp.example.com"
	smtpPort      = "587"
	smtpUsername  = "your_username"
	smtpPassword  = "your_password"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "ssl_checker.pid",
		PidFilePerm: 0644,
		LogFileName: "ssl_checker.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		fmt.Printf("Ugh, can't daemonize: %v\n", err)
		return
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	// Main loop: Let's get this certificate checking show on the road
	for {
		checkCertificates()
		time.Sleep(24 * time.Hour)
	}
}

func checkCertificates() {
	// Get the current time
	currentTime := time.Now()

	// Iterate through the certificate files in the directory
	files, err := os.ReadDir(certDirectory)
	if err != nil {
		fmt.Printf("Error reading the darn certificate directory: %v\n", err)
		return
	}

	for _, file := range files {
		certPath := certDirectory + file.Name()

		// Load the certificate
		certData, err := os.ReadFile(certPath)
		if err != nil {
			fmt.Printf("Can't even read the stupid certificate file: %v\n", err)
			continue
		}

		cert, err := x509.ParseCertificate(certData)
		if err != nil {
			fmt.Printf("Error parsing this certificate: %v\n", err)
			continue
		}

		// Calculate days until expiry
		daysUntilExpiry := int(cert.NotAfter.Sub(currentTime).Hours() / 24)

		// Notify if certificate expires in 30 days or less
		if daysUntilExpiry <= 30 {
			subject := fmt.Sprintf("SSL Certificate Expiry Warning for %s", file.Name())
			message := fmt.Sprintf("The %@ certificate %@ expires in %d days.", file.Name(), daysUntilExpiry)
			sendEmail(subject, message)
		}
	}
}

func sendEmail(subject, message string) {
	// Set up email configuration
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	to := []string{emailTo}
	msg := []byte("Subject: " + subject + "\r\n\r\n" + message)

	// Send email
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, emailFrom, to, msg)
	if err != nil {
		fmt.Printf("Error sending this blasted email: %v\n", err)
	} else {
		fmt.Println("Email sent, finally.")
	}
}
