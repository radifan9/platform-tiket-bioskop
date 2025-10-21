package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radifan9/platform-tiket-bioskop/models"
)

type ScheduleRepository struct {
	db *pgxpool.Pool
}

func NewScheduleRepository(db *pgxpool.Pool) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) CreateSchedule(ctx context.Context, s models.Schedule) (models.Schedule, error) {
	query := `
		INSERT INTO schedules (city_id, cinema_id, movie_id, start_at, ticket_price, show_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`
	err := r.db.QueryRow(ctx, query,
		s.CityID, s.CinemaID, s.MovieID, s.StartAt, s.TicketPrice, s.ShowDate,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)

	return s, err
}
