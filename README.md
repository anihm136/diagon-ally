# diagon-ally
(Yes I know it is a terrible pun at best)
`diagon-ally` is a CLI tool to perform transformative actions when files in a directory are changed i.e, generate/export files when a file is created/modified in a directory. In short, it is a CLI wrapper around a file watcher, with some added conveniences. It is simple to use, with minimal configuration and functionality - it is designed to do one thing (fairly) well. It was originally designed for [org-diagrams](https://github.com/org-diagrams), but can be used as a generic file-watcher independently as well.

## Build instructions
Build the binary with
```sh
go build . -o diagon
```
Place the binary on your path. That's it

## Usage
```
Diagon-ally watches a directory of template images, and exports them to
destination images whenever a new template is created or an existing
template is edited.

Usage:
	diagon [flags]

Flags:
	-d, --dest_dir string     Directory to export to
	-f, --force               Force creation of settings if it does not exist
	-h, --help                help for diagon
	-u, --on_update string    Command to run for export on update
	-s, --source_dir string   Directory to watch for changes
```

## Configuration
`diagon-ally` provides a simple configuration file, which offers only 3 options currently -
1. `WatchDir`: Path to watch
2. `ExportDir`: Path to export files to
3. `OnUpdate`: Command to run when a file in `WatchDir` is created/modified
All relative paths are resolved from the directory in which the tool is started. `OnUpdate` supports two variables - `${OUT}`, the path to export to, and `${IN}`, the path of the changed file to use as the source

The configuration file is located in the default XDG config locations - `~/.config/diagon-ally/settings.txt` on \*NIX, `~/Library/Application Support/diagon-ally/settings.txt` on macOS and `%LOCALAPPDATA%\diagon-ally\settings.txt` on Windows. If no configuration file is found in these locations, default values are used
* `WatchDir`: `./svg`
* `ExportDir`: `./images`
* `OnUpdate`: `inkscape --export-area-drawing --export-area-snap --export-type=png -o ${OUT} ${IN}`
## Dependencies
This tool is written in Golang, and uses the following external libraries -
* [xdg](https://github.com/adrg/xdg) for getting platform-dependent XDG directories for the settings file
* [notify](https://github.com/rjeczalik/notify) for watching file events
* [cobra](https://github.com/spf13/cobra) for generating the CLI tool
In addition, there is a soft dependency on [InkScape](https://inkscape.org/) as the default export tool
