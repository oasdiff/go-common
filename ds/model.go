package ds

type Tenant struct {
	Id           string `datastore:"id"`
	Name         string `datastore:"name"`
	Email        string `datastore:"email"`
	Callback     string `datastore:"callback"`
	SlackChannel string `datastore:"slack_channel"`
	Created      int64  `datastore:"created"`
}

type Webhook struct {
	TenantId string `json:"tenant_id" datastore:"tenant_id"`
	Spec     string `json:"spec" datastore:"spec"`
	Copy     string `json:"copy" datastore:"copy"`
	Created  int64  `json:"created" datastore:"created"`
}
