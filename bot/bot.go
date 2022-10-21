package bot

import (
	"fmt"
	"log"
	"newsbot/providers"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

type Bot struct {
	Token		string
	Every		string
	ChannelID	string
	Provider	providers.Provider

	discordGO	*discordgo.Session
	scheduler	*gocron.Scheduler
}

func (b *Bot) getInitName() string {
	return fmt.Sprintf("Bot %s", b.Token)
}

func (b *Bot) scheduledEvent() {
	b.scheduler.Cron(b.Every).Do(func(){
		b.serveContents()
	})
	b.scheduler.StartBlocking()
}

func (b *Bot) serveContents() {
	contents := b.Provider.ProvideContents()

	for _, content := range contents {
		content.LogSent()
		b.discordGO.ChannelMessageSendComplex(b.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				URL: content.Url,
				Title: content.Title,
				Type: discordgo.EmbedTypeArticle,
			},
		})
	}
}

func (b *Bot) Init() {
	dg, err := discordgo.New(b.getInitName())

	if err != nil {
		log.Fatal("Error initializing discordgo")
	}
	b.discordGO = dg
	b.scheduler = gocron.NewScheduler(time.UTC)
}

func (b *Bot) ServeForever() {
	err := b.discordGO.Open()

	if err != nil {
		log.Fatal("Error opening connection with discord")
	}
	defer b.discordGO.Close()
	fmt.Println("Serving news on discord ...")

	b.scheduledEvent()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func (b *Bot) RegisterContentProvider(cp providers.ContentProvider) {
	b.Provider.RegisterContentProvider(cp)
}