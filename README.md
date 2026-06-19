# cc-util

CapCut desktop utility — a terminal UI for patching, version-locking, and downloading legacy CapCut builds on Windows.

## Features

- **Patch** — Get many features by patching `VECreator.dll`
- **Lock** — Lock CapCut version
- **Download** — Browse and open legacy CapCut versions (1.0.0–5.4.0 Beta6) in browser

## Build

```
make build
```

Requires Go and [`go-winres`](https://github.com/tc-hib/go-winres) for Windows resource embedding. Cross-compiles from any OS (`GOOS=windows GOARCH=amd64`).

## Usage

Run `cc-util.exe` on a Windows machine with CapCut installed. The TUI walks through prechecks (OS, installation, running state) then presents the main menu.

```
cc-util.exe --version
```

## License

[MIT](LICENSE)

## Credits

- [iosdevice](https://iosgods.com/profile/6266834-iosdevice/) — patches
