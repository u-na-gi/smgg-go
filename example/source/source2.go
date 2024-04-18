package source

import "github.com/u-na-gi/smgg-go/example/types"

type Bird struct {
	Name    types.BirdName
	Size    int
	singing bool
	dancing bool
	Wings   types.WingsInterface
}
