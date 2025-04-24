package responses

type BaseResponse[T any] struct {
	IsSuccess bool `json:"ok"`
	Result    *T   `json:"result"`
}

type User struct {
	ID                      int64   `json:"id"` // 64-битный int, чтобы вместить до 52 значащих бит
	IsBot                   bool    `json:"is_bot"`
	FirstName               string  `json:"first_name"`
	LastName                *string `json:"last_name,omitempty"`
	Username                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	IsPremium               bool    `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool    `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool    `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool    `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool    `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    bool    `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           bool    `json:"has_main_web_app,omitempty"`
}

type Update struct {
	ID            int64          `json:"update_id"` // 64-битный int, чтобы вместить до 52 значащих бит
	Message       *Message       `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}

type Message struct {
	Text string `json:"text"`
	Chat *Chat  `json:"chat"`
}

type CallbackQuery struct {
	Data    string   `json:"data"`
	Message *Message `json:"message"`
}

type Chat struct {
	ID int64 `json:"id"`
}
