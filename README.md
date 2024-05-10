# SongMan
Spotify playlist/track information to json blueprint converter, written in GoLang.

## How it works
1. SongMan will take playlist or track link from spotify and make sure its a valid link.

2. Content will be turned into blueprint and exported in `blueprints` directory with `playlist/track` name and `.json` file extension

You can deal with blueprints however you want, download or use for research.

## Usage
1. **Clone the project:**
    ```bash
    git clone https://github.com/FinnTheHero/SongMan.git && cd Songman
    ```
2. **Create `.env` file:**
    ```bash
    mkdir .env
    ```
3. **Put your spotify app client id and client secret in `.env` file:**
    ```
    CLIENT_ID=""
    CLIENT_SECRET=""
    ```
4. **Run the project:**
    ```bash
    go run . <Mode> "<Link>"
    ```
    **Modes:**
    * **`-track`**: Create blueprint only for a track
    * **`-playlist`**: Create blueprint only for a playlist

    **Example:**
    ```bash
    go run . -playlist "https://open.spotify.com/playlist/3cEYpjA9oz9GiPac4AsH4n"
    ```

---

**Author takes no responsibility of how you use this project!**