package nautils

import (
	"errors"

	"github.com/jomei/notionapi"
)

func ObjctToPage(o notionapi.Object) (*notionapi.Page, error) {
	if o == nil {
		return nil, nil
	}
	if o.GetObject().String() != "page" {
		return nil, errors.New("object is not page")
	}
	return o.(*notionapi.Page), nil
}
