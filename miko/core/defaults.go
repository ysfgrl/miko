package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko/assets"
	"image"
	"path"
	"path/filepath"
)

var Loading []*ebiten.Image

func init() {
	Loading, _ = LoadAssets("images/default/loading")
}

func LoadAssets(folder string) ([]*ebiten.Image, error) {

	res := []*ebiten.Image{}
	ents, err := assets.Assets.ReadDir(folder)
	if err != nil {
		return res, err
	}
	for _, ent := range ents {
		name := ent.Name()
		ext := filepath.Ext(name)
		if ext != ".png" {
			continue
		}
		f, err := assets.Assets.Open(path.Join(folder, name))
		if err != nil {
			return res, err
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return res, err
		}
		res = append(res, ebiten.NewImageFromImage(img))
	}
	return res, nil
}
