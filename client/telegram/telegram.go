package telegram

import "net/http"

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
    q := url.Values{}
	q.Add(key:"offset", strconv.Itoa(offset))
	q.Add(key:"limit", strconv.Itoa(limit))
}

func (c *Client) doRequest(method string,query url.Values) ([]byte, error) {
	u:=url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method)
		// Path:c.basePath + "/" +method,
	}
	req,err := http.NewRequest(http.MethodGet, u.String(), body:nil)

    if err != nil {
	return nil, fmt.Errorf(format:"cant do request: %w", err)
    }
}
func (c *Client) SendMessage() {

}
