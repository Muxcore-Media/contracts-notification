# Changelog

## [0.1.1] — 2026-08-31

### Added
- `Severity` enum on `NotifyRequest` (`INFO` / `SUCCESS` / `WARNING` / `ERROR`) replacing free-string severity.
- `Test` RPC for fixture channel pings (`TestRequest` / `TestResponse`).
- `ConfigureRequest.enabled` to disable or clear a channel without deleting the row.
- `ChannelStatus.last_error` and `ChannelStatus.last_success_at` for operator visibility.
- Documented canonical `ConfigureRequest.settings` keys in proto comments and README.
- Forgejo `proto-check` CI job; `compat_test.go` freezing enum numbers, RPC names, and field tags.
- `COMPATIBILITY.md`; proto plugin version pins in `Makefile`.

### Changed
- `muxcore.json` version **0.1.1** (tag `v0.1.1`).
- `Makefile` `clean` no longer deletes published stubs under `muxcore/notification/v1/`.
- `AGENTS.md` rewritten for a proto/stubs contract repo (not a gRPC sidecar).

### Compatibility
- Breaking wire change: `NotifyRequest.severity` is now `Severity` enum (field 3). Consumers must regenerate stubs and map enum values; do not send raw strings.
- Additive: `ConfigureRequest.enabled` (field 3), `ChannelStatus.last_error` (4), `last_success_at` (5), `Test` RPC.

## [0.1.0] — initial

### Added
- Initial `NotificationService` contract (`Notify`, `Configure`, `Status`).
- Generated Go stubs under `muxcore/notification/v1`.
