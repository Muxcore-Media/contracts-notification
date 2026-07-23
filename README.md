# contracts-notification

Protobuf/gRPC contract interfaces for notification modules in MuxCore.

## Proto Service

- **`NotificationService`** (`proto/muxcore/notification/v1/notification.proto`)
  - `Notify` — send a notification to one or more channels
  - `Configure` — enable or update per-channel settings at runtime
  - `Status` — list configured channels and their state

### `Channel` enum

| Value | Name |
|------:|------|
| 0 | `CHANNEL_UNSPECIFIED` |
| 1 | `CHANNEL_DISCORD` |
| 2 | `CHANNEL_SLACK` |
| 3 | `CHANNEL_WEBHOOK` |
| 4 | `CHANNEL_EMAIL` |
| 5 | `CHANNEL_APPRISE` |

Logical contract name: `NotificationProvider` v0.1.0.

## Go import

Generated stubs: `github.com/Muxcore-Media/contracts-notification/muxcore/notification/v1`

## Implementing Modules

Any module advertising capability `notification` may implement this service.

### Known implementers

- [notification-default](https://github.com/Muxcore-Media/notification-default) — Discord, Slack, webhook, email
- [notification-apprise](https://github.com/Muxcore-Media/notification-apprise) — Apprise plus Discord/Slack/webhook
