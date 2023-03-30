# darkmode

## Installation

Run following command.

```shell
go install github.com/orangekame3/darkmode@latest
```

## Usage 

First, run the following command to generate a configuration file.

```shell
darkmode init
```

The following file is generated when the command is executed. Please disable the appropriate platform comment-outs.

```yaml:darkmode.yaml
desktop:
  environment: windows
#  environment: gnome
#  environment: kde
#  environment: xfce

```
