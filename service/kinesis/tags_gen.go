// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package kinesis

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/hashicorp/terraform-provider-aws/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/tags"
	"github.com/hashicorp/terraform-provider-aws/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// ListTags lists kinesis service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn kinesisiface.KinesisAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &kinesis.ListTagsForStreamInput{
		StreamName: aws.String(identifier),
	}

	output, err := conn.ListTagsForStreamWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists kinesis service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).KinesisConn(), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns kinesis service tags.
func Tags(tags tftags.KeyValueTags) []*kinesis.Tag {
	result := make([]*kinesis.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &kinesis.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from kinesis service tags.
func KeyValueTags(ctx context.Context, tags []*kinesis.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns kinesis service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*kinesis.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets kinesis service tags in Context.
func SetTagsOut(ctx context.Context, tags []*kinesis.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// createTags creates kinesis service tags for new resources.
func createTags(ctx context.Context, conn kinesisiface.KinesisAPI, identifier string, tags []*kinesis.Tag) error {
	if len(tags) == 0 {
		return nil
	}

	return UpdateTags(ctx, conn, identifier, nil, KeyValueTags(ctx, tags))
}

// UpdateTags updates kinesis service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(ctx context.Context, conn kinesisiface.KinesisAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.Kinesis)
	if len(removedTags) > 0 {
		for _, removedTags := range removedTags.Chunks(10) {
			input := &kinesis.RemoveTagsFromStreamInput{
				StreamName: aws.String(identifier),
				TagKeys:    aws.StringSlice(removedTags.Keys()),
			}

			_, err := conn.RemoveTagsFromStreamWithContext(ctx, input)

			if err != nil {
				return fmt.Errorf("untagging resource (%s): %w", identifier, err)
			}
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.Kinesis)
	if len(updatedTags) > 0 {
		for _, updatedTags := range updatedTags.Chunks(10) {
			input := &kinesis.AddTagsToStreamInput{
				StreamName: aws.String(identifier),
				Tags:       aws.StringMap(updatedTags.IgnoreAWS().Map()),
			}

			_, err := conn.AddTagsToStreamWithContext(ctx, input)

			if err != nil {
				return fmt.Errorf("tagging resource (%s): %w", identifier, err)
			}
		}
	}

	return nil
}

// UpdateTags updates kinesis service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).KinesisConn(), identifier, oldTags, newTags)
}
