# Functionality
- Run a file watcher in the background on a specific directory
- When an svg changes, export it as png to another directory
- Copy the insertion template into clipboard
- If a new file is to be created, copy a template into the watched directory

## Config
- Plain text file with key-value pairs
	- Directory to watch (absolute/relative)
	- Directory to export to
	- Template file for new file
	- Template file for insertion
- Defaults
	- Watch `svg/` under cwd
	- Export to `images/` under cwd
	- Use default template for new files and insertion, bundled in binary
- Config dir determined by `xdg` module, used XDG specifications

## Watcher
- Use `notify` module for file watcher
- Run subprocess using config values to export

## CLI
- Start server in a directory (and start watcher)
- Create new file (connect to existing server or start new one)
