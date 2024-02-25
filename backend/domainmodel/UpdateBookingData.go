package domainmodel

import (
	"context"
)

func UpdateBookingData(ctx context.Context) {
	tx := ctx.Value("tx")
	tx.ExecContext(ctx, "UPDATE bookings SET user_id = $1, room_id = $2, start_time = $3, end_time = $4 WHERE id = $5", "test", 20)
}
