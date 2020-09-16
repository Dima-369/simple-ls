## Installation

`go get github.com/Gira-X/simple-ls`

This generates the command `simple-ls` in `GOPATH/bin` or `GOBIN`.

## Options

It only supports the command argument `hidden` which only lists hidden directories and files.

Note that `.DS_Store` is not considered in the `... hidden files` count to avoid many directory listings
where the `.DS_Store` is the only hidden file.

## Output



## Easy bash usage

I rarely use the default `ls` and I am instead using those aliases in [fish](https://fishshell.com/)
with `simple-ls` renamed to `l`.

```bash
> type l

l is /Users/Gira/Documents/go/bin/l
```

```bash
> type lh

lh is a function with definition
# Defined in /Users/Gira/.config/fish/config.fish @ line 174
function lh
	l hidden
end
```

