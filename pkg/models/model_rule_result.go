/*
 * Unified Policy Engine I/O Formats
 *
 * Documentation for the input and output formats used in Unified Policy Engine
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

// A single rule result
type RuleResult struct {
	// Whether or not this is a passing or failing result
	Passed bool `json:"passed"`
	// Whether or not this result is ignored
	Ignored bool `json:"ignored"`
	// An optional message that can be returned by a rule
	Message string `json:"message,omitempty"`
	// The ID of the primary resource (if any) associated with this result
	ResourceId string `json:"resource_id,omitempty"`
	// The type of resource (if any) associated with this result. This will typically be used with \"missing resource\" rules.
	ResourceType string `json:"resource_type,omitempty"`
	// A Markdown-formatted set of remediation steps to resolve the issue identified by the rule
	Remediation string `json:"remediation,omitempty"`
	// The severity of this rule result
	Severity string `json:"severity,omitempty"`
	// An arbitrary key-value map that a rule can return in its result.
	Context map[string]interface{} `json:"context,omitempty"`
	// A map of resource ID to a resource object associated with this result.
	Resources map[string]RuleResultResource `json:"resources,omitempty"`
}
