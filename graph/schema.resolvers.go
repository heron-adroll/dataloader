package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/heron.rossi/dataloader/dataloaders"
	"github.com/heron.rossi/dataloader/graph/generated"
	"github.com/heron.rossi/dataloader/graph/model"
)

func (r *accountResolver) Reference(ctx context.Context, obj *model.Account) (*model.Reference, error) {
	ginContext, err := dataloaders.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	loader, _ := ginContext.Get(dataloaders.LoadersKey)
	thunk := loader.(*dataloaders.Loaders).References.Load(context.TODO(), dataloader.StringKey(obj.Domain))
	data, _ := thunk()
	return data.(*model.Reference), nil
}

func (r *queryResolver) Accounts(ctx context.Context, domain *string) ([]*model.Account, error) {
	accounts := []*model.Account{}
	accounts = append(accounts, &model.Account{
		Domain: *domain,
	})
	// this is just to simulate various objects being retrieved
	accounts = append(accounts, &model.Account{
		Domain: "test2",
	})
	accounts = append(accounts, &model.Account{
		Domain: "test 333",
	})
	return accounts, nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type accountResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
