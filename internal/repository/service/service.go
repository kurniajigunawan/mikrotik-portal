package service

import "context"

func (s *Repository) GetServiceByClientID(ctx context.Context, clientID string) (Service, error) {
	query := "SELECT id, name, client_id, client_secret, status, created_at, updated_at FROM services WHERE client_id = $1"
	row := s.db.QueryRowContext(ctx, query, clientID)
	var service Service
	if err := row.Scan(&service.ID, &service.Name, &service.ClientID, &service.ClientSecret, &service.Status, &service.CreatedAt, &service.UpdatedAt); err != nil {
		return Service{}, err
	}
	return service, nil
}

func (s *Repository) GetActiveServices(ctx context.Context) ([]Service, error) {
	query := "SELECT id, name, client_id, client_secret, status, created_at, updated_at FROM services WHERE status = $1"
	rows, err := s.db.QueryContext(ctx, query, StatusActive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var services []Service
	for rows.Next() {
		var service Service
		if err := rows.Scan(&service.ID, &service.Name, &service.ClientID, &service.ClientSecret, &service.Status, &service.CreatedAt, &service.UpdatedAt); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}
