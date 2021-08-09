package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	h "github.com/Ex0dIa-dev/yt-downloader-v2/helper"
)

type Download struct {
	URL            string
	PathToSave     string
	Sections       int
	SectionsFolder string
}

//return a http.Request with the inputed method
func (dw Download) NewHTTPRequest(method string) *http.Request {

	request, err := http.NewRequest(method, dw.URL, nil)
	h.CheckErr(err)

	request.Header.Set("User-Agent", "Youtube-Downloader v2 by Ex0dIa-dev")
	return request
}

func (dw Download) DownloadStart() {

	req := dw.NewHTTPRequest("HEAD")
	resp, err := http.DefaultClient.Do(req)
	h.CheckErr(err)
	h.CheckResponseStatusCode(resp.StatusCode)

	//getting the size of the file, and the size of each section
	fullSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	h.CheckErr(err)

	if dw.Sections == 0 {
		dw.Sections = 15
	}

	sectionSize := fullSize / dw.Sections

	//creating the sections
	var sections = make([][2]int, dw.Sections)
	for i := range sections {

		if i == 0 { //init first byte of first section
			sections[i][0] = 0
		} else { //init first byte of other sections
			sections[i][0] = sections[i-1][1] + 1
		}

		if i < dw.Sections-1 { //last byte of other section
			sections[i][1] = sections[i][0] + sectionSize
		} else { //last byte of last section
			sections[i][1] = fullSize - 1
		}

	}

	//downloading any section
	var wg sync.WaitGroup
	err = os.Mkdir(dw.SectionsFolder, os.ModePerm)
	h.CheckErr(err)

	for i, s := range sections {
		wg.Add(1)
		go func(i int, s [2]int) {
			defer wg.Done()
			dw.DownloadSection(i, s)
		}(i, s)

	}
	wg.Wait()
	dw.MergeTmpFiles(sections)

}

func (dw Download) DownloadSection(i int, s [2]int) {

	//download the section
	//setting the bytes range and do request
	req := dw.NewHTTPRequest("GET")
	req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", s[0], s[1]))
	resp, err := http.DefaultClient.Do(req)
	h.CheckErr(err)
	h.CheckResponseStatusCode(resp.StatusCode)

	//downloadedSize := resp.Header.Get("Content-Length")
	body, err := ioutil.ReadAll(resp.Body)
	h.CheckErr(err)

	//writing body in a tmp file
	err = ioutil.WriteFile(fmt.Sprintf("./%v/section-%v.tmp", dw.SectionsFolder, i), body, os.ModePerm)
	h.CheckErr(err)
}

func (dw Download) MergeTmpFiles(sections [][2]int) {

	//if file exists, remove it
	if h.FileExists(dw.PathToSave) {
		os.Remove(dw.PathToSave)
	}

	fd, err := os.OpenFile(dw.PathToSave, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	h.CheckErr(err)
	defer fd.Close()

	//reading sections tmp files and writing their contents on a only file
	for i := range sections {
		tmpFilename := fmt.Sprintf("./%v/section-%v.tmp", dw.SectionsFolder, i)
		fileContent, err := ioutil.ReadFile(tmpFilename)
		h.CheckErr(err)

		//writing in the final file
		_, err = fd.Write(fileContent)
		h.CheckErr(err)

		//removing tmp file
		err = os.Remove(tmpFilename)
		h.CheckErr(err)
	}

	err = os.Remove(dw.SectionsFolder)
	h.CheckErr(err)
}
