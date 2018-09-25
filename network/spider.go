package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"log"
	"os"
	"strconv"
	"fmt"
	"strings"
	"io"
	"runtime"
	"time"
)



func main() {
	fmt.Println("runtime on " + strconv.Itoa(runtime.NumCPU()) + " goroutine")
	runtime.GOMAXPROCS(runtime.NumCPU())
	url := "http://taohuabt.cc/thread-1820205-1-index.html"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	imgSrcList := ExtractImageSrcList(resp.Body)
	DownloadImage(imgSrcList, runtime.NumCPU())
}

func ExtractImageSrcList(body io.ReadCloser)[]string  {
	docs, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	imageElements := docs.Find(".zoom")
	imageSrcList := make([]string, len(imageElements.Nodes))
	imageElements.Each(func(i int, selection *goquery.Selection) {
		value, ok := selection.Attr("file")
		if ok {
			imageSrcList[i] = value
		}
	})
	return imageSrcList
}


func DownloadImage(imageList []string, cpuNums int)  {
	ch := make(chan int, cpuNums)
	imageMap := make(map[int]map[int]string)
	for i := 0; i < cpuNums; i++{
		imageMap[i] = make(map[int]string)
	}
	for index, imageSrc := range imageList {
		mapType := index % cpuNums
		imageMap[mapType][index] = imageSrc
	}

	for i := 0; i < cpuNums; i++ {
		go Download(imageMap[i], ch, i)
	}
	for i := 0; i < cpuNums; i++ {
		<-  ch
	}
}

func Download(imageMap map[int]string, ch chan <- int, mapIndex int)  {
	for index, imageSrc := range imageMap {
		resp, err := http.Get(imageSrc)
		if err != nil {
			log.Fatal(err)
			continue
		}
		fileLength, _ := strconv.Atoi(resp.Header.Get("content-length"))
		fileType := strings.Split(resp.Header.Get("content-type"), "/")[1]
		content := make([]byte, fileLength)
		resp.Body.Read(content)
		if fileType == "jpeg" {
			fileType = "jpg"
		}
		fileName := "./images/" + strconv.Itoa(index) + "." + fileType
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
			break
		}

		file.Write(content)
		file.Close()
		resp.Body.Close()
		fmt.Println("file " + strconv.Itoa(index) + " downloaded")
		time.Sleep(1e9)
	}
	ch <- mapIndex
}

