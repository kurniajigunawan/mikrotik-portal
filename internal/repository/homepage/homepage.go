package homepage

import (
	"context"
)

func (r *Repository) GetActiveMenu(ctx context.Context) ([]Menu, error) {
	query := `SELECT id, title, description, icon, icon_color, link_url FROM homepage_menus WHERE status = 1`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []Menu
	for rows.Next() {
		var menu Menu
		if err := rows.Scan(&menu.ID, &menu.Title, &menu.Description, &menu.Icon, &menu.IconColor, &menu.Link); err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}
	return menus, nil
}
