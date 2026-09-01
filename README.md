# contracts-notification

Protobuf/gRPC contract interfaces for notification modules in MuxCore.

## Proto Service

- **`NotificationService`** (`proto/muxcore/notification/v1/notification.proto`) — logical interface name **`NotificationProvider`** (used by `notification-default`, `notification-apprise`, and `contracts-reconciler`).
  - `Notify` — send a notification to one or more channels
  - `Configure` — enable, update, or disable per-channel settings at runtime
  - `Status` — list configured channels and their state
  - `Test` — send a fixture ping to one channel (admin-ui / muxcorectl smoke)

Contract repo capability: **`contracts.notification`**. Implementers advertise capability **`notification`**.

### `Channel` enum

| Value | Name |
|------:|------|
| 0 | `CHANNEL_UNSPECIFIED` |
| 1 | `CHANNEL_DISCORD` |
| 2 | `CHANNEL_SLACK` |
| 3 | `CHANNEL_WEBHOOK` |
| 4 | `CHANNEL_EMAIL` |
| 5 | `CHANNEL_APPRISE` |

### `Severity` enum

| Value | Name |
|------:|------|
| 0 | `SEVERITY_UNSPECIFIED` |
| 1 | `SEVERITY_INFO` |
| 2 | `SEVERITY_SUCCESS` |
| 3 | `SEVERITY_WARNING` |
| 4 | `SEVERITY_ERROR` |

### `ConfigureRequest.settings` keys

| Channel(s) | Key | Description |
|------------|-----|-------------|
| Discord, Slack, Webhook | `webhook_url` | HTTPS webhook endpoint |
| Email | `smtp_host` | SMTP server hostname |
| Email | `smtp_port` | SMTP port (default 587 if omitted by implementer) |
| Email | `smtp_user` | SMTP username |
| Email | `smtp_pass` | SMTP password |
| Email | `smtp_from` | From address |
| Email | `to` | Recipient address |
| Apprise | `urls` | Comma-separated Apprise notification URLs |
| Apprise | `token` | Optional Apprise API token |

Set `ConfigureRequest.enabled = false` to disable a channel. Non-portable aliases (`host`, `from`, `pass`, …) must not be used in new callers.

## Go import

Generated stubs: `github.com/Muxcore-Media/contracts-notification/muxcore/notification/v1`

Regenerate after proto edits:

```bash
make proto
go test ./...
```

## Implementing Modules

Any module advertising capability `notification` may implement this service.

### Known implementers

- [notification-default](https://github.com/Muxcore-Media/notification-default) — Discord, Slack, webhook, email
- [notification-apprise](https://github.com/Muxcore-Media/notification-apprise) — Apprise plus Discord/Slack/webhook

See [COMPATIBILITY.md](COMPATIBILITY.md) and [CHANGELOG.md](CHANGELOG.md).
