package setcronjob

type Group struct {
	Id 	 int 	`json:"id"`
	Name string `json:"name"`
}

type Groups interface {
	List() (*[]Group, error)
	Get(id int) (*Group, error)
	Lookup(name string) (*Group, error)
	Add(group *Group) (*Group, error)
	Edit(id int, group *Group) (*Group, error)
	Delete(id int) (*Group, error)
	Vanish(id int) (*Group, error)
	Empty(id int) (*Group, error)
	Items(id int) (*[]Cron, error)
}

type groups struct {
	client Client
}

func newGroups(c Client) Groups {
	return &groups{
		client: c,
	}
}

