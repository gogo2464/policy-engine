package input

import (
	"github.com/snyk/policy-engine/pkg/models"
)

// Annotate a report with source location information
func AnnotateResults(
	configurations Loader,
	results *models.Results,
) {
	for _, inputResult := range results.Results {
		// Retrieve the filepath of the input state by looking in the metadata.
		filepath, haveFilepath := inputResult.Input.Meta["filepath"].(string)
		if !haveFilepath {
			continue
		}

		// Annotate resources in input state
		for _, resources := range inputResult.Input.Resources {
			for rk, resource := range resources {
				location := annotateResourceLocation(
					configurations,
					filepath,
					resource.Namespace,
					resource.ResourceType,
					resource.Id,
				)
				if resource.Meta == nil {
					resource.Meta = map[string]interface{}{}
				}
				if len(location) > 0 {
					resource.Meta["location"] = location
				}
				resources[rk] = resource
			}
		}

		// Annotate resources in rule results.
		for _, ruleResult := range inputResult.RuleResults {
			for _, result := range ruleResult.Results {
				annotateRuleResult(
					configurations,
					filepath,
					result,
				)
			}
		}
	}
}

func annotateRuleResult(
	configurations Loader,
	filepath string,
	result models.RuleResult,
) {
	for _, resource := range result.Resources {
		location := annotateResourceLocation(
			configurations,
			filepath,
			resource.Namespace,
			resource.Type,
			resource.Id,
		)
		resource.Location = location

		for i := range resource.Attributes {
			attributePath := []interface{}{resource.Type, resource.Id}
			attributePath = append(attributePath, resource.Attributes[i].Path...)
			location, _ := configurations.Location(filepath, attributePath)
			if len(location) > 0 {
				loc := toLocation(location[0])
				resource.Attributes[i].Location = &loc
			}
		}
	}
}

func annotateResourceLocation(
	configurations Loader,
	filepath string,
	resourceNamespace string,
	resourceType string,
	resourceId string,
) []models.SourceLocation {
	resourcePath := []interface{}{resourceNamespace, resourceType, resourceId}
	resourceLocs, _ := configurations.Location(filepath, resourcePath)
	if resourceLocs == nil {
		return nil
	}
	locations := []models.SourceLocation{}
	for _, loc := range resourceLocs {
		locations = append(locations, toLocation(loc))
	}
	return locations
}

func toLocation(loc Location) models.SourceLocation {
	return models.SourceLocation{
		Filepath: loc.Path,
		Line:     loc.Line,
		Column:   loc.Col,
	}
}