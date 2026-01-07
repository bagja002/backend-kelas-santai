package utils

import (
	"fmt"
	"log"
	"project-kelas-santai/internal/config"

	"gopkg.in/gomail.v2"
)

const (
	SMTPHost   = "smtp.gmail.com"
	SMTPPort   = 587
	SenderName = "Kelas Santai <kelasantai.bootcamp@gmail.com>"
	AuthEmail  = "kelasantai.bootcamp@gmail.com"
)

// SendOrderSuccessEmail mengirim email notifikasi pesanan berhasil
func SendOrderSuccessEmail(toEmail, userName, message string) error {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	subject := "Selamat Datang di Kelas Santai"
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Selamat Datang di Kelas Santai</title>
			<style>
				body {
					font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
					line-height: 1.6;
					color: #4E342E; /* Mocha Text */
					background-color: #EFEBE9; /* Light Mocha Background */
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 40px auto;
					background-color: #ffffff;
					border-radius: 12px;
					overflow: hidden;
					box-shadow: 0 4px 6px rgba(0,0,0,0.05); /* Soft shadow */
				}
				.header {
					background-color: #795548; /* Matte Mocha Primary */
					padding: 40px 20px;
					text-align: center;
				}
				.header h1 {
					color: #ffffff;
					margin: 0;
					font-size: 24px;
					font-weight: 600;
					letter-spacing: 0.5px;
				}
				.content {
					padding: 40px 30px;
					background-color: #ffffff;
				}
				.greeting {
					font-size: 20px;
					font-weight: 600;
					color: #5D4037;
					margin-bottom: 20px;
				}
				.message-box {
					background-color: #F5F5DC; /* Cream/Beige Accent */
					border-left: 4px solid #8D6E63;
					padding: 15px;
					margin: 20px 0;
					border-radius: 4px;
				}
				.footer {
					background-color: #D7CCC8; /* Lighter Mocha footer */
					padding: 20px;
					text-align: center;
					font-size: 13px;
					color: #5D4037;
				}
				.button {
					display: inline-block;
					padding: 12px 24px;
					background-color: #6D4C41;
					color: #ffffff;
					text-decoration: none;
					border-radius: 6px;
					font-weight: bold;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Kelas Santai</h1>
				</div>
				<div class="content">
					<div class="greeting">Halo, %s!</div>
					<p>Terima kasih telah mendaftar di <strong>Kelas Santai</strong>.</p>
					<p>Kami sangat senang menyambut Anda sebagai bagian dari komunitas kami. Berikut adalah detail pendaftaran Anda:</p>
					
					<div class="message-box">
						<p style="margin: 0;">%s</p>
					</div>

					<p>Selamat belajar dan semoga pengalaman Anda menyenangkan!</p>
					
					<div style="text-align: center;">
						<a href="#" class="button" style="color: #ffffff;">Mulai Belajar</a>
					</div>
				</div>
				<div class="footer">
					<p>&copy; 2024 Kelas Santai. All rights reserved.</p>
					<p>Jakarta</p>
				</div>
			</div>
		</body>
		</html>
	`, userName, message)

	m := gomail.NewMessage()
	m.SetHeader("From", SenderName)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	fmt.Println("Sending email to:", toEmail)
	// Body too long to log fully usually, keeping it clean
	fmt.Println("Email subject:", subject)

	d := gomail.NewDialer(SMTPHost, SMTPPort, AuthEmail, cfg.Web.AppPassword)

	return d.DialAndSend(m)
}
