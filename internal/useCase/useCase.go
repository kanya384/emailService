package useCase

import (
	"bytes"
	"context"
	"emailservice/internal/domain/delivery"
	"emailservice/internal/domain/subscriber"
	"emailservice/internal/domain/template"
	"emailservice/internal/domain/template/path"
	repo "emailservice/internal/repository/postgres"
	"emailservice/internal/useCase/adapters/repository"
	"emailservice/pkg/logger"
	"emailservice/pkg/mail"
	"fmt"
	"os"
	"time"

	tmp "html/template"

	"github.com/google/uuid"
)

type useCase struct {
	deliveryRepo    repository.Delivery
	templatesRepo   repository.Template
	subscribersRepo repository.Subscriber
	logger          logger.Interface
	sender          mail.Sender
	storePath       string
	from            string
	host            string
}

const (
	callBackImageUrl = "http://%s/delivery/%s/%s"
)

func New(repo repo.Repository, logger *logger.Logger, sender mail.Sender, storePath, from, host string) *useCase {
	os.MkdirAll(storePath, os.ModePerm)
	return &useCase{
		deliveryRepo:    repo.Delivery,
		templatesRepo:   repo.Template,
		subscribersRepo: repo.Subscriber,
		logger:          logger,
		storePath:       storePath,
		sender:          sender,
		from:            from,
		host:            host,
	}
}

func (uc *useCase) CreateTemplate(ctx context.Context, templateContent []byte) (result *template.Template, err error) {
	fullPath := uc.storePath + "/" + uuid.NewString() + ".html"
	f, err := os.Create(fullPath)
	if err != nil {
		uc.logger.Error("error creating template file: %s", err.Error())
		return
	}
	defer f.Close()

	_, err = f.Write(templateContent)
	if err != nil {
		uc.logger.Error("error storing template file: %s", err.Error())
		return
	}

	filePath, err := path.NewPath(fullPath)
	if err != nil {
		uc.logger.Error("template file path invalid: %s", err.Error())
		return
	}
	result, err = template.New(*filePath)
	if err != nil {
		uc.logger.Error("create template error: %s", err.Error())
		return
	}

	err = uc.templatesRepo.CreateTemplate(ctx, result)
	if err != nil {
		uc.logger.Error("error saving template to repo: %s", err.Error())
		return
	}
	return
}

func (uc *useCase) CreateDeliveryWithSubscribers(ctx context.Context, delivery *delivery.Delivery, subscribers ...*subscriber.Subscriber) (err error) {
	return uc.deliveryRepo.CreateDeliveryWithSubscribers(ctx, delivery, subscribers...)
}

func (uc *useCase) MarkAsReadedBySubscriber(ctx context.Context, deliveryID uuid.UUID, subscriberID uuid.UUID) (err error) {
	return uc.deliveryRepo.MarkAsReadedBySubscriber(ctx, deliveryID, subscriberID)
}

func (uc *useCase) ProcessDelivery(ctx context.Context, done <-chan struct{}, interval time.Duration) {
L:
	for {
		select {
		case <-done:
			break L
		default:
			err := uc.doProcessDelivery(ctx)
			if err != nil {
				uc.logger.Error("processing delivery error: %s", err.Error())
			}
			time.Sleep(interval)
		}

	}
}

func (uc *useCase) doProcessDelivery(ctx context.Context) (err error) {

	list, err := uc.deliveryRepo.ReadDeliveryTasks(ctx)
	if err != nil {
		return
	}

L:
	for _, del := range list {
		subscribers, err := uc.deliveryRepo.ReadDeliverySubscribers(ctx, del.ID())
		if err != nil {
			uc.logger.Error("get delivery subscribers error: %s", err.Error())
			continue
		}
		template, err := uc.templatesRepo.ReadTemplateById(ctx, del.TemplateID())
		if err != nil {
			uc.logger.Error("get template error: %s", err.Error())
			continue
		}
		var t *tmp.Template
		t, err = t.ParseFiles(template.Path().String())
		if err != nil {
			uc.logger.Error("parse template error: %s", err.Error())
			continue
		}

		buff := new(bytes.Buffer)

		for _, subscriber := range subscribers {
			templateData := NewTemplateStruct(subscriber.Name().String(), subscriber.Surname().String(), int(subscriber.Age()), fmt.Sprintf(callBackImageUrl, uc.host, del.ID().String(), subscriber.ID().String()))
			t.Execute(buff, templateData)
			m := mail.NewMessage()
			m.SetTo(subscriber.Email().String())
			m.SetFrom(uc.from)
			m.SetContentType("text/html")
			m.SetBody(buff.String())

			err = uc.sender.SendEmail(m)
			if err != nil {
				uc.logger.Error("send email error: %s", err.Error())
				continue L
			}
		}

		updateFn := func(oldDelivery *delivery.Delivery) (*delivery.Delivery, error) {
			return delivery.NewWithID(oldDelivery.ID(), oldDelivery.CreatedAt(), time.Now(), oldDelivery.TemplateID(), oldDelivery.SendAt(), true)
		}
		_, err = uc.deliveryRepo.UpdateDelivery(ctx, del.ID(), updateFn)
		if err != nil {
			uc.logger.Error("update delivery error: %s", err.Error())
			return err
		}
	}
	return
}
