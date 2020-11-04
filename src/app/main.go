package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheBiggestBombs/ffxivtoolkit"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var client ffxivtoolkit.Client

func main() {
	loadConfig()

	// Get some details from the config
	ffxivToolkitToken := viper.GetString("FFXIVToolkit.Token")
	defaultLodestoneID := viper.GetString("FFXIVToolkit.FCLodestoneID")

	// get fc details
	client = ffxivtoolkit.New(ffxivToolkitToken)
	logDetails(defaultLodestoneID)
	logMembers(defaultLodestoneID)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func loadConfig() {
	viper.SetDefault("Debug", false)
	viper.SetDefault("FFXIVToolkit.FCLodestoneID", "")
	viper.SetDefault("FFXIVToolkit.Token", "")
	viper.SetDefault("Discord.BotToken", "")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.raidbot")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func logDetails(lodestoneID string) {
	freeCompany := client.FreeCompany.Basic(lodestoneID)

	log.Println("Lodestone ID: " + freeCompany.LodestoneID)
	log.Println("FC Name: " + freeCompany.Name)
}

func logMembers(lodestoneID string) {
	members := client.FreeCompany.Members(lodestoneID)

	log.Println("Total Members: " + fmt.Sprint(len(members)))
}

func openDiscord() {
	ffxivToolkitToken := viper.GetString("Discord.BotToken")

	discord, err := discordgo.New("Bot " + ffxivToolkitToken)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Raidbot is now running. -Press CTRL_C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
