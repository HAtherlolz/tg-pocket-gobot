package telegram

import (
	"context"
	"log"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	pocket "github.com/zhashkevych/go-pocket-sdk"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					b.handleError(update.Message.Chat.ID, err)
				}
				continue
			}

			if err := b.handleMessage(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
		}
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case "start":
		return b.handleStartCommand(message)
	default:
		return b.handleUknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidURL
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauth
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		return errUnableToSave
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.SavedSuccessfully)
	msg.ReplyToMessageID = message.MessageID
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.AlreadyAuth)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.InvalidCommand)
	_, err := b.bot.Send(msg)
	return err
}
