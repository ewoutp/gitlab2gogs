# gitlab2gogs

Migrate your GitLab 9.x repositories to Gogs.

Usage:

```
./gitlab2gogs -gitlab-host https://<yourgitlabhost> \
    -gitlab-api-path /api/v4
    -gitlab-token <your gitlab token> \
    -gitlab-user <gitlab admin user> \
    -gitlab-password <password of gitlab-user> \
    -gitlab-visibilitylevel {private|internal|public} \
    -gogs-url https://<yourgogshost> \
    -gogs-token <your gogs token> \
    -gogs-user <gogs admin username>
```

Organizations are created if they do not yet exists.
Users are created if they do not yet exists.
Existing repositories (in Gogs) are not overwritten.
