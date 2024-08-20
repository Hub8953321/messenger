package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"messager/src/pkg/logger"
)

type PostgresLogger struct {
	logger.Logger
}

var _ pgx.QueryTracer = (*PostgresLogger)(nil)

func NewPostgresLogger(logger logger.Logger) *PostgresLogger {
	return &PostgresLogger{logger}
}

func (l *PostgresLogger) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	if len(data.Args) > 0 {
		l.Info(fmt.Sprintf("%s, with args: %v\n", data.SQL, data.Args))
	} else {
		l.Info(data.SQL)
	}
	return ctx
}

func (l *PostgresLogger) TraceQueryEnd(_ context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	if data.Err != nil {
		l.Error("tag: " + data.CommandTag.String() + " err: " + data.Err.Error())
	}
}
