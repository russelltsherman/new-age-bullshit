# New Age Bullshit Generator

Inspired by [sebpearce](https://github.com/sebpearce/bullshit)

cli application which generates pseudo random new age text
written in golang and includes options to send generated text to TTS engine.

includes output options for the following TTS engines

- `cepstral`
- `espeak`
- `festival`
- `pico`
- `say`

## load dependencies

```sh
GO111MODULE=on go mod download
```

## run

```sh
go run app.go --help
```
