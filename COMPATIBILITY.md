# Compatibility

| Component | Requirement |
|-----------|-------------|
| Contract interface | `NotificationProvider` **v0.1.1** (gRPC service `NotificationService`) |
| Go module path | `github.com/Muxcore-Media/contracts-notification` |
| Generated package | `github.com/Muxcore-Media/contracts-notification/muxcore/notification/v1` |
| MuxCore core | ≥ 0.4.0 (consumers) |
| Capability ID (contract repo) | `contracts.notification` |
| Capability ID (implementers / callers) | `notification` |

Implementers register capability **`notification`** in `muxcore.json` and implement gRPC service **`NotificationService`**. The logical interface name in `contracts-reconciler` and module manifests is **`NotificationProvider`**. The contract repository capability id **`contracts.notification`** identifies this proto module in the umbrella catalog.

Known implementers: `notification-default`, `notification-apprise`.

Known consumers (event → Notify): `media-subtitles`, `playback-monitor`, `playback-guard`, `media-library-maintainer`.

## Severity vocabulary

Portable `Severity` enum values (Discord/Slack color mapping):

| Enum | Meaning |
|------|---------|
| `SEVERITY_INFO` | Informational |
| `SEVERITY_SUCCESS` | Success / completed |
| `SEVERITY_WARNING` | Warning / attention |
| `SEVERITY_ERROR` | Error / failure |

## Configure settings keys

| Channel | Keys |
|---------|------|
| Discord, Slack, Webhook | `webhook_url` |
| Email | `smtp_host`, `smtp_port`, `smtp_user`, `smtp_pass`, `smtp_from`, `to` |
| Apprise | `urls`, `token` |

Set `ConfigureRequest.enabled = false` to disable a channel. `ChannelStatus.description` must not contain secrets (URLs, passwords, tokens).

Breaking proto changes require a new major contract version and coordinated consumer updates.
