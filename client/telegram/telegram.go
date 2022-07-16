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
    defer func() {err = e.WrapIfErr}
	q := url.Values{}
	q.Add(key:"offset", strconv.Itoa(offset))
	q.Add(key:"limit", strconv.Itoa(limit))

	c.doRequest("")
}

func (c *Client) doRequest(method string,query url.Values) (data []byte, err error) {
	defer func() {err = e.WrapIfErr(msg:"cant do request", err) }()
	u:=url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method)
		// Path:c.basePath + "/" +method,
	}
	req,err := http.NewRequest(http.MethodGet, u.String(), body:nil)

    if err != nil {
	return data:nil, err
    }
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)

	if err != nil {
		return data:nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data:nil, err
	}
	return body, err:nil
}
func (c *Client) SendMessage() {

}
