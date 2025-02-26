package nautils

import (
	"context"

	"github.com/jomei/notionapi"
	"github.com/samber/lo"
)

func (c *Client) SearchPage(ctx context.Context, query string) ([]*notionapi.Page, error) {
	rawC := c.RawClient()
	q := notionapi.SearchRequest{
		Query: query,
		Filter: notionapi.SearchFilter{
			Property: "object",
			Value:    "page",
		},
	}
	res, err := rawC.Search.Do(ctx, &q)
	if err != nil {
		return nil, err
	}

	pages := lo.Map(res.Results, func(item notionapi.Object, n int) *notionapi.Page {
		page, err := ObjctToPage(item)
		if err != nil {
			return nil
		}
		return page
	})
	pages = lo.Filter(pages, func(page *notionapi.Page, n int) bool {
		return page != nil
	})
	return pages, nil
}
