package main

import (
	"log"
	"regexp"
)

func init() {
	RegProvider("bitbucket:push", func() Provider { return &BitbucketPushProvider{} })
}

type BitbucketPushProvider struct {
	BitbucketPush
}

func (b *BitbucketPushProvider) FromPayload(data interface{}) error {

	return FromJson(b, data)
}

func (b *BitbucketPushProvider) Name() string {

	return "bitbucket:push"
}

func (b *BitbucketPushProvider) Hook(cfg *HookConfig) bool {
	if b.Repository.Name != cfg.Check.RepositoryName {
		log.Printf("not expected branch_name='%s', want='%s'", b.Repository.Name, cfg.Check.RepositoryName)
		return false
	}

	if b.Author.Username != cfg.Check.UserName {
		log.Printf("not expected actor, username='%s', want='%s'", b.Author.Username, cfg.Check.UserName)
		return false
	}

	for _, change := range b.Push.Changes {
		if change.New.Type != "branch" {
			log.Printf("not expected change type='%s', want='%s', hash='%s'", change.New.Type, "branch", change.New.Target.Hash)
			continue
		}

		if change.New.Name != cfg.Check.BranchName {
			log.Printf("not expected branch_name='%s', want='%s', hash='%s'", change.New.Name, cfg.Check.BranchName, change.New.Target.Hash)
			continue
		}

		if change.New.Target.Author.User.Username != cfg.Check.UserName {
			log.Printf("not expected target author, username='%s', want='%s', hash='%s'", change.New.Target.Author.User.Username, cfg.Check.UserName, change.New.Target.Hash)
			continue
		}

		if change.New.Target.Type != "commit" {
			log.Printf("not expected target type='%s', want='%s', hash='%s'", change.New.Target.Type, "commit", change.New.Target.Hash)
			continue
		}

		if matched, _ := regexp.MatchString("#"+cfg.Check.MessageTagName, change.New.Target.Message); !matched {
			log.Printf("not matched message='%s', tag='%s'", change.New.Target.Message, cfg.Check.MessageTagName)
			continue
		}

		log.Printf("hook from token='%s', provider='%s', hash='%s'", cfg.Token, cfg.ProviderName, change.New.Target.Hash)
		return true
	}

	return false
}
