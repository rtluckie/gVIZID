# CLI Specification

The CLI is built with Cobra + Viper.

Binary name: `gvizid`

## Config

Default config path:

- `~/.config/vizid/config.yaml`

Override via:

- `--config, -c /path/to/config.yaml`

Config keys (v1):

```yaml
timezone: "UTC"
custom: false
warn: true

components:
  year: true
  month: true
  day: true
  hour: true
  minute: true
  second: true
  ms: true
  uuid: true
```

## Commands

### `gvizid gen`

Generate a new VIZID (prints visual form).

Flags:

- `--config, -c` alternate config file
- `--timezone, -t` timezone (default `UTC`)
- `--user-defined, -u` toggles all components off, user must define what components they want.
- `--warn` warn if sort order may break

- component toggles (bool)[default=`TRUE`]:

  Allows the default (and user defined config) components to be toggled on/off
  When used in combination is the `--user-defined` / `-u` flag.
  When `-u` is passed all components are turned off
  - `--year`
  - `--month`
  - `--day`
  - `--hour`
  - `--minute`
  - `--second`
  - `--ms`
  - `--uuid`

### `gvizid decode <vizid>`

Decode a VIZID into its ASCII form:

```
YYYYMMDDhhmmssmmm-PTTCCR
```

### `gvizid encode <ascii>`

Encode an ASCII ID into VIZ form.

---

## Sort order warnings

If custom component toggles disable any high-significance timestamp component while leaving lower-significance components enabled, chronological sort order may break.

Examples of risky customization:

- disabling `year` but keeping `month`
- disabling `day` but keeping `hour`

The CLI MUST warn when `--warn=true`.
