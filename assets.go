package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"image/color"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/tinne26/etxt"
)

const SampleRate = 44100

var (
	fonts           *etxt.FontLibrary
	titleBackground *ebiten.Image
	theBuilding     *ebiten.Image
	buildingLight   *ebiten.Image
	portalLogo      *ebiten.Image

	background  map[string]*ebiten.Image
	textProfile map[string]*TextProfile

	audioContext = audio.NewContext(SampleRate)

	auntJosLetter *audio.Player
	loop0         *audio.Player
	loop1         *audio.Player
	loop2         *audio.Player

	voiceVolume = 0.6
	musicVolume = 0.75
)

func loadAssets() {
	loadImages()
	loadFonts()
	loadTextProfiles()
	loadAudio()
}

func loadImages() {
	log.Printf("Loading images...")
	titleBackground = loadImage(FileSystem, "images/title-background.png")
	theBuilding = loadImage(FileSystem, "images/building-rect.png")
	buildingLight = loadImage(FileSystem, "images/building-rect-big-lightened.png")
	portalLogo = loadImage(FileSystem, "images/building-square-200px.png")

	background = make(map[string]*ebiten.Image)
	background["title"] = titleBackground
}

// This is taken directly from https://github.com/mroobit/untitled-sidescroller/blob/main/helper.go#L120
func loadImage(fs embed.FS, path string) *ebiten.Image {
	rawFile, err := fs.Open(path)
	if err != nil {
		log.Fatalf("Error opening file %s: %v\n", path, err)
	}
	defer rawFile.Close()

	img, err := png.Decode(rawFile)
	if err != nil {
		log.Fatalf("Error decoding file %s: %v\n", path, err)
	}
	loadedImage := ebiten.NewImageFromImage(img)
	return loadedImage
}

func loadFonts() {
	fmt.Println("Loading fonts...")
	fonts = etxt.NewFontLibrary()
	_, _, err := fonts.ParseEmbedDirFonts("fonts", FileSystem)
	if err != nil {
		log.Fatalf("Error while loading fonts: %s", err.Error())
	}
	fonts.EachFont(func(name string, _ *etxt.Font) error {
		fmt.Println(name)
		return nil
	})
}

type TextProfile struct {
	Name   string
	Font   string
	AlignY string
	AlignX string
	Size   int
	Color  [4]uint8
}

func (g *Game) ConfigureTextRenderer() {
	fmt.Println("Configuring text renderer...")
	renderer := etxt.NewStdRenderer()
	cache := etxt.NewDefaultCache(10 * 1024 * 1024)
	renderer.SetCacheHandler(cache.NewHandler())
	g.Text = renderer
	g.SetTextProfile(textProfile["default"])
}

func (g *Game) SetTextProfile(p *TextProfile) {
	y := map[string]etxt.VertAlign{
		"YCenter":  etxt.YCenter,
		"Top":      etxt.Top,
		"Baseline": etxt.Baseline,
	}
	x := map[string]etxt.HorzAlign{
		"XCenter": etxt.XCenter,
		"Left":    etxt.Left,
		"Right":   etxt.Right,
	}

	g.Text.SetFont(fonts.GetFont(p.Font))
	g.Text.SetAlign(y[p.AlignY], x[p.AlignX])
	g.Text.SetSizePx(p.Size)
	g.Text.SetColor(color.RGBA{p.Color[0], p.Color[1], p.Color[2], p.Color[3]})
}

func loadTextProfiles() {
	fmt.Println("Loading text profiles...")
	var rawProfiles []*TextProfile
	profileData, err := FileSystem.ReadFile("data/text-profiles.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(profileData, &rawProfiles)
	if err != nil {
		log.Fatal("Error when unmarshalling: ", err)
	}

	textProfile = make(map[string]*TextProfile)

	for _, p := range rawProfiles {
		textProfile[p.Name] = p
	}
}

func loadAudio() {
	// TODO: load each audio file into its own audioPlayer
	auntJosLetter = loadAudioPlayer("audio/aunt-jos-letter.ogg", "single-play")
	loop0 = loadAudioPlayer("audio/loop-0.ogg", "loop")
	loop1 = loadAudioPlayer("audio/loop-1.ogg", "loop")
	loop2 = loadAudioPlayer("audio/loop-2.ogg", "loop")

}

func loadAudioPlayer(path string, pType string) *audio.Player {
	audioRaw, err := FileSystem.ReadFile(path)
	if err != nil {
		log.Fatalf("Error opening file %s: %v\n", path, err)
	}
	audioBytes, err := vorbis.DecodeWithSampleRate(SampleRate, bytes.NewReader(audioRaw))
	if err != nil {
		log.Fatalf("Error creating audio stream for %s: %v\n", path, err)
	}
	if pType == "single-play" {
		player, err := audio.CurrentContext().NewPlayer(audioBytes)
		if err != nil {
			log.Fatalf("Error creating new audio player for %s: %v\n", path, err)
		}
		player.SetVolume(voiceVolume)
		return player
	} else {
		infiniteReader := audio.NewInfiniteLoop(audioBytes, audioBytes.Length())
		player, err := audioContext.NewPlayer(infiniteReader)
		if err != nil {
			log.Fatalf("Error while creating infinite player: %s: %v\n", path, err)
		}
		player.SetVolume(musicVolume)
		return player

	}
}
