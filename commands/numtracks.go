/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numtracks.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// NumTracksCommand is a command that outputs the current number of tracks in
// the audio queue.
type NumTracksCommand struct{}

// Aliases is a method that returns the current aliases for the add command.
func (c *NumTracksCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.numtracks")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *NumTracksCommand) IsAdmin() bool {
	return viper.GetBool("permissions.numtracks")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *NumTracksCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}