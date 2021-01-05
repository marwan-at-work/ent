// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
	// optional schema name for the Task table.
	TaskSchema string
	// optional schema name for the Team table.
	TeamSchema string
	// optional schema name for the User table.
	UserSchema string
}

// hooks per client, for fast access.
type hooks struct {
	Task []ent.Hook
	Team []ent.Hook
	User []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// WithTaskSchema allows an alternate
// database name to be passed into ent operations.
func WithTaskSchema(schema string) Option {
	return func(c *config) {
		c.TaskSchema = schema
	}
} // WithTeamSchema allows an alternate
// database name to be passed into ent operations.
func WithTeamSchema(schema string) Option {
	return func(c *config) {
		c.TeamSchema = schema
	}
} // WithUserSchema allows an alternate
// database name to be passed into ent operations.
func WithUserSchema(schema string) Option {
	return func(c *config) {
		c.UserSchema = schema
	}
}
