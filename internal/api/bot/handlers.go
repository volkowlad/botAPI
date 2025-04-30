package bot

import (
	"context"
	"fmt"
	"gopkg.in/telebot.v4"
	"log/slog"
	"strconv"
	"strings"
)

func (w *Wrapper) textHandler(c telebot.Context) error {
	id := c.Chat().ID
	mtx := getChatMutex(id)
	mtx.Lock()
	defer mtx.Unlock()

	ctx := context.TODO()

	txt := c.Text()

	if string([]rune(txt)[:4]) == w.config.Name {
		message := strings.ReplaceAll(txt, w.config.Name, "")

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

func (w *Wrapper) helloHandler(c telebot.Context) error {
	id := c.Chat().ID
	slog.Info(fmt.Sprintf("add to new chat %v", id))

	return c.Send(fmt.Sprintf(`Я личный бот-помощник - Волк
Покажите как вы умеете выть, чтобы попасть ко мне в стаю
Напишите %s, чтобы ваш клич был услышан
А также не забудьте сделать меня админом, чтобы я мог отправять сообщения😉`, codeCall))
}
