package advert

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
)

func NewAdvertSelector(db *pgx.Conn) Selector {
	return Selector{
		db,
	}
}

type Selector struct {
	db *pgx.Conn
}

func (s *Selector) MatchAdvert(dates, devices []string) []string {
	sqlQuery := `SELECT advert_id FROM t_adverts_targeting t1`

	hasDates := len(dates) > 0
	hasDevices := len(devices) > 0

	if hasDates || hasDevices {
		sqlQuery += " WHERE "
	}

	if hasDates {
		sqlQuery += "("
		var sqlDateChunks []string
		for _, date := range dates {
			sqlDateChunks = append(sqlDateChunks, fmt.Sprintf("t1.dates @> '[\"%s\"]'", date))
		}
		sqlQuery += strings.Join(sqlDateChunks, " OR ")
		sqlQuery += ")"
	}

	if hasDevices {
		sqlQuery += " AND ("
		var sqlDateChunks []string
		for _, date := range devices {
			sqlDateChunks = append(sqlDateChunks, fmt.Sprintf("t1.devices @> '[\"%s\"]'", date))
		}
		sqlQuery += strings.Join(sqlDateChunks, " OR ")
		sqlQuery += ")"
	}

	sqlQuery += " ORDER BY cost"
	ctx := context.Background()

	var advertIds []string

	rows, _ := s.db.Query(ctx, sqlQuery)
	for rows.Next() {
		var advertId string
		_ = rows.Scan(&advertId)

		advertIds = append(advertIds, advertId)
	}

	return advertIds
}
