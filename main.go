package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	var baseFolder string
	var folderUrl string
	flag.StringVar(&baseFolder, "o", ".", "download folder, default is . (means download to current folder)")
	flag.StringVar(&folderUrl, "u", "", "download url, such as http://www.opengrok-server.com/xxx/xxx")
	flag.Parse()

	_, err := url.ParseRequestURI(folderUrl)
	if err != nil {
		log.Fatalf("invalid url %s, exception %s", folderUrl, err.Error())
	}
	absBaseFolder, err := filepath.Abs(baseFolder)
	if err != nil {
		log.Fatalf("failed to get absolute path for %s", baseFolder)
	}

	err = os.MkdirAll(absBaseFolder, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create directory %s, exception %s", absBaseFolder, err.Error())
	}

	client := getHttpClient()
	defer client.CloseIdleConnections()

	err = downloadFolder(client, absBaseFolder, folderUrl)
	if err != nil {
		log.Fatalf("failed to download %s to %s, exception %s", folderUrl, absBaseFolder, err.Error())
	}
}

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
		},
	}
}

func downloadFolder(client *http.Client, folder string, folderUrl string) error {
	var doc, err = loadUrlAsDoc(client, folderUrl)
	if err == nil {
		list := htmlquery.Find(doc, "//table[@id='dirlist']//a[@href and not(@title)]")
		for _, node := range list {
			childFilename := htmlquery.SelectAttr(node, "href")
			// ignore ..
			if childFilename == ".." {
				continue
			}

			childUrl := folderUrl + childFilename
			childFile := filepath.Join(folder, childFilename)
			if strings.HasSuffix(childFilename, "/") {
				err = os.MkdirAll(childFile, os.ModePerm)
				if err == nil {
					log.Println(childFile, " folder created")
					err = downloadFolder(client, childFile, childUrl)
				}
			} else {
				err = downloadFile(client, childFile, childUrl)
			}

			if err != nil {
				break
			}
		}
	}

	return err
}

func loadUrlAsDoc(client *http.Client, folderUrl string) (*html.Node, error) {
	response, err := client.Get(folderUrl)
	if err == nil {
		body := response.Body
		defer closeBody(body)
		bytes, err := ioutil.ReadAll(body)
		if err == nil {
			return htmlquery.Parse(strings.NewReader(string(bytes)))
		}
	}
	return nil, err
}

func downloadFile(client *http.Client, fileSavePath string, fileUrl string) error {
	var fileDetailDoc, err = loadUrlAsDoc(client, fileUrl)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get file detail page %s, exception %s",
			fileSavePath, err.Error()))
	}
	downloadSpanNode := htmlquery.FindOne(fileDetailDoc, "//a[@href]/span[@id='download']")
	if downloadSpanNode == nil {
		return errors.New(fmt.Sprint("span url not found", fileUrl))
	}
	downloadUrlNode := downloadSpanNode.Parent
	downloadLink := htmlquery.SelectAttr(downloadUrlNode, "href")
	if downloadLink == "" {
		return errors.New(fmt.Sprint("failed to get download link", fileUrl))
	}

	downloadLink = getBaseUrl(fileUrl) + downloadLink
	childTmpFile := fileSavePath + ".tmp"
	err = downloadAndSaveFile(client, downloadLink, childTmpFile)
	if err != nil {
		return err
	}
	err = os.Rename(childTmpFile, fileSavePath)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to rename %s to %s, exception %s",
			childTmpFile, fileSavePath, err.Error()))
	}
	log.Println(fileSavePath, "downloaded")
	return nil
}

func getBaseUrl(url string) string {
	const schemaSeparator = "//"
	const schemaSeparatorLength = len(schemaSeparator)

	index := strings.Index(url, schemaSeparator)
	if index != -1 {
		//s := url[:index+2]
		resourceStart := strings.Index(url[index+schemaSeparatorLength:], "/")
		if resourceStart != -1 {
			return url[:index+resourceStart+schemaSeparatorLength]
		}
	}
	return url
}

func downloadAndSaveFile(client *http.Client, fileUrl string, file string) error {
	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Fatal("failed to download file", fileUrl, err.Error())
	}
	body := resp.Body
	defer closeBody(body)

	fileObj, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("failed to write file", file, err.Error())
		return nil
	}
	defer fileObj.Close()
	buf := make([]byte, 4096)

	for {
		nRead, err := body.Read(buf)

		if err != nil && err != io.EOF {
			return errors.New(fmt.Sprint("failed to read data from resp", fileUrl, err.Error()))
		}

		nWrite, writeErr := fileObj.Write(buf[0:nRead])
		if writeErr != nil {
			return errors.New(fmt.Sprint("failed to write data to file", file, err.Error()))
		}

		if nWrite != nRead {
			return errors.New(fmt.Sprintf("failed to write data to file %s, read=%d, write=%d", file, nRead, nWrite))
		}
		if err == io.EOF {
			break
		}

	}

	return nil
}

func closeBody(body io.ReadCloser) {
	func() {
		closeErr := body.Close()
		if closeErr != nil {
			log.Println("failed to close data", closeErr.Error())
		}
	}()
}
