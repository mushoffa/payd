package repository

import (
	"context"
	"fmt"

	"payd/infrastructure/database"
	"payd/infrastructure/trace/embedded"
	"payd/shift/domain/entity"
	"payd/shift/domain/repository"
	"payd/shift/domain/valueobject"

	"github.com/jackc/pgx/v5"
)

type shifts struct {
	embedded.Monitor
	db database.DatabaseService
}

func NewShiftsRepository(db database.DatabaseService) domain.ShiftRepository {
	r := &shifts{db: db}
	r.Init()
	return r
}

func (r *shifts) Create(ctx context.Context, shift *entity.Shift) (int, error) {
	childCtx, span := r.Trace(ctx, "Repository.Create")
	defer span.End()

	const sql = `insert into shifts(date, start_time, end_time, role, location) values($1,$2,$3,$4,$5) returning id;`

	date := shift.Date.Time()
	start_time := shift.StartTime.Time()
	end_time := shift.EndTime.Time()
	role := shift.Role
	location := shift.Location

	id, err := r.db.Insert(childCtx, sql, date, start_time, end_time, role, location)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *shifts) FindAll(ctx context.Context) ([]entity.Shift, error) {
	childCtx, span := r.Trace(ctx, "Repository.FindAll")
	defer span.End()

	var shifts []entity.Shift

	const sql = `select * from shifts order by date`

	rows, err := r.db.QueryMany(childCtx, sql)
	if err != nil {
		return shifts, err
	}

	tables, err := pgx.CollectRows(rows.(pgx.Rows), pgx.RowToStructByName[ShiftTable])
	if err != nil {
		return shifts, err
	}

	for _, table := range tables {

		shift := entity.Shift{
			ID:        table.ID,
			Date:      valueobject.ShiftDate(table.Date),
			StartTime: valueobject.ShiftTime(table.StartTime),
			EndTime:   valueobject.ShiftTime(table.EndTime),
			Role:      table.Role,
			Location:  table.Location,
		}

		shifts = append(shifts, shift)
	}

	return shifts, nil
}

func (r *shifts) FindByID(ctx context.Context, id int) (entity.Shift, error) {
	childCtx, span := r.Trace(ctx, "Repository.FindByID", r.Attribute("shift.id", fmt.Sprintf("%d", id)))
	defer span.End()

	var shift entity.Shift
	const sql = `select * from shifts where id=$1`

	row, err := r.db.QueryMany(childCtx, sql, id)
	if err != nil {
		return shift, err
	}

	table, err := pgx.CollectOneRow(row.(pgx.Rows), pgx.RowToStructByName[ShiftTable])
	if err != nil {
		return shift, err
	}

	shift.ID = table.ID
	shift.Date = valueobject.ShiftDate(table.Date)
	shift.StartTime = valueobject.ShiftTime(table.StartTime)
	shift.EndTime = valueobject.ShiftTime(table.EndTime)
	shift.Role = table.Role
	shift.Location = table.Location
	if table.EmployeeName != nil {
		shift.EmployeeName = *table.EmployeeName
	}

	if table.EmployeeName != nil {
		shift.EmployeeID = *table.EmployeeID
	}

	return shift, nil
}

func (r *shifts) Update(ctx context.Context, shift *entity.Shift) error {
	return nil
}

func (r *shifts) Delete(ctx context.Context, id int) error {
	const sql = `delete from shifts where id=$1`

	_, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}
