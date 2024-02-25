package internal

const GET_ALL_BOOKINGS_DATA = "SELECT * FROM bookings"
const GET_A_BOOKING_DATA = "SELECT * FROM bookings WHERE id = $1"
const UPDATE_BOOKING_DATA = "UPDATE bookings SET user_id = $1, room_id = $2, start_time = $3, end_time = $4 WHERE id = $5"
const UPDATE_BOOKING_STATUS = "UPDATE bookings SET status = $1 WHERE id = $2"
