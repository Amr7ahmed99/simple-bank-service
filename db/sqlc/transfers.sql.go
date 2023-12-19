// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: transfers.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  owner, from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3, $4
) RETURNING id, owner, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	Owner         string `json:"owner"`
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.queryRow(ctx, q.createTransferStmt, createTransfer,
		arg.Owner,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteTransferStmt, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, owner, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.queryRow(ctx, q.getTransferStmt, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, owner, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
LIMIT $1
`

func (q *Queries) ListTransfers(ctx context.Context, limit int32) ([]Transfer, error) {
	rows, err := q.query(ctx, q.listTransfersStmt, listTransfers, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfers 
SET owner= $2, from_account_id= $3, 
to_account_id= $4, amount= $5
WHERE id= $1
RETURNING id, owner, from_account_id, to_account_id, amount, created_at
`

type UpdateTransferParams struct {
	ID            int64  `json:"id"`
	Owner         string `json:"owner"`
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.queryRow(ctx, q.updateTransferStmt, updateTransfer,
		arg.ID,
		arg.Owner,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
