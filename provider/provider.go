// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"errors"
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"google.golang.org/api/youtube/v3"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "youtubeplaylist-native"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[Playlist, PlaylistArgs, PlaylistState](),
		},
	})
}

type Playlist struct{}

type PlaylistArgs struct {
	// todo: if Id is left empty, Title must be provided in order to create the playlist
	Title       string `pulumi:"title,optional"`
	Description string `pulumi:"description,optional"`
	Id          string `pulumi:"id,optional"`
}

type PlaylistState struct {
	PlaylistArgs
}

func (Playlist) Create(ctx p.Context, name string, input PlaylistArgs, preview bool) (string, PlaylistState, error) {
	youtubeService, err := youtube.NewService(ctx)
	if err != nil {
		return "", PlaylistState{}, errors.New(fmt.Sprintf("failed to create youtube service: %v", err))
	}
	state := PlaylistState{PlaylistArgs: input}
	if preview {
		return name, state, nil
	}

	// If the user didn't provide an ID, we'll create a new playlist.
	if input.Id == "" {
		id, err := createPlaylist(youtubeService, input)
		if err != nil {
			return "", PlaylistState{}, errors.New(fmt.Sprintf("failed to create playlist: %v", err))
		}
		state.Id = id
	} else {
		playlist, err := getPlaylist(youtubeService, input.Id)
		if err != nil {
			return "", PlaylistState{}, errors.New(fmt.Sprintf("failed to get playlist: %v", err))
		}
		state.Title = playlist.Snippet.Title
		state.Description = playlist.Snippet.Description
	}
	return name, state, nil
}

func getPlaylist(youtubeService *youtube.Service, id string) (*youtube.Playlist, error) {
	call := youtubeService.Playlists.List([]string{"snippet"}).Id(id)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	if len(response.Items) == 0 {
		return nil, errors.New(fmt.Sprintf("playlist with id %s not found", id))
	}
	return response.Items[0], nil
}

func createPlaylist(youtubeService *youtube.Service, input PlaylistArgs) (string, error) {
	playlist := &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title:       input.Title,
			Description: input.Description,
		},
	}
	call := youtubeService.Playlists.Insert([]string{"snippet"}, playlist)
	response, err := call.Do()
	if err != nil {
		return "", err
	}
	return response.Id, nil
}
