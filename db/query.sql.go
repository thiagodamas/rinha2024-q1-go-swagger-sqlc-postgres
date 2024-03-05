// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const credito = `-- name: Credito :one
with dados as (
	select id_conta, limite, nome_cliente from dados_bancarios 
	where dados_bancarios.id_conta=$3
	for update
),
informacoes as (
    select 
	case
	  when (saldos.saldo+dados.limite) >= $1 then true
	  else false
	end as autorizado,	  
	dados.id_conta,saldos.saldo,dados.limite,(saldos.saldo+dados.limite) as disponivel from dados
	join saldos on dados.id_conta=saldos.id_conta
	for update
),
credito as (
	update saldos ss set saldo=ss.saldo+$1
	from informacoes i
	where ss.id_conta=i.id_conta
	returning ss.id_conta, ss.saldo,i.limite
),
inserehistorico as (
    insert into transacoes (id_conta,tipo_operacao,valor,descricao)
	select id_conta, 'c', $1, $2::text from informacoes where informacoes.autorizado
	returning id, id_conta, tipo_operacao, valor, descricao, created_at
)
select inf.autorizado, inf.id_conta, inf.saldo as saldo_anterior, inf.limite, coalesce(c.saldo,inf.saldo) as saldo, $1::bigint as valor, $2::text as descricao
from informacoes inf
left join credito c on inf.id_conta=c.id_conta
left join inserehistorico i on c.id_conta=i.id_conta
`

type CreditoParams struct {
	Valor     int64
	Descricao string
	IDConta   int32
}

type CreditoRow struct {
	Autorizado    bool
	IDConta       int32
	SaldoAnterior int64
	Limite        int64
	Saldo         int64
	Valor         int64
	Descricao     string
}

func (q *Queries) Credito(ctx context.Context, arg CreditoParams) (CreditoRow, error) {
	row := q.db.QueryRow(ctx, credito, arg.Valor, arg.Descricao, arg.IDConta)
	var i CreditoRow
	err := row.Scan(
		&i.Autorizado,
		&i.IDConta,
		&i.SaldoAnterior,
		&i.Limite,
		&i.Saldo,
		&i.Valor,
		&i.Descricao,
	)
	return i, err
}

const debito = `-- name: Debito :one
with dados as (
	select id_conta, limite, nome_cliente from dados_bancarios 
	where dados_bancarios.id_conta=$3
	for update
),
informacoes as (
    select 
	case
	  when (saldos.saldo+dados.limite) >= $1 then true
	  else false
	end as autorizado,	  
	dados.id_conta,saldos.saldo,dados.limite,(saldos.saldo+dados.limite) as disponivel from dados
	join saldos on dados.id_conta=saldos.id_conta
	for update
),
debito as (
	update saldos ss set saldo=ss.saldo-$1
	from informacoes i
	where ss.id_conta=i.id_conta and i.autorizado
	returning ss.id_conta, ss.saldo,i.limite
),
inserehistorico as (
    insert into transacoes (id_conta,tipo_operacao,valor,descricao)
	select id_conta, 'd', $1, $2::text from informacoes where informacoes.autorizado
	returning id, id_conta, tipo_operacao, valor, descricao, created_at
)
select inf.autorizado, inf.id_conta, inf.saldo as saldo_anterior, inf.limite, coalesce(d.saldo,inf.saldo) as saldo, $1::bigint as valor, $2::text as descricao
from informacoes inf
left join debito d on inf.id_conta=d.id_conta
left join inserehistorico i on d.id_conta=i.id_conta
`

type DebitoParams struct {
	Valor     int64
	Descricao string
	IDConta   int32
}

type DebitoRow struct {
	Autorizado    bool
	IDConta       int32
	SaldoAnterior int64
	Limite        int64
	Saldo         int64
	Valor         int64
	Descricao     string
}

func (q *Queries) Debito(ctx context.Context, arg DebitoParams) (DebitoRow, error) {
	row := q.db.QueryRow(ctx, debito, arg.Valor, arg.Descricao, arg.IDConta)
	var i DebitoRow
	err := row.Scan(
		&i.Autorizado,
		&i.IDConta,
		&i.SaldoAnterior,
		&i.Limite,
		&i.Saldo,
		&i.Valor,
		&i.Descricao,
	)
	return i, err
}

const extrato = `-- name: Extrato :many
with dados as (
	select id_conta, limite, nome_cliente from dados_bancarios 
	where dados_bancarios.id_conta=$1
	for update
),
informacoes as (
    select dados.id_conta,saldos.saldo,dados.limite,(saldos.saldo+dados.limite) as disponivel from dados
	join saldos on dados.id_conta=saldos.id_conta
),
extrato as (
	select id, id_conta, tipo_operacao, valor, descricao, created_at from transacoes
	order by id desc
)
select i.saldo, now()::timestamp without time zone as data_extrato, i.limite, COALESCE(e.tipo_operacao, 'e')::text as tipo_operacao, e.valor, COALESCE(e.descricao,'')::text as descricao, e.created_at as realizada_em
from informacoes i left join extrato e on e.id_conta=i.id_conta
order by e.id desc limit 10
`

type ExtratoRow struct {
	Saldo        int64
	DataExtrato  pgtype.Timestamp
	Limite       int64
	TipoOperacao string
	Valor        pgtype.Int8
	Descricao    string
	RealizadaEm  pgtype.Timestamp
}

func (q *Queries) Extrato(ctx context.Context, idConta int32) ([]ExtratoRow, error) {
	rows, err := q.db.Query(ctx, extrato, idConta)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ExtratoRow
	for rows.Next() {
		var i ExtratoRow
		if err := rows.Scan(
			&i.Saldo,
			&i.DataExtrato,
			&i.Limite,
			&i.TipoOperacao,
			&i.Valor,
			&i.Descricao,
			&i.RealizadaEm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}