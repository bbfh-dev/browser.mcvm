package widget

import (
	"slices"
	"strings"

	"github.com/bbfh-dev/browser.mcvm/pkg/mcvm"
	"github.com/bbfh-dev/browser.mcvm/tui/style"
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

func (widget PackageWidget) Hint() string {
	return widget.Id
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

func (widget PackageWidget) Extra() string {
	entries := []string{
		style.WithIcon(
			" ",
			strings.ReplaceAll(strings.Join(widget.Modloaders(), ", "), "like", ""),
		),
	}

	categories := widget.Categories()
	if len(categories) > 0 {
		entries = append(entries, strings.Join(categories, ", "))
	}

	if slices.Contains(widget.SupportedSides(), "server") {
		entries = append(entries, style.IconFallback(" ", "Server"))
	}

	if slices.Contains(widget.SupportedSides(), "client") {
		entries = append(entries, style.IconFallback("󰍹 ", "Client"))
	}

	authors := widget.Authors()
	if len(categories) > 0 {
		entries = append(entries, style.WithIcon(" ", strings.Join(authors, ", ")))
	}

	return strings.Join(entries, " 󰧞 ")
}

func getJSONStringArray(source string, query string) (entries []string) {
	results := gjson.Get(source, query).Array()
	for _, result := range results {
		entries = append(entries, result.String())
	}

	return entries
}

func (widget PackageWidget) Authors() []string {
	return getJSONStringArray(widget.metadataJSON, "metadata.authors")
}

func (widget PackageWidget) Maintainers() []string {
	return getJSONStringArray(widget.metadataJSON, "metadata.package_maintainers")
}

func (widget PackageWidget) Categories() []string {
	return getJSONStringArray(widget.metadataJSON, "metadata.categories")
}

func (widget PackageWidget) Modloaders() []string {
	return getJSONStringArray(widget.metadataJSON, "properties.supported_modloaders")
}

func (widget PackageWidget) SupportedSides() []string {
	return getJSONStringArray(widget.metadataJSON, "properties.supported_sides")
}
