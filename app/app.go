package app

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/Ex0dIa-dev/yt-downloader-v2/downloader"
	fff "github.com/Ex0dIa-dev/yt-downloader-v2/ffmpeg-functions"
	h "github.com/Ex0dIa-dev/yt-downloader-v2/helper"
)

type ToDownload struct {
	YoutubeUrl   string
	YoutubeTitle string

	Video downloader.Download
	Audio downloader.Download

	Format         string
	OutputFilename string

	AudioTMPFilename string
}

func (obj ToDownload) DownloadWBEM(wg *sync.WaitGroup) {

	defer wg.Done()
	obj.Audio.PathToSave = obj.OutputFilename
	obj.Audio.DownloadStart()
}

func (obj ToDownload) DownloadMP3(wg *sync.WaitGroup) {
	defer wg.Done()

	obj.Audio.DownloadStart()
	fmt.Println("[+]Converting WEBM to MP3...")
	fff.MPEGtoMP3(obj.Audio.PathToSave, obj.AudioTMPFilename)

}

func (obj ToDownload) DownloadMP4(wg *sync.WaitGroup) {
	defer wg.Done()

	obj.Video.DownloadStart()

}

//return a []string which contains the download urls
func (obj ToDownload) GetDirectURLs() []string {

	out, err := exec.Command("youtube-dl", "-g", obj.YoutubeUrl).Output()
	h.CheckErr(err)

	urls := strings.Split(string(out), "\n")

	//urls[0] = video link, urls[1] = audio link
	return urls
}

//return the youtube title as string
func (obj ToDownload) GetYoutubeTitle() string {

	out, err := exec.Command("youtube-dl", "-e", obj.YoutubeUrl).Output()
	h.CheckErr(err)

	title := string(out)
	title = strings.Replace(title, "\n", "", -1)
	return title
}
