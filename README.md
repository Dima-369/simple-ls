## Installation

`go get github.com/Gira-X/simple-ls`

This generates the command `simple-ls` in `GOPATH/bin` or `GOBIN`.

## Options

It only supports a single flag `hidden` which only lists hidden directories and files.

Note that `.DS_Store` is not considered in the `... hidden files` count to avoid many directory listings
where the `.DS_Store` is the only hidden file.

## Output


