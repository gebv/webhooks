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

## TODO

* github [doc](https://developer.github.com/webhooks/)
* maybe universal listener interface to any events
