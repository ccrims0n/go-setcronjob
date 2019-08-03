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

func (g *groups) List() (*[]Group, error) {

}

func (g *groups) Get(id int) (*Group, error) {

}

func (g *groups) Lookup(name string) (*Group, error) {

}

func (g *groups) Add(group *Group) (*Group, error) {

}

func (g *groups) Edit(id int, group *Group) (*Group, error) {

}

func (g *groups) Delete(id int) (*Group, error) {

}

func (g *groups) Vanish(id int) (*Group, error) {

}

func (g *groups) Empty(id int) (*Group, error) {

}

func (g *groups) Items(id int) (*[]Cron, error) {

}
