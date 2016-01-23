package main

import (
	"time"
)

// see https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-Push

type BitbucketUser struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	UUID        string `json:"uuid"`
}

type BitbucketRepository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	UUID     string `json:"uuid"`

	Links map[string]interface{}

	Website   string `json:"website"`
	SCM       string `json:"scm"`
	IsPrivate bool   `json:"is_private"`
}

type BitbucketEventMeta struct {
	Author     BitbucketUser       `json:"actor"`
	Repository BitbucketRepository `json:"repository"`
}

type BitbucketPush struct {
	BitbucketEventMeta

	Push BitbucketPushChanges `json:"push"`
}

type BitbucketPushChanges struct {
	Changes []BitbucketPushChange
}

type PushTarget struct {
	Hash   string `json:"hash"`
	Author struct {
		User BitbucketUser `json:"user"`
		Raw  string        `json:"raw"`
	} `json:"author"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
	Type    string    `json:"type"`
}

type PushInforation struct {
	Type       string              `json:"type"`
	Name       string              `json:"name"`
	Repository BitbucketRepository `json:"repository"`
	Target     PushTarget          `json:"target"`
}

type BitbucketPushChange struct {
	// new
	New PushInforation `json:"new"`
	// old
	// created
	// forced
	// closed
	Links     map[string]interface{}
	Commits   BitbucketCommits `json:"commits"`
	Truncated bool             `json:"truncated"`
}

type BitbucketCommits []BitbucketCommit

type BitbucketCommit struct {
	Hash    string
	Type    string
	Message string
	Author  BitbucketUser

	// Links to the change on Bitbucket (html), in the API (commits), and in the form of a diff (diff).
	Links map[string]interface{}
}
