package ds

type Tenant struct {
	Id           string `datastore:"id"`
	Name         string `datastore:"name"`
	Email        string `datastore:"email"`
	Callback     string `datastore:"callback"`
	SlackChannel string `datastore:"slack_channel"`
	Created      int64  `datastore:"created"`
	Comment      int64  `datastore:"comment"`
}

type Webhook struct {
	Id       string `datastore:"id"`
	Name     string `datastore:"name"`
	TenantId string `datastore:"tenant_id"`
	Spec     string `datastore:"spec"`
	Copy     string `datastore:"copy"`
	Created  int64  `datastore:"created"`
	Updated  int64  `datastore:"updated"`
	Owner    string `datastore:"owner"`
	Repo     string `datastore:"repo"`
	Path     string `datastore:"path"`
	Branch   string `datastore:"branch"` // revision branch to compare the base spec
}
