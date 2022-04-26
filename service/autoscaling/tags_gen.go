// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package autoscaling

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/tags"
	"github.com/hashicorp/terraform-provider-aws/tfresource"
)

// GetTag fetches an individual autoscaling service tag for a resource.
// Returns whether the key value and any errors. A NotFoundError is used to signal that no value was found.
// This function will optimise the handling over ListTags, if possible.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func GetTag(conn *autoscaling.AutoScaling, identifier string, resourceType string, key string) (*tftags.TagData, error) {
	input := &autoscaling.DescribeTagsInput{
		Filters: []*autoscaling.Filter{
			{
				Name:   aws.String("auto-scaling-group"),
				Values: []*string{aws.String(identifier)},
			},
			{
				Name:   aws.String("key"),
				Values: []*string{aws.String(key)},
			},
		},
	}

	output, err := conn.DescribeTags(input)

	if err != nil {
		return nil, err
	}

	listTags := KeyValueTags(output.Tags, identifier, resourceType)

	if !listTags.KeyExists(key) {
		return nil, tfresource.NewEmptyResultError(nil)
	}

	return listTags.KeyTagData(key), nil
}

// ListTags lists autoscaling service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(conn *autoscaling.AutoScaling, identifier string, resourceType string) (tftags.KeyValueTags, error) {
	input := &autoscaling.DescribeTagsInput{
		Filters: []*autoscaling.Filter{
			{
				Name:   aws.String("auto-scaling-group"),
				Values: []*string{aws.String(identifier)},
			},
		},
	}

	output, err := conn.DescribeTags(input)

	if err != nil {
		return tftags.New(nil), err
	}

	return KeyValueTags(output.Tags, identifier, resourceType), nil
}

// []*SERVICE.Tag handling

// ListOfMap returns a list of autoscaling in flattened map.
//
// Compatible with setting Terraform state for strongly typed configuration blocks.
//
// This function strips tag resource identifier and type. Generally, this is
// the desired behavior so the tag schema does not require those attributes.
// Use (tftags.KeyValueTags).ListOfMap() for full tag information.
func ListOfMap(tags tftags.KeyValueTags) []interface{} {
	var result []interface{}

	for _, key := range tags.Keys() {
		m := map[string]interface{}{
			"key":   key,
			"value": aws.StringValue(tags.KeyValue(key)),

			"propagate_at_launch": aws.BoolValue(tags.KeyAdditionalBoolValue(key, "PropagateAtLaunch")),
		}

		result = append(result, m)
	}

	return result
}

// ListOfStringMap returns a list of autoscaling tags in flattened map of only string values.
//
// Compatible with setting Terraform state for legacy []map[string]string schema.
// Deprecated: Will be removed in a future major version without replacement.
func ListOfStringMap(tags tftags.KeyValueTags) []interface{} {
	var result []interface{}

	for _, key := range tags.Keys() {
		m := map[string]string{
			"key":   key,
			"value": aws.StringValue(tags.KeyValue(key)),

			"propagate_at_launch": strconv.FormatBool(aws.BoolValue(tags.KeyAdditionalBoolValue(key, "PropagateAtLaunch"))),
		}

		result = append(result, m)
	}

	return result
}

// Tags returns autoscaling service tags.
func Tags(tags tftags.KeyValueTags) []*autoscaling.Tag {
	var result []*autoscaling.Tag

	for _, key := range tags.Keys() {
		tag := &autoscaling.Tag{
			Key:               aws.String(key),
			Value:             tags.KeyValue(key),
			ResourceId:        tags.KeyAdditionalStringValue(key, "ResourceId"),
			ResourceType:      tags.KeyAdditionalStringValue(key, "ResourceType"),
			PropagateAtLaunch: tags.KeyAdditionalBoolValue(key, "PropagateAtLaunch"),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from autoscaling service tags.
//
// Accepts the following types:
//   - []*autoscaling.Tag
//   - []*autoscaling.TagDescription
//   - []interface{} (Terraform TypeList configuration block compatible)
//   - *schema.Set (Terraform TypeSet configuration block compatible)
func KeyValueTags(tags interface{}, identifier string, resourceType string) tftags.KeyValueTags {
	switch tags := tags.(type) {
	case []*autoscaling.Tag:
		m := make(map[string]*tftags.TagData, len(tags))

		for _, tag := range tags {
			tagData := &tftags.TagData{
				Value: tag.Value,
			}

			tagData.AdditionalBoolFields = make(map[string]*bool)
			tagData.AdditionalBoolFields["PropagateAtLaunch"] = tag.PropagateAtLaunch
			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			m[aws.StringValue(tag.Key)] = tagData
		}

		return tftags.New(m)
	case []*autoscaling.TagDescription:
		m := make(map[string]*tftags.TagData, len(tags))

		for _, tag := range tags {
			tagData := &tftags.TagData{
				Value: tag.Value,
			}
			tagData.AdditionalBoolFields = make(map[string]*bool)
			tagData.AdditionalBoolFields["PropagateAtLaunch"] = tag.PropagateAtLaunch
			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			m[aws.StringValue(tag.Key)] = tagData
		}

		return tftags.New(m)
	case *schema.Set:
		return KeyValueTags(tags.List(), identifier, resourceType)
	case []interface{}:
		result := make(map[string]*tftags.TagData)

		for _, tfMapRaw := range tags {
			tfMap, ok := tfMapRaw.(map[string]interface{})

			if !ok {
				continue
			}

			key, ok := tfMap["key"].(string)

			if !ok {
				continue
			}

			tagData := &tftags.TagData{}

			if v, ok := tfMap["value"].(string); ok {
				tagData.Value = &v
			}

			tagData.AdditionalBoolFields = make(map[string]*bool)
			if v, ok := tfMap["propagate_at_launch"].(bool); ok {
				tagData.AdditionalBoolFields["PropagateAtLaunch"] = &v
			}

			// Deprecated: Legacy map handling
			if v, ok := tfMap["propagate_at_launch"].(string); ok {
				b, _ := strconv.ParseBool(v)
				tagData.AdditionalBoolFields["PropagateAtLaunch"] = &b
			}

			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			result[key] = tagData
		}

		return tftags.New(result)
	default:
		return tftags.New(nil)
	}
}

// UpdateTags updates autoscaling service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(conn *autoscaling.AutoScaling, identifier string, resourceType string, oldTagsSet interface{}, newTagsSet interface{}) error {
	oldTags := KeyValueTags(oldTagsSet, identifier, resourceType)
	newTags := KeyValueTags(newTagsSet, identifier, resourceType)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &autoscaling.DeleteTagsInput{
			Tags: Tags(removedTags.IgnoreAWS()),
		}

		_, err := conn.DeleteTags(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &autoscaling.CreateOrUpdateTagsInput{
			Tags: Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.CreateOrUpdateTags(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
