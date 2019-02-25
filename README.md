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

**Tracking time**  
Track time. Specify a description and hit enter to start tracking time.
```
./timetracker
```
You can add the flag `--cli` to not use the tray for stopping the time and instead use the command line.

**List**  
List all saved entries
```
./timetracker list
```

**Clear**  
Clear all saved entries
```
./timetracker clear
```

**Delete**  
Delete the entry with the specified ID
```
./timetracker delete [ID]
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