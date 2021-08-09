/*
	Made by Ex0dIa-dev
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Ex0dIa-dev/yt-downloader-v2/app"
	fff "github.com/Ex0dIa-dev/yt-downloader-v2/ffmpeg-functions"
	h "github.com/Ex0dIa-dev/yt-downloader-v2/helper"
)

//flags
func init() {
	flag.StringVar(&obj.YoutubeUrl, "u", "", "insert url")
	flag.StringVar(&obj.Format, "f", "mp3", "insert format")
	flag.StringVar(&obj.OutputFilename, "o", "", "output filename")
}

var obj app.ToDownload

func main() {

	flag.Parse()
	if obj.YoutubeUrl == "" {
		fmt.Println("[-]Please enter a valid url.")
		os.Exit(1)
	}

	if !h.CheckSupportedFormat(obj.Format) {
		fmt.Println("[-]Please enter a supported format.")
		os.Exit(1)
	}

	startTime := time.Now()

	fmt.Println("[+]Verifying Url...")
	obj.YoutubeTitle = obj.GetYoutubeTitle()
	fmt.Println("[+]Getting Download Urls...")
	urls := obj.GetDirectURLs()

	obj.Video.URL = urls[0]
	obj.Video.PathToSave = "video.mp4"
	obj.Video.SectionsFolder = "tmp-video-folder"

	obj.Audio.URL = urls[1]
	obj.Audio.PathToSave = "audio.webm"
	obj.Audio.SectionsFolder = "tmp-audio-folder"

	if obj.OutputFilename == "" {
		obj.OutputFilename = fmt.Sprintf("%v.%v", obj.YoutubeTitle, obj.Format)
	}

	obj.AudioTMPFilename = "audio-tmp.mp3"

	switch obj.Format {
	case "mp3":
		var wg sync.WaitGroup
		wg.Add(1)
		fmt.Println("[+]Download Started...")
		obj.DownloadMP3(&wg)
		wg.Wait()

		err := os.Rename(obj.AudioTMPFilename, obj.OutputFilename)
		h.CheckErr(err)

		fmt.Printf("[+]Operation Completed in %v seconds.\n", time.Since(startTime).Seconds())

	case "mp4":
		var wg sync.WaitGroup
		wg.Add(2)

		fmt.Println("[+]Download Started...")
		go obj.DownloadMP3(&wg)
		go obj.DownloadMP4(&wg)
		wg.Wait()

		fff.MergeMP4MP3(obj.Video.PathToSave, obj.AudioTMPFilename, obj.OutputFilename)

		fmt.Printf("[+]Operation Completed in %v seconds.\n", time.Since(startTime).Seconds())

	case "webm":
		var wg sync.WaitGroup
		wg.Add(1)

		fmt.Println("[+]Download Started...")
		obj.DownloadWBEM(&wg)
		wg.Wait()
		fmt.Printf("[+]Operation Completed in %v seconds.\n", time.Since(startTime).Seconds())
	}

}
