package traversal

import (
	"strings"
	"testing"
)

func TestVertexLabel_SerializesAsTemplateAction(t *testing.T) {
	g := NewTraversal().V().HasLabel(VertexLabel("aws.iam.User"))
	query := g.String()

	if !strings.Contains(query, `{{ "aws.iam.User" | tenant }}`) {
		t.Fatalf("expected template action in query, got: %s", query)
	}
	if strings.Contains(query, `hasLabel("aws.iam.User")`) {
		t.Fatalf("raw label must not appear in query without template action: %s", query)
	}
}

func TestEdgeLabel_AddE_SerializesAsTemplateAction(t *testing.T) {
	g := NewTraversal().AddE(EdgeLabel("BELONGS_TO"))
	query := g.String()

	if !strings.Contains(query, `{{ "BELONGS_TO" | tenant }}`) {
		t.Fatalf("expected template action in query, got: %s", query)
	}
}

func TestVertexLabel_AddV_SerializesAsTemplateAction(t *testing.T) {
	g := NewTraversal().AddV(VertexLabel("aws.iam.User"))
	query := g.String()

	if !strings.Contains(query, `{{ "aws.iam.User" | tenant }}`) {
		t.Fatalf("expected template action in query, got: %s", query)
	}
}

func TestEdgeLabel_Out_SerializesAsTemplateAction(t *testing.T) {
	g := NewTraversal().V().Out(EdgeLabel("BELONGS_TO"))
	query := g.String()

	if !strings.Contains(query, `{{ "BELONGS_TO" | tenant }}`) {
		t.Fatalf("expected template action in query, got: %s", query)
	}
}

func TestEdgeLabel_In_SerializesAsTemplateAction(t *testing.T) {
	g := NewTraversal().V().In(EdgeLabel("BELONGS_TO"))
	query := g.String()

	if !strings.Contains(query, `{{ "BELONGS_TO" | tenant }}`) {
		t.Fatalf("expected template action in query, got: %s", query)
	}
}

func TestEdgeLabel_Both_SerializesAsTemplateAction(t *testing.T) {
	g := NewTraversal().V().Both(EdgeLabel("BELONGS_TO"))
	query := g.String()

	if !strings.Contains(query, `{{ "BELONGS_TO" | tenant }}`) {
		t.Fatalf("expected template action in query, got: %s", query)
	}
}

func TestString_Out_BackwardCompatible(t *testing.T) {
	g := NewTraversal().V().Out("BELONGS_TO")
	query := g.String()

	if !strings.Contains(query, `"BELONGS_TO"`) {
		t.Fatalf("plain string label should appear quoted in query, got: %s", query)
	}
	if strings.Contains(query, "tenant") {
		t.Fatalf("plain string label must not emit a template action, got: %s", query)
	}
}

func TestHasLabel_BackwardCompatible(t *testing.T) {
	g := NewTraversal().V().HasLabel("aws.iam.User")
	query := g.String()

	if !strings.Contains(query, `"aws.iam.User"`) {
		t.Fatalf("plain string label should appear quoted in query, got: %s", query)
	}
	if strings.Contains(query, "tenant") {
		t.Fatalf("plain string label must not emit a template action, got: %s", query)
	}
}
