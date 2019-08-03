package setcronjob

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Cron struct {
	Id               int    `json:"id,omitempty"`
	Group            int    `json:"group"`
	Expression       string `json:"expression"`
	Timezone         string `json:"timezone"`
	DnsCache         bool   `json:"dnsCache"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Url              string `json:"url"`
	HttpMethod       string `json:"httpMethod"`
	HttpHeaders      string `json:"httpHeaders"`
	PostData         string `json:"postData"`
	Fail             int    `json:"fail"`
	Status           int    `json:"status"`
	Name             string `json:"name"`
	Notify           int    `json:"notify"`
	NotifyEvery      int    `json:"notifyEvery"`
	IgnoreHttpStatus bool   `json:"ignoreHttpStatus"`
	RetryFailed      bool   `json:"retryFailed"`
	FailureThreshold int    `json:"failureThreshold"`
	Pattern          string `json:"pattern"`
	Points           int    `json:"points"`
}

type CronLog struct {
	JobId  int
	No     int
	result []struct {
		Timezone          string
		Output            string
		Downloaded        int
		Url               string
		HttpStatus        int
		Error             string
		Message           string
		Ip                string
		SourceIp          string
		Time              time.Time
		StartTime         float32
		ExecutionTime     float32
		Status            int
		ConnectTime       float32
		NameLookupTime    float32
		PreTransferTime   float32
		StartTransferTime float32
		OutputShortened   string
		VerboseData		  string
	}
	StartTime     float32
	ExecutionTime float32
	Status        int
}

type CronConfig struct {

}

type Crons interface {
	List(keyword string) (*[]Cron, error)
	Get(id int) (*Cron, error)
	Add(cron *Cron) (*Cron, error)
	Edit(id int, cron *Cron) (*Cron, error)
	Enable(id int) (*Cron, error)
	Disable(id int) (*Cron, error)
	Delete(id int) (error)
	Run(id int) (*Cron, error)
	Logs(id int) (*Cron, error)
	Failures(id int) (*Cron, error)
}

type crons struct {
	client Client
}

func newCrons(c Client) Crons {
	return &crons{
		client: c,
	}
}

func (c *crons) List(keyword string) (*[]Cron, error) {
	return nil, nil
}

func (c *crons) Get(id int) (*Cron, error) {
	_, err := c.client.newRequest(http.MethodGet, "cron", "get", url.Values{"id": {fmt.Sprint(id)}}, nil)

	return nil, err
}

func (c *crons) Add(cron *Cron) (*Cron, error) {
	body, err := json.Marshal(cron)
	_, err = c.client.newRequest(http.MethodGet, "cron", "get", nil, strings.NewReader(string(body)))

	return nil, err
}

func (c *crons) Edit(id int, cron *Cron) (*Cron, error) {
	return nil, nil
}

func (c *crons) Enable(id int) (*Cron, error) {
	return nil, nil
}

func (c *crons) Disable(id int) (*Cron, error) {
	return nil, nil
}

func (c *crons) Delete(id int) (error) {
	_, err := c.client.newRequest(http.MethodGet, "cron", "delete", url.Values{"id": {fmt.Sprint(id)}}, nil)

	return err
}

func (c *crons) Run(id int) (*Cron, error) {
	return nil, nil
}

func (c *crons) Logs(id int) (*Cron, error) {
	return nil, nil
}

func (c *crons) Failures(id int) (*Cron, error) {
	return nil, nil
}

