package traversal

import "fmt"

// VertexLabel is a strongly typed label for graph vertices.
// When serialized to a Gremlin query string it emits a Go template action
// that the executor resolves to a tenant-prefixed value at query time.
//
// Example:
//
//	g.V().HasLabel(VertexLabel("aws.iam.User"))
//	// serialises to: g.V().hasLabel("{{ "aws.iam.User" | tenant }}")
//	// executed as:   g.V().hasLabel("acme-aws.iam.User")
type VertexLabel string

// templateAction returns the Go template action string for this label.
func (v VertexLabel) templateAction() string {
	return fmt.Sprintf(`{{ "%s" | tenant }}`, string(v))
}

// TemplateAction returns the Go template action string for this label.
// It is the exported form of templateAction for use by the query executor.
func (v VertexLabel) TemplateAction() string {
	return v.templateAction()
}

// EdgeLabel is a strongly typed label for graph edges.
// It follows the same template-based tenant-prefix convention as VertexLabel.
type EdgeLabel string

// templateAction returns the Go template action string for this label.
func (e EdgeLabel) templateAction() string {
	return fmt.Sprintf(`{{ "%s" | tenant }}`, string(e))
}

// TemplateAction returns the Go template action string for this label.
// It is the exported form of templateAction for use by the query executor.
func (e EdgeLabel) TemplateAction() string {
	return e.templateAction()
}