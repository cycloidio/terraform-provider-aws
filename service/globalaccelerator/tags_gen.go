// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package globalaccelerator

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	tftags "github.com/hashicorp/terraform-provider-aws/tags"
)

// ListTags lists globalaccelerator service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(conn *globalaccelerator.GlobalAccelerator, identifier string) (tftags.KeyValueTags, error) {
	input := &globalaccelerator.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(input)

	if err != nil {
		return tftags.New(nil), err
	}

	return KeyValueTags(output.Tags), nil
}

// []*SERVICE.Tag handling

// Tags returns globalaccelerator service tags.
func Tags(tags tftags.KeyValueTags) []*globalaccelerator.Tag {
	result := make([]*globalaccelerator.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &globalaccelerator.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from globalaccelerator service tags.
func KeyValueTags(tags []*globalaccelerator.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(m)
}

// UpdateTags updates globalaccelerator service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(conn *globalaccelerator.GlobalAccelerator, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(oldTagsMap)
	newTags := tftags.New(newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &globalaccelerator.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     aws.StringSlice(removedTags.IgnoreAWS().Keys()),
		}

		_, err := conn.UntagResource(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &globalaccelerator.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.TagResource(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
