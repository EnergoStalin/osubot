package handlers

func RegisterCommand(prefix string, callback MessageCallback) {
	messageHandlerMap[prefix] = callback
}
