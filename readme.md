# Moon Phase Emoji Generator

This **Go** program determines the current phase of the Moon and displays it as an emoji. It can be integrated into an environment like **PowerShell** with **Oh My Posh** to show the lunar phase in your prompt.

## Features
- Calculates the current phase of the Moon based on a reference new moon.
- Returns an emoji corresponding to the phase:
- Useful for creating an environment variable usable in your PowerShell prompt.

## Installation
1. **Prerequisites:** Make sure Go is installed on your system. If not, download it from [golang.org](https://golang.org/).

2. **Compile and install the program:**
   ```bash
   go install github.com/ctrl-vfr/flying-to-the-moon
   ```
   Ensure that `GOPATH/bin` is included in your `PATH` environment variable.

3. **Create the environment variable:**
   Add the following to your **PowerShell** profile (`$PROFILE`) so the program generates a `MOON_PHASE` variable:
   ```powershell
   $Env:MOON_PHASE = & "flying-to-the-moon.exe"
   ```

   **Example:**
   If you're using **Oh My Posh**, add to your theme as follows:
   ```json
   {
     "foreground": "#00deee",
     "powerline_symbol": "\ue0b0",
     "properties": {
       "always_enabled": true
     },
     "style": "plain",
     "template": " {{ if gt .Code 0 }}💥{{ end }}{{ if eq .Code 0 }}{{ .Env.MOON_PHASE }} {{ end }} ",
     "type": "exit"
   }
   ```

   This will display the lunar phase emoji directly in your prompt.

## Usage
Simply run the program to display the lunar phase emoji in the console:
```bash
./flying-to-the-moon
```

## Dependencies
- **Go Language** (standard library)
