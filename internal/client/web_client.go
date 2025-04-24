package client

import (
	"bytes"
	"com.github/dsvdev/telego/internal/client/models/requests"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"com.github/dsvdev/telego/internal/client/models/common"
	"com.github/dsvdev/telego/internal/client/models/responses"
)

const (
	apiUrl = "https://api.telegram.org/bot"
)

func resolveUrl(token string, method common.ClientMethod) string {
	return apiUrl + token + "/" + string(method)
}

// doTelegramRequest выполняет HTTP-запрос к Telegram Bot API и парсит ответ в BaseResponse[T].
// method: "GET", "POST" и т.д.
// body: nil для GET-запросов, иначе io.Reader с JSON-данными.
func doTelegramRequest[T any](url, method string, body io.Reader) (*T, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(data))
	}

	var parsed responses.BaseResponse[T]
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if !parsed.IsSuccess {
		return nil, errors.New("telegram API returned not ok")
	}
	if parsed.Result == nil {
		return nil, errors.New("telegram API returned nil result")
	}

	return parsed.Result, nil
}

func GetMe(token string) (*responses.User, error) {
	requestUrl := resolveUrl(token, common.GetMe)
	return doTelegramRequest[responses.User](requestUrl, http.MethodGet, nil)
}

func GetUpdates(token string, offset int64) (*[]*responses.Update, error) {
	urlObj, err := url.Parse(resolveUrl(token, common.GetUpdates))
	if err != nil {
		return nil, err
	}

	q := urlObj.Query()
	q.Set("offset", strconv.FormatInt(offset, 10))
	q.Set("timeout", "10")
	urlObj.RawQuery = q.Encode()

	return doTelegramRequest[[]*responses.Update](urlObj.String(), http.MethodGet, nil)
}

func SendMessage(token string, req *requests.SendMessageRequest) (*responses.Message, error) {
	requestUrl := resolveUrl(token, common.SendMessage)
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return doTelegramRequest[responses.Message](requestUrl, http.MethodPost, bytes.NewBuffer(jsonBody))
}
