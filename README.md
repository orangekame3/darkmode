# darkmode

## Installation

Run following command.

```shell
go install github.com/orangekame3/darkmode@latest
```

## Usage 

First, run the following command to generate a configuration file.

```shell
darkmode init --env windows
```

The following file is generated when the command is executed. Please disable the appropriate platform comment-outs.

```yaml:darkmode.yaml
desktop:
  environment: windows
  dark-theme: none
  light-theme: none
```


If you using gnome-desktop, run the following command.

```shell
darkmode init --env gnome
```

The following file is generated when the command is executed. Please disable the appropriate platform comment-outs.

```yaml:darkmode.yaml
desktop:
  environment: gnome
  dark-theme: Adwaita-dark
  light-theme: Adwaita
```
