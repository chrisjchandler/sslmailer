# SSLMailer - Because Certificates Have Bad Timing

SSL certificates have a knack for expiring at the most inconvenient times, leaving us feeling like the world's unluckiest sysadmins. But fret not! SSLMailer is here to help you stay ahead of certificate chaos.

## What Does It Do?

SSLMailer is your trusty sidekick for handling SSL/TLS certificate matters with style and punctuality:

1. **Certificate Expiry Checker:** It diligently examines the SSL/TLS certificates on your system and identifies those with less than 30 days left to live.

2. **Email Notifications:** When a certificate is about to expire, SSLMailer crafts and sends an email notification to let you know. No more surprises!

## Getting Started

Getting SSLMailer up and running is as smooth as it gets:

1. **Clone This Repo:**

   ```bash
   git clone https://github.com/yourusername/sslmailer.git
   cd sslmailer
Configuration:

Edit the sslmailer.go file to set the certDirectory variable to the path where your certificates reside. SSLMailer won't rest until it finds them!
Fill in your email server details (smtpServer, smtpPort, smtpUsername, smtpPassword) so SSLMailer can whisk notifications straight to your inbox.

Build it
go build sslmailer.go

Run it 
./sslmailer

Relax
Sip your coffee, walk your dog (or cat) and let SSLMailer handle those grumpy certificates.

Notes and Considerations
SSLMailer operates as a scheduled service, performing its checks and email notifications every 24 hours. Certificates don't wait for anyone, and neither does SSLMailer!

While SSLMailer is grumpy about certificates, it may need a bit of customization if your certificates are in unconventional places or formats.

Protect your email and server details like treasure! We wouldn't want anyone else receiving your certificate notifications.

Contributions
Contributions are welcome!

License 
This project is released under the "SSL Certificates Savior License." Feel free to use it, tweak it, or ignore it – your certificates, your rules.
