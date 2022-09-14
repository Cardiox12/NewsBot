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
	Every		int
	ChannelID	string
	Provider	providers.Provider

	discordGO	*discordgo.Session
	scheduler	*gocron.Scheduler
}

func (b *Bot) getInitName() string {
	return fmt.Sprintf("Bot %s", b.Token)
}

func (b *Bot) scheduledEvent() {
	b.scheduler.Every(b.Every).Second().Do(func(){
		b.serveContents()
	})
	b.scheduler.StartBlocking()
}

func (b *Bot) serveContents() {
	contents := b.Provider.ProvideContents()

	for _, content := range contents {
		fmt.Println("Sending...")
		b.discordGO.ChannelMessageSend(b.ChannelID, content.String())
	}
}

func (b *Bot) Init() {
	dg, err := discordgo.New(b.getInitName())

	if err != nil {
		log.Fatal("Error initializing discordgo")
	}
	b.discordGO = dg
	b.scheduler = gocron.NewScheduler(time.UTC)
	// b.discordGO.Identify.Intents = discordgo.IntentsGuildMessages
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