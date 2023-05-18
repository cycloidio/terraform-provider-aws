// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package swf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
	"github.com/hashicorp/terraform-provider-aws/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/tags"
	"github.com/hashicorp/terraform-provider-aws/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// ListTags lists swf service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn swfiface.SWFAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &swf.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists swf service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).SWFConn(), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns swf service tags.
func Tags(tags tftags.KeyValueTags) []*swf.ResourceTag {
	result := make([]*swf.ResourceTag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &swf.ResourceTag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from swf service tags.
func KeyValueTags(ctx context.Context, tags []*swf.ResourceTag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns swf service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*swf.ResourceTag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets swf service tags in Context.
func SetTagsOut(ctx context.Context, tags []*swf.ResourceTag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// UpdateTags updates swf service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(ctx context.Context, conn swfiface.SWFAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.SWF)
	if len(removedTags) > 0 {
		input := &swf.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.SWF)
	if len(updatedTags) > 0 {
		input := &swf.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates swf service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).SWFConn(), identifier, oldTags, newTags)
}
