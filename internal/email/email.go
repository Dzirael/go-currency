package email

import (
	"log"
	"net/smtp"

	"github.com/Dzirael/go-curenncy/internal/pkg/config"
	models "github.com/Dzirael/go-curenncy/internal/pkg/models/users"
	"github.com/Dzirael/go-curenncy/internal/pkg/persistence"
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		if err := SendToAll(); err != nil {
			log.Printf("notification job failed %s", err)
		}
	})

	c.Start()

}

func SendToAll() error {
	s := persistence.GetUserRepository()

	users, err := s.All()
	if err != nil {
		return err
	}

	for _, user := range *users {
		if err := sendToUser(user); err != nil {
			log.Printf("send email failed %s", err)
			continue
		}
	}
	return nil
}

func sendToUser(user models.User) (err error) {
	conf := config.GetConfig()
	message := []byte("Subject: Test mail\r\n\r\nEmail body\r\n")
	auth := smtp.PlainAuth("", conf.Email.From, conf.Email.Password, conf.Email.Smtp)
	err = smtp.SendMail(conf.Email.Smtp+":587", auth, conf.Email.From, []string{user.Email}, message)
	return
}
