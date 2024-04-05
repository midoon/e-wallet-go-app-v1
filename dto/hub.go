package dto

type Hub struct {
	NotificationChan map[string]chan NotificationData
}
