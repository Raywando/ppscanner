# ppscanner

A tool to detect prototype pollution based on popular vulnerable libraries: https://github.com/BlackFan/client-side-prototype-pollution

# Installation

```
sudo apt install -y chromium-browser
go get -u github.com/Raywando/ppscanner
```

# Demo

`cat urls.txt | ppscanner`

`echo https://target.com | ppscanner`

![image](https://user-images.githubusercontent.com/33800255/120795894-c34c4100-c542-11eb-9d9b-51414b1ee789.png)
