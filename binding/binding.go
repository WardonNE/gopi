package binding

var (
	BindingJSON = &jsonBinding{}
	BindingXML  = &xmlBinding{}
	BindingYAML = &yamlBinding{}
	BindingTOML = &tomlBinding{}
	BindingURI  = &uriBinding{}
	BindingFORM = &formBinding{}
)

type Binding interface {
	Parser() Parser
}
