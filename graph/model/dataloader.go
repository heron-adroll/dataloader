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
				wait:     1000 * time.Millisecond,
				fetch: func(domains []string) ([]*Reference, []error) {

					references := []*Reference{}
					for _, domain := range domains {
						references = append(references, &Reference{
							ID: domain,
						})
					}

					fmt.Println(domains)
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
