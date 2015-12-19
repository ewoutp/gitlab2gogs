# gitlab2gogs

Migrate your Gitlab repositories to Gogs.

Usage:

```
./gitlab2gogs -gitlab-host https://<yourgitlabhost> \
    -gitlab-api-path /api/v3
    -gitlab-token <your gitlab token> \
    -gitlab-user <gitlab admin user> \
    -gitlab-password <password of gitlab-user> \
    -gogs-url https://<yourgogshost> \
    -gogs-token <your gogs token> \
    -gogs-user <gogs admin username>
```

Organizations are created if they do not yet exists.
Existing repositories (in Gogs) are not overwritten.
