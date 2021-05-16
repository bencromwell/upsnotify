# upsnotify

Called by UPSmon, in turn calls Mailgun's API.

See `upsmon.conf` at `/etc/nut/upsmon.conf`
Place this binary in the path defined as `NOTIFYCMD`:

```
NOTIFYCMD /usr/local/bin/upsnotify
```

UPSmon calls the script with the environment variables `NOTIFYTYPE` and `UPSNAME`. The only arg provided is the event message.

## config

Provided by environment variables or `/etc/default/upsnotify`

```
MAILGUN_PRIVATE_KEY=key-foo
MAILGUN_DOMAIN=mg.example.org
ALERT_EMAIL=youremail@example.com
ALERT_FROM=upsnotifer@example.org
```
