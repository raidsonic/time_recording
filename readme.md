The content provided is already in Markdown format. Here is the same content again for your reference:

# Time Recording Project

This project is a command-line tool for recording work time entries. It is written in Go and uses a configuration file for settings.

## Features

- Parse and validate command-line arguments for time, task, day, and project.
- Format time entries.
- Load settings from a JSON configuration file.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/raidsonic/time-recording.git
    cd time-recording
    ```

2. Build the project:
    ```sh
    go build
    ```

## Usage

Run the command with the required arguments:
```sh
./time-recording --time 0800-1600 --task "Development work" --day 2024-10-01 --project "ProjectName"
```

## Configuration

The configuration file `time_recording.config.json` should be placed in the same directory as the executable. Example configuration:
```json
{
    "project-default": "default",
    "filepath": "~/time_records"
}
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.