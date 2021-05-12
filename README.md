# opengrok-downloader
a simple opengrok downloader.  Can download an entire folder from opengrok
## download execute binary from `https://github.com/wssiqi/opengrok-downloader/releases  

## windows usage
```bash
>opengrok-downloader.exe -h
Usage of opengrok-downloader.exe:
  -o string
        download folder, default is . (means download to current folder)
  -u string
        download url, such as http://www.opengrok-server.com/xxx/xxx
```
```
>opengrok-downloader.exe -o myFolder -u http://www.opengrok-server.com/xxx/xxx
15:04:33.771780 myFolder\src  folder created
15:04:35.928120 myFolder\src\main.go downloaded
15:04:36.637137 myFolder\src\test.go downloaded
```


## linux usage
```bash
sh opengrok-downloader -h
Usage of opengrok-downloader.exe:
  -o string
        download folder, default is . (means download to current folder)
  -u string
        download url, such as http://www.opengrok-server.com/xxx/xxx
```
```
$opengrok-downloader.exe -o myFolder -u http://www.opengrok-server.com/xxx/xxx
15:04:33.771780 myFolder/src  folder created
15:04:35.928120 myFolder/src/main.go downloaded
15:04:36.637137 myFolder/src/test.go downloaded
```