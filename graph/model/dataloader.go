package model

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const loadersKey = "dataloaders"

type Loaders struct {
	Reference ReferenceLoader
}

func DataLoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(loadersKey, &Loaders{
			Reference: ReferenceLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(domains []string) ([]*Reference, []error) {
					idIdentifier := "100"
					if domains[0] == "test2" {
						idIdentifier = "200"
					}
					references := []*Reference{}
					references = append(references, &Reference{
						ID: idIdentifier,
					})
					return references, nil
				},
			},
		})

		c.Next()
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
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
