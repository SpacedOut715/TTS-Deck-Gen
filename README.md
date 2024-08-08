# What is it?

This is a small tool used for creating custom deck images for TTS (TableTop Simulator). In TTS deck can be "cut-down" from a single image that is max 10000x10000 pixels, or composed of 10x7 cards.

## Build

Use standard go build command.

```bash
go build
```

## Usage

```go
./tts-deck-gen $COMMAND --param1 --param2

where $COMMAND can be:
# auto-locate 
  params:
  --search-dir - path to root directory that will be searched, only END subdirectories are valid.
  (if test-images directory is provided, valid directories would be: A-green, B-big-deck, C-yellow)
  --export-dir - path to directory where deck images will be exported.
# with-config
  params:
  --config-path - path to json configuration file
```

### Example usage

```
./tts-deck-gen auto-locate --search-dir "D:/Path/To/Root/Deck/Directory" --export-dir "D:/Path/To/Generated"
./tts-deck-gen with-config --config-path "D:/Path/To/Config/Directory/config.json"
```

### Example config

```
{
    "decks": [
        {
            "deckPath": "D:/Path/To/Deck/Directory",
            "deckFileName": "Example_Name1"
        },
        {
            "deckPath": "D:/Path/To/Another/Deck/Directory",
            "deckFileName": "Example_Name2"
        }
    ],
    "exportPath": "D:/Path/To/Generated"
}
```

## License

Do whatever the f*** you want.
