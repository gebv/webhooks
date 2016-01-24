# Webhook

Event listener from a variety of services.

``` shell
./bin/app.bin -addr=127.0.0.1:8081
```

# Supported providers

* Bitbucket webhook push [doc](https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-Push)
* 

# Example config for bitbucket webhook push

config.json
``` json
{
    "Hooks": [
        {
            "Token": "93e42247fc03",
            "ProviderName": "bitbucket:push",
            "Description": "",
            "ShellCommand": "workspace/deploy.sh",
            "Check": {
                "UserName": "username",
                "BranchName": "branchname",
                "RepositoryName": "repositoryname",
                "Tag": "deploy"
            },
            "IsEnabled": true
        }
    ]
}
```

Checks
* branch name
* repository name
* author commit
* tag in the commit message

# My example for hugo site

[Hugo](https://gohugo.io) great project!
And so, what would the project automatically updated after ```git push```. 

Your project has the following structure
```
├── Makefile // magic makefile
├── blog // hugo site
│   ├── archetypes
│   ├── config.toml
│   ├── content
│   ├── data
│   ├── images
│   ├── layouts
│   ├── public
│   ├── static
│   └── themes
├── config.json // webhooks config
└── deploy.sh // for webhooks config ShellCommand
```

Already installed hugo and webhooks.

## Makefile

All the magic in the Makefile.
``` Makefile
.PHONY: deploy hugo_start hugo_stop webhooks_start webhooks_stop

PIDFILEHUGO := ./hugo.pid
PIDFILEWEBHOOKS := ./webhooks.pid
HUGOBASEURL := "https://example.com/cms/"
PORT := "8082"

ifneq ("$(wildcard $(PIDFILEHUGO))","")
    PIDHUGO=$(shell cat $(PIDFILEHUGO))
endif

ifneq ("$(wildcard $(PIDFILEWEBHOOKS))","")
    PIDWEBHOOKS=$(shell cat $(PIDFILEWEBHOOKS))
endif

defult: deploy

deploy:
    @echo Deploy...
    @echo "$(BUILDSTAMP)]" >> ./logs/deploy.log
    $(shell git pull >> ./logs/deploy.log)

hugo_start:
ifndef PIDHUGO
    @echo Starting hugo server...
    @$(shell hugo server -w --port=$(PORT) --source=./blog --config=./blog/config.toml --renderToDisk --baseURL=$(HUGOBASEURL) --appendPort=false > ./logs/hugo_current.log 2>&1 & echo $$! > $(PIDFILEHUGO))
    @echo PID=$(shell cat $(PIDFILEHUGO))
else
    @echo PID='$(PIDHUGO)'
endif

hugo_stop:
ifdef PIDHUGO
    @echo PID=$(PIDHUGO) Stopping hugo server...
    @rm $(PIDFILEHUGO)
    @kill -9 $(PIDHUGO)
    @echo "================END================" >> ./logs/hugo_current.log
    @echo $(BUILDSTAMP) >> ./logs/hugo_current.log
    @cp ./logs/hugo_current.log ./logs/hugo_$(BUILDSTAMP).log
    @echo > ./logs/hugo_current.log
else
    @echo It is stopped
endif

webhooks_start:
ifndef PIDWEBHOOKS
    @echo Starting webhooks listener...
    @$(shell webhooks -addr=127.0.0.1:8081 > ./logs/webhooks_current.log 2>&1 & echo $$! > $(PIDFILEWEBHOOKS))
    @echo PID=$(shell cat $(PIDFILEWEBHOOKS))
else
    @echo PID='$(PIDWEBHOOKS)'
endif

webhooks_stop:
ifdef PIDWEBHOOKS
    @echo PID=$(PIDWEBHOOKS) Stopping webhooks listener...
    @rm $(PIDFILEWEBHOOKS)
    @kill -9 $(PIDWEBHOOKS)
    @echo "================END================" >> ./logs/webhooks_current.log
    @echo $(BUILDSTAMP) >> ./logs/webhooks_current.log
    @cp ./logs/webhooks_current.log ./logs/webhooks_$(BUILDSTAMP).log
    @echo > ./logs/webhooks_current.log
else
    @echo It is stopped
endif
```

## Nginx config
```
server {
    listen       80;
    server_name example.com;

    location /cms/webhooks {
# only IPs bitbucket
        proxy_set_header X-Real-IP  $remote_addr;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host $host;
        proxy_pass http://127.0.0.1:8081;
    }

    location /cms/ {
        proxy_set_header X-Real-IP  $remote_addr;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host $host;
        proxy_pass http://127.0.0.1:8082; # port hugo
    }
}
```

[Bitbucket IPs](https://confluence.atlassian.com/bitbucket/manage-webhooks-735643732.html#Managewebhooks-trigger_webhook)

## webhooks config.json
``` json
{
    "Hooks": [
        {
            "Token": "fc0393e4",
            "ProviderName": "bitbucket:push",
            "Description": "",
            "ShellCommand": "./deploy.sh",
            "Check": {
                "UserName": "username",
                "BranchName": "master",
                "RepositoryName": "examplecom",
                "Tag": "deploy"
            },
            "IsEnabled": true
        }
    ]
}
```

## Start

Git clone and
* Configure the Bitbucket webhooks. URL = http://example.com/cms/webhooks?webhook_key=fc0393e4
* make hugo_start
* make webhooks_start

If the commit message contains a tag **#deploy**, webhooks will executed and will updated the site.

# TODO

* github [doc](https://developer.github.com/webhooks/)
* maybe universal listener interface to any events
* rename to something like webeventlistener~wel~elvis
