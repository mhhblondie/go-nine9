package helper

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var SecretKey = os.Getenv("JWT_SECRET")

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Vérifiez que le token est signé avec l'algorithme attendu
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func GetToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	// Split the header into two parts: "Bearer" and the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		// Handle error: invalid or missing Authorization header
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or missing Authorization header",
		})
	}

	token := parts[1]

	return token, nil
}

func GetClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	claims, ok := c.Locals("userClaims").(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Unauthorized: Claims missing or not jwt.MapClaims")
	}
	return claims, nil
}

func GeneratePassword(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+{}[]|;':\",.<>/?`~")

	password := make([]rune, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	for _, c := range password {
		if !unicode.IsUpper(c) && !unicode.IsLower(c) && !unicode.IsDigit(c) && c < '!' || c > '~' {
			return GeneratePassword(length)
		}
	}

	return string(password)
}

func isManager(user_role string) bool {
	if user_role == "manager" {
		return true
	}
	return false
}

func isOwner(user_role string) bool {
	if user_role == "owner" {
		return true
	}
	return false
}

var jwtKey = os.Getenv("JWT_SECRET")

func GenerateToken(id uuid.UUID, role string, firstname string, email string, salonID uuid.UUID) (string, error) {
	NewClaim := jwt.MapClaims{
		"id":        id.String(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 1 jour
		"role":      role,
		"firstname": firstname,
		"email":     email,
		"salonID":   salonID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaim)

	signedToken, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func SendConfirmationEmail(msg string, recipent string, MailSubject string) {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")

	to := []string{
		recipent,
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	from_msg := fmt.Sprintf("From: %s\r\n", from)
	to_msg := fmt.Sprintf("To: %s\r\n", recipent)
	subject := fmt.Sprintf("Subject: %s\r\n", MailSubject)
	body := msg

	message := []byte(from_msg + to_msg + subject + "\r\n" + body)

	fmt.Println("message" + msg)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}

func CreateMailBody(recipientName string, email string,password string, salon string) string {
	return fmt.Sprintf(`Bonjour %s,
	votre employeur vous a créé un compte sur notre application de réservation de coiffeur.
	Voici vos identifiants de connexion:
	Email: %s
	Mot de passe: %s
	Connectez-vous à l'application pour gérer vos réservations et votre emploi du temps.
	Cordialement,
	Votre équipe %s`, recipientName, email, password, salon)
}



func CreateConfirmationEmailBody(recipientName string, date string, person string, salon string) string {
	return fmt.Sprintf(`Bonjour %s,

Nous sommes heureux de vous confirmer que votre réservation a été effectuée avec succès.

Détails de la réservation :
Date : %s
Avec votre coiffeur : %s

Nous vous remercions de votre réservation et nous avons hâte de vous accueillir.

Cordialement,
Votre équipe %s`, recipientName, date, person, salon)
}

func CreateDeleteStaffBody(recipientName string, date string, salon string, tel string) string {
	return fmt.Sprintf(`Bonjour %s,

Nous sommes désolé de vous indiquer que votre compte a été supprimé le %s

Vous ne pourrez plus recevoir de réservation du salon : %s

Merci de contacter votre ancien manager si soucis il y au :%s

Cordialement,
Planity`, recipientName, date, salon, tel)
}

func CreateNewStaffBody(recipientName string, date string, salon string, password string) string {
	return fmt.Sprintf(`Bonjour %s,

Votre compte a bien été crée le %s

Pour le salon : %s

Utilisez votre mail et ce mot de passe donné par le manager ou l'admin :  %s

Cordialement,
Planity`, recipientName, date, salon, password)
}
