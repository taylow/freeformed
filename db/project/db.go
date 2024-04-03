// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.changeProjectTeamStmt, err = db.PrepareContext(ctx, changeProjectTeam); err != nil {
		return nil, fmt.Errorf("error preparing query ChangeProjectTeam: %w", err)
	}
	if q.createProjectStmt, err = db.PrepareContext(ctx, createProject); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProject: %w", err)
	}
	if q.deleteProjectStmt, err = db.PrepareContext(ctx, deleteProject); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProject: %w", err)
	}
	if q.disableProjectStmt, err = db.PrepareContext(ctx, disableProject); err != nil {
		return nil, fmt.Errorf("error preparing query DisableProject: %w", err)
	}
	if q.enableProjectStmt, err = db.PrepareContext(ctx, enableProject); err != nil {
		return nil, fmt.Errorf("error preparing query EnableProject: %w", err)
	}
	if q.listProjectsStmt, err = db.PrepareContext(ctx, listProjects); err != nil {
		return nil, fmt.Errorf("error preparing query ListProjects: %w", err)
	}
	if q.listProjectsByTeamStmt, err = db.PrepareContext(ctx, listProjectsByTeam); err != nil {
		return nil, fmt.Errorf("error preparing query ListProjectsByTeam: %w", err)
	}
	if q.readProjectStmt, err = db.PrepareContext(ctx, readProject); err != nil {
		return nil, fmt.Errorf("error preparing query ReadProject: %w", err)
	}
	if q.softDeleteProjectStmt, err = db.PrepareContext(ctx, softDeleteProject); err != nil {
		return nil, fmt.Errorf("error preparing query SoftDeleteProject: %w", err)
	}
	if q.unsoftDeleteProjectStmt, err = db.PrepareContext(ctx, unsoftDeleteProject); err != nil {
		return nil, fmt.Errorf("error preparing query UnsoftDeleteProject: %w", err)
	}
	if q.updateProjectStmt, err = db.PrepareContext(ctx, updateProject); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProject: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.changeProjectTeamStmt != nil {
		if cerr := q.changeProjectTeamStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing changeProjectTeamStmt: %w", cerr)
		}
	}
	if q.createProjectStmt != nil {
		if cerr := q.createProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProjectStmt: %w", cerr)
		}
	}
	if q.deleteProjectStmt != nil {
		if cerr := q.deleteProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProjectStmt: %w", cerr)
		}
	}
	if q.disableProjectStmt != nil {
		if cerr := q.disableProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing disableProjectStmt: %w", cerr)
		}
	}
	if q.enableProjectStmt != nil {
		if cerr := q.enableProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing enableProjectStmt: %w", cerr)
		}
	}
	if q.listProjectsStmt != nil {
		if cerr := q.listProjectsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProjectsStmt: %w", cerr)
		}
	}
	if q.listProjectsByTeamStmt != nil {
		if cerr := q.listProjectsByTeamStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProjectsByTeamStmt: %w", cerr)
		}
	}
	if q.readProjectStmt != nil {
		if cerr := q.readProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing readProjectStmt: %w", cerr)
		}
	}
	if q.softDeleteProjectStmt != nil {
		if cerr := q.softDeleteProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing softDeleteProjectStmt: %w", cerr)
		}
	}
	if q.unsoftDeleteProjectStmt != nil {
		if cerr := q.unsoftDeleteProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unsoftDeleteProjectStmt: %w", cerr)
		}
	}
	if q.updateProjectStmt != nil {
		if cerr := q.updateProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProjectStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                      DBTX
	tx                      *sql.Tx
	changeProjectTeamStmt   *sql.Stmt
	createProjectStmt       *sql.Stmt
	deleteProjectStmt       *sql.Stmt
	disableProjectStmt      *sql.Stmt
	enableProjectStmt       *sql.Stmt
	listProjectsStmt        *sql.Stmt
	listProjectsByTeamStmt  *sql.Stmt
	readProjectStmt         *sql.Stmt
	softDeleteProjectStmt   *sql.Stmt
	unsoftDeleteProjectStmt *sql.Stmt
	updateProjectStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                      tx,
		tx:                      tx,
		changeProjectTeamStmt:   q.changeProjectTeamStmt,
		createProjectStmt:       q.createProjectStmt,
		deleteProjectStmt:       q.deleteProjectStmt,
		disableProjectStmt:      q.disableProjectStmt,
		enableProjectStmt:       q.enableProjectStmt,
		listProjectsStmt:        q.listProjectsStmt,
		listProjectsByTeamStmt:  q.listProjectsByTeamStmt,
		readProjectStmt:         q.readProjectStmt,
		softDeleteProjectStmt:   q.softDeleteProjectStmt,
		unsoftDeleteProjectStmt: q.unsoftDeleteProjectStmt,
		updateProjectStmt:       q.updateProjectStmt,
	}
}