package templates

import "fmt"

type TemplateEngine interface {
	Render(templateID string, data map[string]any) (string, error)
}

type SmapleTemplateEngine struct{}

func (s *SmapleTemplateEngine) Render(templateID string, data map[string]any) (string, error) {
	return fmt.Sprintf("TemplateID %s, Data %s \n", templateID, data), nil
}
