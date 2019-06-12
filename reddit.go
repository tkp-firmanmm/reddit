//Package reddit implements reddit API for fetching post
package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//Item describe a Reddit post
type Item struct {
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}

//Response describe response from the server
type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

//Get return the most recent items posted on reddit
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}

	return items, nil
}

//String return a formatted string with Title, Comments and URL of the post
func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
	case 1:
		com = " (1 comment)"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}
