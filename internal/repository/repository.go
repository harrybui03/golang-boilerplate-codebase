package repository

import (
	"context"
	"fmt"

	"golang-boilerplate/ent"
	"golang-boilerplate/internal/repository/user"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	User() user.Repository
	DoInTx(ctx context.Context, txFunc TxFunc, repo RepositoryRegistry) error
}

type TxFunc func(ctx context.Context, repo RepositoryRegistry) error

// impl is the implementation of the repository registry.
type impl struct {
	client *ent.Client
	user   user.Repository
}

// New creates a new repository registry.
func New(client *ent.Client) RepositoryRegistry {
	return &impl{
		client: client,
		user:   user.New(client),
	}
}

// User returns the user repository.
func (i impl) User() user.Repository {
	return i.user
}

// DoInTx allows to combine multiple repository operations in a single tx
func (i impl) DoInTx(ctx context.Context, txFunc TxFunc, repo RepositoryRegistry) error {
	tx, err := i.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := txFunc(ctx, repo); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
