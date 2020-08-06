package dataloaders

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
	"github.com/heron.rossi/dataloader/graph/model"
)

const LoadersKey = "dataloaders"

type Loaders struct {
	References *dataloader.Loader
}

func DataLoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(LoadersKey, &Loaders{
			References: dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
				var results []*dataloader.Result
				// we should use keys to make the necessary API calls
				for _, val := range keys {
					results = append(results, &dataloader.Result{Data: &model.Reference{ID: val.String()}})
				}
				return results
			}),
		})
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
