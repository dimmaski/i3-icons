# i3-icons-daemon

<p align="center">
   <img src="assets/i3-icons.png" alt="i3"/>
</p>

```sh
go build .
```

```sh
~ » cat ~/.i3/icons.json
{
    "firefox": "\uf269",
    "spotify": "\uf1bc",
    "chrome": "\uf268",
    "code": "\uf121",
    "edit": "\uf044",
    "nautilus": "\uf07b",
    "vlc": "\uf04b",
    "terminal": "\uf120",
    "_wildcard": "\uf128"
}
```

```sh
# add the following line in your i3 config
exec_always i3-icons -separator "|" -config ~/.i3/icons.json
```
