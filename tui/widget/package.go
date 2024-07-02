package widget

import (
	"github.com/bbfh-dev/browser.mcvm/pkg/mcvm"
	"github.com/tidwall/gjson"
)

type PackageWidget struct {
	Id           string
	Ready        bool
	metadataJSON string
}

func NewPackageWidget(id string) PackageWidget {
	return PackageWidget{
		Id:           id,
		metadataJSON: "",
	}
}

func (widget *PackageWidget) Load() {
	metadata, err := mcvm.PackageMetadata(widget.Id)
	widget.metadataJSON = metadata
	widget.Ready = err == nil && gjson.Valid(widget.metadataJSON)
}

func (widget PackageWidget) Name() string {
	return gjson.Get(widget.metadataJSON, "metadata.name").String()
}

func (widget PackageWidget) Description() string {
	return gjson.Get(widget.metadataJSON, "metadata.description").String()
}

func (widget PackageWidget) Valid() bool {
	return widget.Ready
}
