package bot

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func (w *Wrapper) textHandler(c telebot.Context) error {
	ctx := context.TODO()

	txt := c.Text()

	if string([]rune(txt)[:4]) == w.config.Name {
		message := strings.ReplaceAll(txt, w.config.Name, "")

		id := c.Chat().ID
		idStr := strconv.Itoa(int(id))
		userName := c.Sender().Username

		messageName := userName + " написал:" + message

		err := w.redisMsg.AddMessage(ctx, idStr, messageName)
		if err != nil {
			slog.Error("Redis AddMessage after response error", err)
			return err
		}

		chat, err := w.redisMsg.GetChat(ctx, idStr)
		if err != nil {
			slog.Error("Yandex GetChat error", err)
			return err
		}

		answer, err := w.openaiSrv.ChatCompetition(ctx, chat)
		if err != nil {
			slog.Error("no massage", err)
			return c.Send("Луна не полная, не могу овтетить на ваш вопрос")

		}

		slog.Info("Выполнен зов", id, userName)
		return c.Send(answer)
	}
	return nil
}

func (w *Wrapper) startHandler(c telebot.Context) error {

	slog.Info("запрос на правильный зов", c.Sender().ID, c.Sender().Username)
	return c.Send(`Напиши "` + w.config.Name + `, <текст сообщения>", чтобы выпонить зов
Напиши "` + deleteCall + `", чтобы забыть нашу переписку`)
}

func (w *Wrapper) deleteHandler(c telebot.Context) error {
	ctx := context.TODO()

	id := c.Chat().ID
	idStr := strconv.Itoa(int(id))

	err := w.redisMsg.Delete(ctx, idStr)
	if err != nil {
		slog.Error(fmt.Sprintf("Redis Delete error: %v", err))
		return c.Send("У меня не получается забыть таких прекрасных людей")
	}

	return c.Send("Волк стал одиноким в этом обсуждении")
}
