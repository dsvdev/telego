package common

type ClientMethod string

const (
	GetMe       ClientMethod = "getMe"
	GetUpdates  ClientMethod = "getUpdates"
	SendMessage ClientMethod = "sendMessage"
)
