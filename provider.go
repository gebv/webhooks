package main

import (
	"errors"
	"io"
	"log"
	"os/exec"
)

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-RepositoryEvents
type buildProvider func() Provider

var registredProvider = make(map[string]buildProvider)

func RegProvider(name string, builder buildProvider) error {
	if _, exist := registredProvider[name]; exist {

		return errors.New("provider existing, name='" + name + "'")
	}

	registredProvider[name] = builder

	return nil
}

func Execute(token string, data io.Reader) error {
	config := Cfg.Hooks.ByToken(token)

	if config == nil {
		return ErrNotFound
	}

	if !config.IsEnabled {
		return ErrNotAllowed
	}

	builder, exist := registredProvider[config.ProviderName]

	if !exist {
		log.Printf("not registred provider='%s', token='%s'", config.ProviderName, token)
		return ErrNotSupported
	}

	provider := builder()

	if err := provider.FromPayload(data); err != nil {
		log.Printf("decoded payload, err='%s'", err)

		return ErrNotValid
	}

	if provider.Hook(config) {
		log.Printf("execute command='%s'", config.ShellCommand)

		out, err := exec.Command(config.ShellCommand).Output()
		if err != nil {
			log.Printf("%s", err)
		}
		log.Printf("%s", out)

	}

	return nil
}

type Provider interface {
	Name() string
	FromPayload(interface{}) error
	Hook(*HookConfig) bool
}
