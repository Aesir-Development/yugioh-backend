package dbConnection

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"reflect"

	"github.com/Aesir-Development/yugioh-backend/pkg/card"
)

// CARDS TABLE
/*
	"id INT NOT NULL AUTO_INCREMENT," +
	"name VARCHAR(255) NOT NULL," +
	"type VARCHAR(255) NOT NULL," +
	"frame_type VARCHAR(255) NOT NULL," +
	"description VARCHAR(255) NOT NULL," +
	"attack INT NOT NULL," +
	"defense INT NOT NULL," +
	"level INT NOT NULL," +
	"race VARCHAR(255) NOT NULL," +
	"attribute VARCHAR(255) NOT NULL," +
	"card_sets JSON," +
	"card_images JSON," +
	"card_prices JSON," +
	"PRIMARY KEY (id)" +
*/

// GetCards - Get all cards from API
func GetCards() []card.Card {
	resp, err := http.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading response body")
	}

	cards := card.ParseCards(body)

	return cards
}

type jsonWrapper struct {
	cardImages []card.CardImage
	cardSets []card.CardSet
	cardPrices []card.CardPrice
}

// SaveCards - Save cards to DB
func SaveCards(cards []card.Card) {
	for _, card := range cards {
		cardImagesJSON, cardSetsJSON, cardPricesJSON := CardJson(jsonWrapper {
			cardImages: card.CardImages, 
			cardSets: card.CardSets, 
			cardPrices: card.CardPrices,
		})
	
		_, err := DB.Exec("INSERT INTO cards (name, type, frame_type, description, attack, defense, level, race, attribute, card_sets, card_images, card_prices) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		card.Name, card.Type, card.FrameType, card.Description, card.Attack, card.Defense, card.Level, card.Race, card.Attribute, cardSetsJSON, cardImagesJSON, cardPricesJSON)
	
		if err != nil {
			panic(err)
		}
	}
}

func CardJson(wrapper jsonWrapper) (string, string, string) {
	cardImagesJSON, err := json.Marshal(wrapper.cardImages)
	if err != nil {
		panic(err)
	}

	cardSetsJSON, err := json.Marshal(wrapper.cardSets)
	if err != nil {
		panic(err)
	}

	cardPricesJSON, err := json.Marshal(wrapper.cardPrices)
	if err != nil {
		panic(err)
	}

	return string(cardImagesJSON), string(cardSetsJSON), string(cardPricesJSON)
}

func FetchCardByName(name string) card.Card {
	result, err := DB.Query("SELECT * FROM cards WHERE name = ?", name)
	if err != nil {
		panic(err)
	}

	card := ScanRows(*result)

	return card
}

// TODO - Better error handling for this
// NOTE - Untested, use with caution
func FetchCardsByName(name string, limit int) []card.Card {
	result, err := DB.Query("SELECT * FROM cards WHERE name = ? LIMIT ?", name, limit)
	if err != nil {
		panic(err)
	}

	cards := []card.Card {}

	for result.Next() {
		card := ScanRows(*result)
		cards = append(cards, card)
	}

	return cards
}

func FetchCardByID(ID int) card.Card {
	result, err := DB.Query("SELECT * FROM cards WHERE id = ?", ID)
	if err != nil {
		panic(err)
	}

	card := ScanRows(*result)

	return card
}

// map for DB names to Struct names
var m map[string]string = map[string]string{
	"id": "ID",
	"name": "Name",
	"type": "Type",
	"frame_type": "FrameType",
	"description": "Description",
	"attack": "Attack",
	"defense": "Defense",
	"level": "Level",
	"race": "Race",
	"attribute": "Attribute",
	"card_sets": "CardSets",
	"card_images": "CardImages",
	"card_prices": "CardPrices",
}
	
type CardSQLWrapper struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	FrameType string `json:"frameType"`
	Description string `json:"desc"`
	Attack int `json:"atk"`
	Defense int `json:"def"`
	Level int `json:"level"`
	Race string `json:"race"`
	Attribute string `json:"attribute"`
	CardSets []uint8 `json:"card_sets"`
	CardImages []uint8 `json:"card_images"`
	CardPrices []uint8 `json:"card_prices"`
}

// Extracting the card data from the SQL query, because it's not supported by default
func ScanRows(rows sql.Rows) card.Card {
	columnNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	returnCard := card.Card{}
	
	for rows.Next() {
		cardWrapped := CardSQLWrapper {}
		pointers := make([]interface{}, len(columnNames))
		structVal := reflect.ValueOf(&cardWrapped).Elem()
		for i, colName := range columnNames {
			colName = m[colName]

			if(colName == "") {
				continue
			}

			fieldVal := structVal.FieldByName(colName)
			if !fieldVal.IsValid() {
				println("Column name: " + colName)
				log.Fatalf("No such field: %s in obj", colName)
			}
			fieldAddr := fieldVal.Addr()
			pointers[i] = fieldAddr.Interface()
		}
		err := rows.Scan(pointers...)
		if err != nil {
			log.Fatal(err)
		}

		cardSets, cardImages, cardPrices := CardStructsFromUint8(cardWrapped.CardSets, cardWrapped.CardImages, cardWrapped.CardPrices)

		returnCard = card.Card {
			ID: cardWrapped.ID,
			Name: cardWrapped.Name,
			Type: cardWrapped.Type,
			FrameType: cardWrapped.FrameType,
			Description: cardWrapped.Description,
			Attack: cardWrapped.Attack,
			Defense: cardWrapped.Defense,
			Level: cardWrapped.Level,
			Race: cardWrapped.Race,
			Attribute: cardWrapped.Attribute,
			CardSets: cardSets,
			CardImages: cardImages,
			CardPrices: cardPrices,
		}

		println(returnCard.Name)

	}

	return returnCard

}

func CardStructsFromUint8(CardSet []uint8, CardImage []uint8, CardPrice []uint8) ([]card.CardSet, []card.CardImage, []card.CardPrice) {
	cardSets := []card.CardSet {}
	cardImages := []card.CardImage {}
	cardPrices := []card.CardPrice {}
	
	err := json.Unmarshal(CardSet, &cardSets)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(CardImage, &cardImages)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(CardPrice, &cardPrices)
	if err != nil {
		panic(err)
	}


	return cardSets, cardImages, cardPrices
}