package quiltdl

import (
	"encoding/json"
	"errors"
	"github.com/pufferpanel/pufferpanel/v2"
	"github.com/pufferpanel/pufferpanel/v2/environments"
	"path"
)

const QuiltMetadataUrl = "https://meta.quiltmc.org/v3/versions/installer"

type Quiltdl struct {
	TargetFile string
}

type QuiltMetadata struct {
	Url string `json:"url"`
}

func (f *Quiltdl) Run(env pufferpanel.Environment) error {
	env.DisplayToConsole(true, "Downloading metadata from %s\n", QuiltMetadataUrl)
	response, err := pufferpanel.HttpGet(QuiltMetadataUrl)
	if err != nil {
		return err
	}
	defer pufferpanel.Close(response.Body)

	var metadata []QuiltMetadata
	err = json.NewDecoder(response.Body).Decode(&metadata)
	if err != nil {
		return err
	}
	if len(metadata) == 0 {
		return errors.New("No metadata available from Quilt, unable to download installer")
	}

	file, err := environments.DownloadViaMaven(metadata[0].Url, env)
	if err != nil {
		return err
	}

	err = pufferpanel.CopyFile(file, path.Join(env.GetRootDirectory(), "quilt-installer.jar"))
	if err != nil {
		return err
	}

	return nil
}
