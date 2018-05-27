package templates

import "fmt"

const(
	templateDirectoryFmt = "templates/templates_html/%s"
)

func GetTemplatePath(templateName string) string{
	return fmt.Sprintf(templateDirectoryFmt, templateName)
}
