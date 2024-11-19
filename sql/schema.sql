CREATE TABLE `Users` (
  `UserID` int NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(255) DEFAULT NULL,
  `LastName` varchar(255) DEFAULT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `Username` varchar(255) DEFAULT NULL,
  `Preferences` json DEFAULT NULL,
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Trips` (
  `tripId` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `userId` int DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `place_id` varchar(255) DEFAULT NULL,
  `photo_uri` varchar(500) DEFAULT NULL,
  `latitude` int DEFAULT NULL,
  `longitude` int DEFAULT NULL,
  PRIMARY KEY (`tripId`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Itinerary_Items` (
  `ItemID` int NOT NULL AUTO_INCREMENT,
  `TripID` int DEFAULT NULL,
  `Title` varchar(255) DEFAULT NULL,
  `Location` varchar(255) DEFAULT NULL,
  `Date` datetime DEFAULT NULL,
  PRIMARY KEY (`ItemID`),
  KEY `TripID` (`TripID`),
  CONSTRAINT `Itinerary_Items_ibfk_1` FOREIGN KEY (`TripID`) REFERENCES `Trips` (`TripID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

