package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/heron.rossi/dataloader/graph/generated"
	"github.com/heron.rossi/dataloader/graph/model"
)

func (r *accountResolver) Reference(ctx context.Context, obj *model.Account) (*model.Reference, error) {
	ginContext, err := model.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return model.For(ginContext).Reference.Load(obj.Domain)
}

func (r *queryResolver) Accounts(ctx context.Context, domain *string) ([]*model.Account, error) {
	accounts := []*model.Account{}
	accounts = append(accounts, &model.Account{
		Domain: *domain,
	})

	return accounts, nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type accountResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
