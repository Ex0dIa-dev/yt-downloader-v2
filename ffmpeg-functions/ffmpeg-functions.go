package ffmpegfunctions

import (
	"os"
	"os/exec"

	h "github.com/Ex0dIa-dev/yt-downloader-v2/helper"
)

func MPEGtoMP3(input, output string) {

	cmd := exec.Command("ffmpeg", "-i", input, "-vn", "-ab", "128k", "-ar", "44100", "-y", output)
	err := cmd.Run()
	h.CheckErr(err)

	err = os.Remove(input)
	h.CheckErr(err)

}

func MergeMP4MP3(inputMP4, inputMP3, output string) {

	if h.FileExists(output) {
		err := os.Remove(output)
		h.CheckErr(err)
	}

	cmd := exec.Command("ffmpeg", "-i", inputMP4, "-i", inputMP3, "-map", "0:v", "-map", "1:a", "-c:v", "copy", "-c:a", "copy", "-y", output)
	err := cmd.Run()
	h.CheckErr(err)

	err = os.Remove(inputMP4)
	h.CheckErr(err)
	err = os.Remove(inputMP3)
	h.CheckErr(err)
}
