-- name: GetItineraryItems :many
SELECT * FROM itinerary_items
WHERE tripId = ?;

-- name: CreateItineraryItem :exec
INSERT INTO itinerary_items (tripId, title, date, url, phone, address, poiId, rating, price)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);