package pg

import (
	. "bonjour_nails/internal/masters"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Repository) GetMasters() ([]Master, error) {
	query := fmt.Sprintf(`
		select id, first_name, last_name, phone, position
		  from masters`)
	res, err := r.Query(query)
	if err != nil {
		return nil, fmt.Errorf("database error: %s", err)
	}
	var masters []Master
	for res.Next() {
		m := Master{}
		err := res.Scan(&m.ID, &m.FirstName, &m.LastName, &m.Phone, &m.Position)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("masters not found")
		}
		if err != nil {
			return nil, fmt.Errorf("user scanner error: %w", err)
		}
		masters = append(masters, m)
	}
	return masters, nil
}
