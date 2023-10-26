package utils

import "fmt"

type Notification interface {
	RenderNotification(s string) string
}

func CreateNotification(message string, renderer Notification) string {
	return renderer.RenderNotification(message)
}

type SuccessNotification struct{}

func (r SuccessNotification) RenderNotification(s string) string {
	return fmt.Sprintf("<div class='notification success show' hx-on:click='this.remove()'>%s</div>", s)
}

type ErrorNotification struct{}

func (r ErrorNotification) RenderNotification(s string) string {
	return fmt.Sprintf("<div class='notification error show' hx-on:click='this.remove()'>%s</div>", s)
}
