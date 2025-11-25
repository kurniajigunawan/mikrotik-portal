package event

import "context"

func (r *Repository) CreateEvent(ctx context.Context, event CreateEventRequest) error {
	query := `INSERT INTO events (service_id, event_type, value, status) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, event.ServiceID, event.EventType, event.Value, event.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SetStatus(ctx context.Context, eventID, serviceID int64, status Status) error {
	query := `UPDATE events SET status = $1 WHERE id = $2 AND service_id = $3`
	_, err := r.db.ExecContext(ctx, query, status, eventID, serviceID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetActiveEventsByServiceID(ctx context.Context, serviceID int64) ([]Event, error) {
	query := `SELECT id, service_id, event_type, value, status, created_at, updated_at FROM events WHERE service_id = $1 AND status = $2`
	rows, err := r.db.QueryContext(ctx, query, serviceID, StatusWaitingToConsume)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.ServiceID, &event.EventType, &event.Value, &event.Status, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
