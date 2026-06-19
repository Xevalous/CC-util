# cc-util

CapCut desktop utility — TUI (Bubble Tea) with patch, version-lock, and download features.

## Build

```
make build    # cross-compile to cc-util.exe (Windows amd64)
```

Requires `go-winres` for embedding Windows resources (`rsrc_windows_amd64.syso` is gitignored). Both `.syso` and `cc-util.exe` are in `.gitignore`.

No tests, no lint, no CI. There is no test command.

## Structure

- `main.go` — entrypoint, runs Bubble Tea program
- `ui/` — TUI views (precheck → menu → patch/lock/download)
- `patch/` — DLL patching logic (binary byte replacement in VECreator.dll)
- `lock/` — version lock via ProductInfo.xml + Windows `attrib`
- `util/process.go` — CapCut process management (`tasklist`/`taskkill`), app dir resolution

## Flow

```
Precheck → Menu → [Patch | Lock | Download]
```

Download view includes scrollable list of legacy CapCut versions (1.0.0–5.4.0 Beta6) with 4.0.0 as recommended. Opens in browser or shows raw URL as fallback.

## Platform

Windows-only at runtime. Uses `LOCALAPPDATA`, `tasklist`, `taskkill`, `attrib`. Cross-compiles from any OS with `GOOS=windows GOARCH=amd64`.
