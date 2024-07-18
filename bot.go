package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func start(Token string, Intents discordgo.Intent) {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Errorf("error creating Discord session: %s", err.Error())
		return
	}

	// Register the message func as a callback for message events.
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = Intents

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Errorf("error opening connection %s\n", err.Error())
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Warnf("#Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

	log.Println("Graceful shutdown")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	parameters := strings.Fields(m.Content)
	if len(parameters) == 0 {
		return
	}

	stack := newStack()
	buff := ""
	invalidBracketArrangement := false

	for _, r := range m.Content {
		if isOpenBracket(r) {
			stack.push(r)
		} else if isClosingBracket(r) {
			if bracketAwaitingClosure, err := stack.peek(); err == nil {
				if matches(bracketAwaitingClosure, r) {
					stack.pop()
				} else {
					invalidBracketArrangement = true
				}
			}
		}
	}

	if stack.IsEmpty() {
		return
	}

	for !stack.IsEmpty() {
		opening, _ := stack.pop()
		switch opening {
		case '(':
			buff += ")"
		case '[':
			buff += "]"
		case '{':
			buff += "}"
		}
	}
	if invalidBracketArrangement {
		buff += "\nInvalid Bracket Arrangement Detected"
	}
	buff += "\n-# Bracket Closing Service"
	s.ChannelMessageSend(m.ChannelID, buff)
}

func isOpenBracket(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

func isClosingBracket(char rune) bool {
	return char == ')' || char == ']' || char == '}'
}

func matches(open, close rune) bool {
	return false ||
		(open == '(' && close == ')') ||
		(open == '[' && close == ']') ||
		(open == '{' && close == '}')
}
