# SongMan
Spotify playlist/track information to json, mp4 and mp3 converter, written in GoLang.

## How it works

**! FFMPEG is required for this project to work !**

1. SongMan will take playlist or track link from spotify and make sure its a valid link.

2. Content will be turned into blueprint and exported in `blueprints` directory with `playlist/track` name and `.json` file extension

3. If specified, blueprints will be used to find the corresponding video on Youtube and download it in `videos` directory with `.mp4` file extension.

4. If downloaded, mp3 will be extracted from mp4 videos in `music` direcotry with `.mp3` file extension.

## Usage
1. **Clone the project:**
    ```bash
    git clone https://github.com/FinnTheHero/SongMan.git && cd Songman
    ```
2. **Create `.env` file:**
    ```bash
    mkdir .env
    ```
3. **Put your Spotify client id & secret and 2 of Youtube Data API keys in `.env` file:**
    ```
    CLIENT_ID=""
    CLIENT_SECRET=""
    YOUTUBE_API_KEY_1=""
    YOUTUBE_API_KEY_2=""
    ```
    > 2 API keys for youtube is optional, you can use only one if you want so. However, you have to change it on your own in the code `src/download/youtube.go` when calling function `GetVideo()`.
4. **Run the project:**
    ```bash
    go run . -link "https://open.spotify.com/playlist/3cEYpjA9oz9GiPac4AsH4n"
    ```

    Add `-download true` option to attempt downloading. However, this requires `ffmpeg` to be installed on the device. This method will also extract audio as an mp3 from downloaded videos.

    ```bash
    go run . -link "https://open.spotify.com/playlist/3cEYpjA9oz9GiPac4AsH4n" -download true
    ```

---

**Author takes no responsibility of how you use this project!**