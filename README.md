# Tracker

It's kind of stupid... just lets you save things from the command line
into a file and list them back out... mostly just something silly for me
to work on in my spare time.

### Why

The main reason was I wanted something simple that I could run from the command
line to remember things that I wanted to bring up at the retrospective.

`Retrium` - the app we use - doesn't have an API, CLI, or Slack app, so I
figured I would just throw something together.

Been a while since I'd had a chance to write some `Go` code, so I went with
that.

### Command line

> tracker setup

Create a config from the defaults and ensure the directory exists

> tracker list

List all tracks

> tracker add [name] [text]

Create a new track with the given name

> tracker show [name]

Display the entries in the track

> tracker clear [name]

Remove the entries in the track
