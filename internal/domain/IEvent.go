
package domain

type DiscordNotifier interface {
    Send(channel, message string) error
}