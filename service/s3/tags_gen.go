// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	tftags "github.com/hashicorp/terraform-provider-aws/tags"
	"github.com/hashicorp/terraform-provider-aws/types"
)

// []*SERVICE.Tag handling

// Tags returns s3 service tags.
func Tags(tags tftags.KeyValueTags) []*s3.Tag {
	result := make([]*s3.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &s3.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from s3 service tags.
func KeyValueTags(ctx context.Context, tags []*s3.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns s3 service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*s3.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets s3 service tags in Context.
func SetTagsOut(ctx context.Context, tags []*s3.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}
