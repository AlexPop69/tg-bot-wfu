package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/AlexPop69/tg-bot-wfu/internal/service"
	"github.com/sirupsen/logrus"
)

const (
	// apiURL - базовый путь (префикс с которого начинаются все запросы)
	apiURL = "https://api.telegram.org/bot"

	adminPrefix = "/admin"

	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

// Bot представляет собой структуру, содержащую основные компоненты для работы с Telegram API
type Bot struct {
	token    string
	apiURL   string
	services *service.Service
	client   *http.Client
}

// NewBot создает новый экземпляр бота с необходимыми сервисами
func NewBot(services *service.Service) (*Bot, error) {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("failed get token")
	}

	return &Bot{
			token:    token,
			apiURL:   apiURL + token,
			services: services,
			client:   &http.Client{Timeout: 10 * time.Second},
		},
		nil
}

// Start запускает бота и начинает обработку обновлений
func (b *Bot) Start() error {
	offset := 0

	for {
		// Получение обновлений от Telegram API
		updates, err := b.getUpdates(offset)
		if err != nil {
			logrus.Printf("Error getting updates: %v\n", err)
			continue
		}

		// Обработка каждого обновления
		for _, update := range updates {
			if update.Message != nil {
				b.handleUpdate(update.Message)
				offset = update.UpdateID + 1 // Увеличиваем offset, чтобы получать только новые обновления
			}
		}
	}
}

// getUpdates выполняет запрос к Telegram API для получения обновлений
func (b *Bot) getUpdates(offset int) ([]Update, error) {
	// Подготовка запроса
	url := fmt.Sprintf("%s/%s?offset=%d&timeout=%v",
		b.apiURL, getUpdatesMethod, offset, b.client.Timeout)

	resp, err := b.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Чтение и декодирование ответа
	var updateResp UpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		return nil, err
	}

	if !updateResp.OK {
		return nil, fmt.Errorf("failed to get updates from Telegram API")
	}

	return updateResp.Result, nil
}

// sendMessage отправляет сообщение в Telegram чат.
func (b *Bot) sendMessage(chatID int64, text string) error {
	// Подготовка сообщения для отправки
	data := url.Values{}
	data.Set("chat_id", fmt.Sprintf("%d", chatID))
	data.Set("text", text)

	resp, err := b.client.PostForm(fmt.Sprintf("%s/%s", b.apiURL, sendMessageMethod), data)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer resp.Body.Close()

	// Проверка успешности запроса
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}

// handleUpdate обрабатывает обновления, полученные от Telegram API
func (b *Bot) handleUpdate(message *Message) {
	logrus.Infof(`Received message "%s" from user: %s`, message.Text, message.From.Username)

	if strings.HasPrefix(message.Text, adminPrefix) {
		if !b.isAdmin(message.From.Username) {
			b.sendMessage(message.Chat.ID, "У вас нет прав для выполнения этой команды.")

			return
		}

		// обработчик команд от админа

	} else {
		// обработчик команд от остальных пользователей
	}
}

// Проверка является ли пользователь администратором
func (b *Bot) isAdmin(username string) bool {
	admin, err := b.services.Admin.Authentication(username)
	if err != nil {
		logrus.Errorf("Error checking admin status for user %s: %v", username, err)
		return false
	}

	if admin == nil {
		logrus.Warnf("User %s is not an admin", username)
		return false
	}

	return true
}
