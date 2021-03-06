# i3-icons-daemon

Is a lightweight golang daemon service that updates your i3 workspace based on the programs that are running in your workspace's windows.

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

## i3 conf

```sh
# add the following line in your i3 config
exec_always i3-icons -separator "|" -config ~/.i3/icons.json
```

```sh
# change the following confis in order to allow switching and moving workspaces

# from this
bindsym $mod+N workspace N
bindsym $mod+Shift+N move container to workspace N

# to this
bindsym $mod+N workspace number N
bindsym $mod+Shift+N move container to workspace number N
```
