-- name: GetItineraryItems :many
SELECT * FROM Itinerary_Items
WHERE TripID = ?;