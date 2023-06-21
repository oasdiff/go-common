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
	Name     string `datastore:"name"`
	TenantId string `datastore:"tenant_id"`
	Spec     string `datastore:"spec"`
	Copy     string `datastore:"copy"`
	Created  int64  `datastore:"created"`
}
