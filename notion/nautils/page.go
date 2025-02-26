package nautils

import (
	"context"

	"github.com/jomei/notionapi"
)

func PageTitle(p *notionapi.Page) *notionapi.TitleProperty {
	rawTitle, ok := p.Properties["title"]
	if !ok {
		return nil
	}
	title, ok := rawTitle.(*notionapi.TitleProperty)
	if !ok {
		return nil
	}
	return title
}

func (c *Client) PageBlock(ctx context.Context, p *notionapi.Page) (notionapi.Block, error) {
	return c.RawClient().Block.Get(ctx, notionapi.BlockID(p.ID.String()))
}

func (c *Client) PageBlocks(ctx context.Context, p *notionapi.Page) ([]notionapi.Block, error) {
	firstRes, err := c.RawClient().Block.GetChildren(ctx, notionapi.BlockID(p.ID.String()), nil)
	if err != nil {
		return nil, err
	}

	reses := []*notionapi.GetChildrenResponse{firstRes}
	for {
		lastRes := reses[len(reses)-1]
		if !lastRes.HasMore {
			break
		}
		res, err := c.RawClient().Block.GetChildren(ctx, notionapi.BlockID(p.ID.String()), &notionapi.Pagination{
			StartCursor: notionapi.Cursor(lastRes.NextCursor),
		})
		if err != nil {
			break
		}
		reses = append(reses, res)
	}
	// utils.PrintAsJSON(reses)

	blocks := []notionapi.Block{}
	for _, res := range reses {
		blocks = append(blocks, res.Results...)
	}
	return blocks, nil
}
