package shared

type Security struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}
