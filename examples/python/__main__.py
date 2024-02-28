import pulumi
import pulumi_youtubeplaylist_native as youtubeplaylist_native

my_playlist = youtubeplaylist_native.Playlist("myPlaylist",
    title="pulumi-test-playlist",
    item_ids=[
        "wdbf7-tlDV8",
        "o-PviKVN0J4",
    ])
pulumi.export("output", {
    "value": my_playlist.playlist_id,
})
