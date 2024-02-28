import * as pulumi from "@pulumi/pulumi";
import * as youtubeplaylist_native from "@pulumi/youtubeplaylist-native";

const myPlaylist = new youtubeplaylist_native.Playlist("myPlaylist", {
    title: "test-playlist",
    itemIds: [
        "wdbf7-tlDV8",
        "o-PviKVN0J4",
    ],
});
export const output = {
    value: myPlaylist.playlistId,
};
