# keyspeed - a typing speed applet for i3blocks

## Install

 - clone the repo
 ```bash
 git clone github.com/indeedhat/keyspeed
 ```

 - install
 ```bash
 make
 sudo make install
 ```

 - disable password requirement for sudo
 > making any executable not require a password is dangerous, do this at your own risk
 ```bash
 sudo echo "$USER ALL=(root) NOPASSWD: /usr/bin/keyspeed" >> /etc/sudoers
 ```

 - Add to your i3 blocks config
 ```
[cpm]
interval=persist
command=sudo keyspeed
```

## Help
```
Typing speed for i3blocks

Options:

    --help, -h[=false]
    display this message

    --cpm, -c[=false]
    display characters per minute (wpm is default)

    --pad, -p[=3]
    pad the display with leading 0's

    --best, -b[=false]
    keep track of your best time for this session

    --interval, -i[=5]
    Polling interval to uodate the count in seconds
```
