package handlers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ContactForm struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Message string `json:"message" binding:"required"`
}

// SendContactMessage обработчик для отправки сообщения с формы
func (app *App) SendContactMessage(c *gin.Context) {
	var form ContactForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Неверные данные формы: " + err.Error(),
		})
		return
	}

	// Валидация обязательных полей
	if strings.TrimSpace(form.Name) == "" || strings.TrimSpace(form.Email) == "" || strings.TrimSpace(form.Message) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Все обязательные поля должны быть заполнены",
		})
		return
	}

	// Отправка email
	if err := sendEmail(form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Ошибка при отправке сообщения: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Сообщение успешно отправлено! Мы свяжемся с вами в ближайшее время.",
	})
}

// sendEmail отправляет email через SMTP
func sendEmail(form ContactForm) error {
	// Получение SMTP настроек из переменных окружения
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	toEmail := os.Getenv("TO_EMAIL")

	// Если TO_EMAIL не установлен, используем SMTP_USER
	if toEmail == "" {
		toEmail = smtpUser
	}

	// Проверка наличия обязательных переменных окружения
	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" {
		return fmt.Errorf("SMTP configuration is incomplete")
	}

	// Аутентификация SMTP
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Формирование email сообщения
	subject := "Новое сообщение с формы обратной связи AuraPrint"
	body := createEmailBody(form)

	// Формирование MIME сообщения
	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		body + "\r\n")

	// Отправка email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{toEmail}, msg)
	if err != nil {
		return err
	}

	return nil
}

// createEmailBody создает HTML тело email
func createEmailBody(form ContactForm) string {
	phone := form.Phone
	if phone == "" {
		phone = "Не указан"
	}

	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; color: #333; }
        .header { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; color: white; text-align: center; }
        .content { background: #f9f9f9; padding: 20px; border-radius: 10px; margin: 20px 0; }
        .field { margin-bottom: 15px; }
        .field strong { color: #2c3e50; }
        .message { background: white; padding: 15px; border-radius: 5px; border-left: 4px solid #667eea; margin-top: 10px; }
        .footer { color: #666; font-size: 12px; text-align: center; margin-top: 20px; }
    </style>
</head>
<body>
    <div class="header">
        <h1>Новое сообщение с AuraPrint</h1>
    </div>
    
    <div class="content">
        <div class="field">
            <strong>👤 Имя:</strong> ` + form.Name + `
        </div>
        
        <div class="field">
            <strong>📧 Email:</strong> ` + form.Email + `
        </div>
        
        <div class="field">
            <strong>📞 Телефон:</strong> ` + phone + `
        </div>
        
        <div class="field">
            <strong>💬 Сообщение:</strong>
            <div class="message">` + strings.ReplaceAll(form.Message, "\n", "<br>") + `</div>
        </div>
    </div>
    
    <div class="footer">
        <p>Сообщение отправлено с формы обратной связи сайта AuraPrint</p>
        <p>© ` + time.Now().Format("2006") + ` AuraPrint. Все права защищены.</p>
    </div>
</body>
</html>`
}
