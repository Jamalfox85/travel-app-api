-- name: GetItineraryItems :many
SELECT * FROM Itinerary_Items
WHERE TripID = ?;

-- name: CreateItineraryItem :exec
INSERT INTO Itinerary_Items (tripId, title, date, url, phone, address, poi_id)
VALUES (?, ?, ?, ?, ?, ?, ?);