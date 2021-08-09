### yt-downloader-v2

****

**Dependencies:**

- [youtube-dl] https://github.com/ytdl-org/youtube-dl)

- [ffmpeg](https://ffmpeg.org/)

****

**Build:**

```bash
go build
```

**Use:**

```bash
./yt-downloader-v2 -u 'url'
```

```bash
./yt-downloader-v2 -u 'url' -f <format>
```

```bash
./yt-downloader-v2 -u 'url' -f <format> -o 'output.mp3'
```

**Example:**

```bash
./yt-downloader-v2 -u 'https://www.youtube.com/watch?v=XbGs_qK2PQA'
```

```bash
./yt-downloader-v2 -u 'https://www.youtube.com/watch?v=XbGs_qK2PQA' -f mp4
```

```bash
./yt-downloader-v2 -u 'https://www.youtube.com/watch?v=XbGs_qK2PQA' -f mp4 -o 'rapgod.mp4'
```

****

**Flags:**

- '**-u**' ==> enter youtube video url
- '**-f**' ==> enter file format(default: mp3)
- '**-o**' ==> enter output filename(ex. 'song.mp3' or 'video.mp4')

**Supported Formats:**

- **MP3**
- **MP4**
- **WEBM**

****

**Notes**

New supported formats will be added next...
