package main

import (
	youtubeplaylistNative "github.com/pulumi/pulumi-youtubeplaylist-native/sdk/go/youtubeplaylist-native"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myPlaylist, err := youtubeplaylistNative.NewPlaylist(ctx, "myPlaylist", &youtubeplaylistNative.PlaylistArgs{
			Title: pulumi.String("test-playlist"),
			ItemIds: pulumi.StringArray{
				pulumi.String("wdbf7-tlDV8"),
				pulumi.String("o-PviKVN0J4"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myPlaylist.PlaylistId,
		})
		return nil
	})
}
