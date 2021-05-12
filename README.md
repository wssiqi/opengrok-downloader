# opengrok-downloader
a simple opengrok downloader.  Can download entire folder from opengrok
## build  
### 1. download code
`git clone https://github.com/wssiqi/opengrok-downloader`
### 2. build package
`cd opengrok-downloader`  
`go build`  
you can find the executable binary file is locate in current path  
on windows, it will be:
`opengrok-downloader.exe`  
on linux, it will be:
`opengrok-downloader`

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