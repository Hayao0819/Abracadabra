package nautils

import (
	"context"

	"github.com/jomei/notionapi"
)

type FullPage struct {
	Page   *notionapi.Page
	Blocks []notionapi.Block
	Title  *notionapi.TitleProperty
}

func (c *Client) FullPage(p *notionapi.Page) (*FullPage, error) {
	blocks, err := c.PageBlocks(context.Background(), p)
	if err != nil {
		return nil, err
	}
	title := PageTitle(p)
	return &FullPage{
		Page:   p,
		Blocks: blocks,
		Title:  title,
	}, nil
}

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

func (c *Client) FullPageFromID(ctx context.Context, pageId notionapi.PageID) (*FullPage, error) {
	page, err := c.RawClient().Page.Get(ctx, pageId)
	if err != nil {
		return nil, err
	}

	fullPage, err := c.FullPage(page)
	if err != nil {
		return nil, err
	}
	return fullPage, nil
}
