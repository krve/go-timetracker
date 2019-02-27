[![CircleCI](https://circleci.com/gh/krve/go-timetracker.svg?style=svg)](https://circleci.com/gh/krve/go-timetracker)

# Timetracker
This is a timetracker CLI app I am currently building to learn Go.

## Building
```
git clone git@github.com:krve/go-timetracker.git
cd go-timetracker
go get
go build
```

## Usage
```
USAGE:
   timetracker [global options] command [command options] [arguments...]

COMMANDS:
     list, l  show the list of entries
     clear    clear all entries
     reindex  reindex all entires
     delete   delete the entry with the specified id
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Testing
```
git clone git@github.com:krve/go-timetracker.git
cd go-timetracker
go get
go test
```

## Contributing
I'm not accepting contributions as this is currently only for personal use. I may open up for contribution in the future.
