// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package youtubeplaylistnative

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type Playlist struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput   `pulumi:"description"`
	Id          pulumi.StringPtrOutput   `pulumi:"id"`
	Items       pulumi.StringArrayOutput `pulumi:"items"`
	Title       pulumi.StringPtrOutput   `pulumi:"title"`
}

// NewPlaylist registers a new resource with the given unique name, arguments, and options.
func NewPlaylist(ctx *pulumi.Context,
	name string, args *PlaylistArgs, opts ...pulumi.ResourceOption) (*Playlist, error) {
	if args == nil {
		args = &PlaylistArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Playlist
	err := ctx.RegisterResource("youtubeplaylist-native:index:Playlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlaylist gets an existing Playlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlaylist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlaylistState, opts ...pulumi.ResourceOption) (*Playlist, error) {
	var resource Playlist
	err := ctx.ReadResource("youtubeplaylist-native:index:Playlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Playlist resources.
type playlistState struct {
}

type PlaylistState struct {
}

func (PlaylistState) ElementType() reflect.Type {
	return reflect.TypeOf((*playlistState)(nil)).Elem()
}

type playlistArgs struct {
	Description *string  `pulumi:"description"`
	Id          *string  `pulumi:"id"`
	Items       []string `pulumi:"items"`
	Title       *string  `pulumi:"title"`
}

// The set of arguments for constructing a Playlist resource.
type PlaylistArgs struct {
	Description pulumi.StringPtrInput
	Id          pulumi.StringPtrInput
	Items       pulumi.StringArrayInput
	Title       pulumi.StringPtrInput
}

func (PlaylistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*playlistArgs)(nil)).Elem()
}

type PlaylistInput interface {
	pulumi.Input

	ToPlaylistOutput() PlaylistOutput
	ToPlaylistOutputWithContext(ctx context.Context) PlaylistOutput
}

func (*Playlist) ElementType() reflect.Type {
	return reflect.TypeOf((**Playlist)(nil)).Elem()
}

func (i *Playlist) ToPlaylistOutput() PlaylistOutput {
	return i.ToPlaylistOutputWithContext(context.Background())
}

func (i *Playlist) ToPlaylistOutputWithContext(ctx context.Context) PlaylistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlaylistOutput)
}

type PlaylistOutput struct{ *pulumi.OutputState }

func (PlaylistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Playlist)(nil)).Elem()
}

func (o PlaylistOutput) ToPlaylistOutput() PlaylistOutput {
	return o
}

func (o PlaylistOutput) ToPlaylistOutputWithContext(ctx context.Context) PlaylistOutput {
	return o
}

func (o PlaylistOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Playlist) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PlaylistOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Playlist) pulumi.StringPtrOutput { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PlaylistOutput) Items() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Playlist) pulumi.StringArrayOutput { return v.Items }).(pulumi.StringArrayOutput)
}

func (o PlaylistOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Playlist) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlaylistInput)(nil)).Elem(), &Playlist{})
	pulumi.RegisterOutputType(PlaylistOutput{})
}
