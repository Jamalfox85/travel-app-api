-- name: GetItineraryItems :many
SELECT * FROM itinerary_items
WHERE TripID = ?;

-- name: CreateItineraryItem :exec
INSERT INTO itinerary_items (tripId, title, date, url, phone, address, poiId)
VALUES (?, ?, ?, ?, ?, ?, ?);