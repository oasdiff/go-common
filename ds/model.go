package ds

type Tenant struct {
	Id           string `datastore:"id" json:"id"`
	Name         string `datastore:"name" json:"name"`
	Email        string `datastore:"email" json:"email"`
	Callback     string `datastore:"callback" json:"callback"`
	SlackChannel string `datastore:"slack_channel" json:"slack_channel"`
	Created      int64  `datastore:"created" json:"created"`
	Comment      int64  `datastore:"comment" json:"comment"`
}

type Webhook struct {
	Id        string `datastore:"id" json:"id"`
	Name      string `datastore:"name" json:"name"`
	TenantId  string `datastore:"tenant_id" json:"tenant_id"`
	Spec      string `datastore:"spec" json:"spec"`
	Copy      string `datastore:"copy" json:"copy"`
	Created   int64  `datastore:"created" json:"created"`
	Updated   int64  `datastore:"updated" json:"updated"`
	Owner     string `datastore:"owner" json:"owner"`
	Repo      string `datastore:"repo" json:"repo"`
	Path      string `datastore:"path" json:"path"`
	Branch    string `datastore:"branch" json:"branch"` // revision branch to compare the base spec
	Changelog string `datastore:"changelog" json:"changelog"`
}
