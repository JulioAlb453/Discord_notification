package main

import (
	service "discordNotification/internal/Application/services"
	"discordNotification/internal/infraestructure/discord"
	"discordNotification/internal/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func main() {
    
    webhookURLs := map[string]string{
        "desarrollo": "https://discord.com/api/webhooks/1349920173519409183/CoI6bD2KUbmIvWpWwCForLEnJi83l9PPZZlCGelGSoroPs_DFMJFWMV7XRrq3Ug7F4B3",
        "pruebas":    "https://discord.com/api/webhooks/1349920164317237278/wmDuYSAERo0ev-e-P_88TRJf8j5XagOD0gELlL9yK5OpU22dBlQra8GChQSNBTkcExiY",
        "general":    "https://discord.com/api/webhooks/1349964247911763980/YwVtuEJgxvWKcnttifzZI59CVVqPuEyT_B_Y0_OyDP__ewHw_wNfUcG-2OSej-L01xji",
    }

    
    discordAdapter := discord.NewDiscordAdapter(webhookURLs)

    
    eventService := service.NewEventService(discordAdapter)

    
	ginHandler := http.NewGinHandler(eventService)
    
    r := gin.Default()
    r.POST("/webhook", ginHandler.HandleGitHubEvent)

   
    if err := r.Run(":4040"); err != nil {
        panic("No se pudo iniciar el servidor: " + err.Error())
    }
}