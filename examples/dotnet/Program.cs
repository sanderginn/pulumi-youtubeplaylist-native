using System.Collections.Generic;
using System.Linq;
using Pulumi;
using YoutubeplaylistNative = Pulumi.YoutubeplaylistNative;

return await Deployment.RunAsync(() => 
{
    var myPlaylist = new YoutubeplaylistNative.Playlist("myPlaylist", new()
    {
        Title = "pulumi-test-playlist",
        ItemIds = new[]
        {
            "wdbf7-tlDV8",
            "o-PviKVN0J4",
        },
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myPlaylist.PlaylistId },
        },
    };
});

