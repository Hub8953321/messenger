package repository

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	UserPass string
	DBName   string
	SSLMode  string
}

func doWithTries(fn func() error, attempts uint8, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--

			continue
		}
		return nil
	}
	return err
}

//go:embed queries/init.up.sql
var upQuery string

func createTablesAndIndexes(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, upQuery)
	return err
}

//go:embed queries/init.down.sql
var downQuery string

func deleteAllTablesAndIndexes(ctx context.Context, pool *pgxpool.Pool) {
	_, err := pool.Exec(ctx, downQuery)
	if err != nil {
		log.Println(err)
	}
}

func NewPostgresDB(ctx context.Context, maxAttempts uint8, cfg Config, logger *PostgresLogger) (pool *pgxpool.Pool, err error) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.UserName, cfg.UserPass, cfg.Host, cfg.Port, cfg.DBName)
	err = doWithTries(func() error {
		ctx1, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		config, err := pgxpool.ParseConfig(url)
		if err != nil {
			return err
		}

		config.ConnConfig.Tracer = logger
		pool, err = pgxpool.NewWithConfig(ctx1, config)
		if err != nil {
			return err
		}

		if err = pool.Ping(ctx1); err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	logger.Info("creating tables for postgres")
	err = createTablesAndIndexes(ctx, pool)
	if err != nil {
		logger.Fatal(err.Error())
	} else {
		logger.Info("all tables for postgres created")
	}

	return pool, nil
}

func RegenerateIndexes(ctx context.Context, pool *pgxpool.Pool) {
	_, err := pool.Exec(ctx, `
		DROP INDEX IF EXISTS parent_user_idx;
		CREATE INDEX IF NOT EXISTS parent_user_idx ON posts(parent_user_id);
		
		DROP INDEX IF EXISTS posts_likes_user_idx;
		DROP INDEX IF EXISTS posts_likes_parent_post_idx;
		CREATE INDEX IF NOT EXISTS posts_likes_user_idx ON posts_likes(user_id);
		CREATE INDEX IF NOT EXISTS posts_likes_parent_post_idx ON posts_likes(parent_post_id);
		
		DROP INDEX IF EXISTS posts_dislikes_user_idx;
		DROP INDEX IF EXISTS posts_dislikes_parent_post_idx;
		CREATE INDEX IF NOT EXISTS posts_dislikes_user_idx ON posts_dislikes(user_id);
		CREATE INDEX IF NOT EXISTS posts_dislikes_parent_post_idx ON posts_dislikes(parent_post_id);
		
		DROP INDEX IF EXISTS parent_posts_idx;
		CREATE INDEX IF NOT EXISTS parent_posts_idx ON surveys(parent_post_id);
		
		DROP INDEX IF EXISTS surveys_voices_user_idx;
		CREATE INDEX IF NOT EXISTS surveys_voices_user_idx ON surveys_voices(user_id);
		
		DROP INDEX IF EXISTS parent_post_idx;
		CREATE INDEX IF NOT EXISTS parent_post_idx ON comments(parent_post_id);

		DROP INDEX IF EXISTS comments_likes_user_idx;
		DROP INDEX IF EXISTS comments_likes_parent_comment_idx;
		CREATE INDEX IF NOT EXISTS comments_likes_user_idx ON comments_likes(user_id);
		CREATE INDEX IF NOT EXISTS comments_likes_parent_comment_idx ON comments_likes(parent_comment_id);

		DROP INDEX IF EXISTS comments_dislikes_user_idx;
		DROP INDEX IF EXISTS comments_dislikes_parent_comment_idx;
		CREATE INDEX IF NOT EXISTS comments_dislikes_user_idx ON comments_dislikes(user_id);
		CREATE INDEX IF NOT EXISTS comments_dislikes_parent_comment_idx ON comments_dislikes(parent_comment_id);
		
		DROP INDEX IF EXISTS parent_chat_idx;
		CREATE INDEX IF NOT EXISTS parent_chat_idx ON messages(parent_chat_id);
	`)
	if err != nil {
		log.Println(err)
	}
}
