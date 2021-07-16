# opengrok-downloader
a simple opengrok downloader.  Can download an entire folder from opengrok
## download
download execute binary from `https://github.com/wssiqi/opengrok-downloader/releases  

## usage
### windows usage
```bash
>opengrok-downloader.exe -h
Usage of opengrok-downloader.exe:
  -o string
        download folder, default is . (means download to current folder)
  -u string
        download url, such as http://www.opengrok-server.com/xxx/xxx/
```
```
>opengrok-downloader.exe -o myFolder -u http://www.opengrok-server.com/xxx/xxx/
15:04:33.771780 myFolder\src  folder created
15:04:35.928120 myFolder\src\main.go downloaded
15:04:36.637137 myFolder\src\test.go downloaded
```


### linux usage
```bash
sh opengrok-downloader -h
Usage of opengrok-downloader.exe:
  -o string
        download folder, default is . (means download to current folder)
  -u string
        download url, such as http://www.opengrok-server.com/xxx/xxx/
```
```
$opengrok-downloader.exe -o myFolder -u http://www.opengrok-server.com/xxx/xxx/
15:04:33.771780 myFolder/src  folder created
15:04:35.928120 myFolder/src/main.go downloaded
15:04:36.637137 myFolder/src/test.go downloaded
```

## proxy
opengrok-downloader will automatic load proxy settings from environment "http_proxy" or "HTTP_PROXY", "https_proxy" or "HTTPS_PROXY"  
if the proxy not work, you can try follow settings
### windows
```
>set http_proxy=1.1.1.1:8080
>set https_proxy=1.1.1.1:8080
>opengrok-downloader.exe -o android -u http://androidxref.com/9.0.0_r3/xref/bionic/
```

### linux
```
$ export http_proxy=1.1.1.1:8080
$ export https_proxy=1.1.1.1:8080
$sh opengrok-downloader -o android -u http://androidxref.com/9.0.0_r3/xref/bionic/
```
