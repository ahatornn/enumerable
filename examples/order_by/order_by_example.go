// go run examples/order_by/order_by_example.go
package main

import (
	"fmt"

	"github.com/ahatornn/enumerable"
	"github.com/ahatornn/enumerable/comparer"
)

type Movie struct {
	Title    string
	Director string
	Actors   []string
	Year     int
	Rating   float64
	Tags     []string
}

func main() {
	movies := []Movie{
		{Title: "Film 01", Director: "Dir1", Actors: []string{"Actor1", "Actor2"}, Year: 2020, Rating: 8.5, Tags: []string{"drama", "mystery"}},
		{Title: "Film 02", Director: "Dir2", Actors: []string{"Actor3", "Actor4"}, Year: 2021, Rating: 8.2, Tags: []string{"action", "sci-fi"}},
		{Title: "Film 03", Director: "Dir3", Actors: []string{"Actor5", "Actor6"}, Year: 2022, Rating: 9.0, Tags: []string{"comedy", "romance"}},
		{Title: "Film 04", Director: "Dir4", Actors: []string{"Actor7", "Actor8"}, Year: 2022, Rating: 8.7, Tags: []string{"comedy", "family"}},
		{Title: "Film 05", Director: "Dir5", Actors: []string{"Actor9", "Actor1"}, Year: 2019, Rating: 9.1, Tags: []string{"drama", "biography"}},
		{Title: "Film 06", Director: "Dir6", Actors: []string{"Actor2", "Actor3"}, Year: 2021, Rating: 8.8, Tags: []string{"comedy", "crime"}},
		{Title: "Film 07", Director: "Dir6_1", Actors: []string{"Actor3", "Actor4"}, Year: 2021, Rating: 9.2, Tags: []string{"comedy", "crime"}},
		{Title: "Film 08", Director: "Dir7", Actors: []string{"Actor4", "Actor5"}, Year: 2021, Rating: 8.3, Tags: []string{"comedy", "thriller"}},
		{Title: "Film 09", Director: "Dir8", Actors: []string{"Actor6", "Actor7"}, Year: 2020, Rating: 8.1, Tags: []string{"comedy", "musical"}},
		{Title: "Film 10", Director: "Dir9", Actors: []string{"Actor8", "Actor9"}, Year: 2023, Rating: 8.9, Tags: []string{"comedy", "satire"}},
		{Title: "Film 11", Director: "Dir10", Actors: []string{"Actor1", "Actor2"}, Year: 2023, Rating: 8.6, Tags: []string{"comedy", "war"}},
		{Title: "Film 12", Director: "Dir11", Actors: []string{"Actor3", "Actor4"}, Year: 2023, Rating: 8.5, Tags: []string{"thriller", "horror"}},
		{Title: "Film 13", Director: "Dir10_1", Actors: []string{"Actor3", "Actor5"}, Year: 2023, Rating: 9.0, Tags: []string{"comedy", "war"}},
	}

	favoriteActors := enumerable.FromSlice([]string{"Actor3", "Actor8"})
	targetTag := "comedy"

	// Let's assume we are looking for movies that:
	// 1. Contain the tag "comedy"
	// 2. Have a rating > 8.2
	// 3. Have the main actor "Actor3" or "Actor8"
	// Then we sort by year (descending) and rating (descending) and take the first 5.
	topMovies := enumerable.FromSliceAny(movies).
		Where(func(m Movie) bool {
			return m.Rating > 8.2 &&
				enumerable.FromSlice(m.Tags).Contains(targetTag) &&
				enumerable.FromSlice(m.Actors).Intersect(favoriteActors).Any()
		}).
		OrderByDescending(comparer.KeyComparerInt(func(m Movie) int { return m.Year })).
		ThenByDescending(comparer.KeyComparerFloat64(func(m Movie) float64 { return m.Rating })).
		Take(5).
		ToSlice()

	fmt.Println("\n--- Using enumerable ---")
	fmt.Printf("Top 5 movies with tag '%s', rating > 8.2, and actors one of %s)\n", targetTag, favoriteActors.ToSlice())
	fmt.Println("(sorted: year desc., rating desc.):")
	for i, m := range topMovies {
		fmt.Printf("  %d. %s (%d) - Actors: %s, Rating: %.1f, Tags: %v\n",
			i+1, m.Title, m.Year, m.Actors, m.Rating, m.Tags)
	}
}
