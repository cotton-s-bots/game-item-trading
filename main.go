package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tradingBot/commands"

	// Import DisGo packages
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	h := handler.New()

	token := os.Getenv("DISCORD_BOT_TOKEN")
	guildID := snowflake.GetEnv("DISCORD_GUILD_ID")

	// Create a new client instance
	client, err := disgo.New(token,
		// Set gateway configuration options
		bot.WithEventListeners(h),
		bot.WithGatewayConfigOpts(
			// Set enabled intents
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
		// Listen to the Ready event in order to know when the bot is connected
		bot.WithEventListenerFunc(func(e *events.Ready) {
			fmt.Println("Bot is connected as", e.User.Username)
		}),
	)
	if err != nil {
		panic(err)
	}

	h.SlashCommand("/ping", commands.PingCommandHandler)
	if err := handler.SyncCommands(client, commands.Commands, []snowflake.ID{guildID}); err != nil {
		panic(err)
	}

	// Ensure we close the client on exit
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client.Close(ctx)
	}()

	// Connect to the gateway
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.OpenGateway(ctx); err != nil {
		panic(err)
	}

	// Wait here until CTRL+C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
	fmt.Println("Shutting down bot...")
}
