package product

import (
	"fmt"
	"github.com/Nariett/arox-pkg/response"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

func (e *endpoint) Gets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := e.stores.Users().Get("test")
		if err != nil {
			response.NotFound(w, "users not found")

			return
		}

		fmt.Println(x)
		//products := []struct {
		//	Brand       string
		//	Name        string
		//	Category    constants.Category
		//	Price       float64
		//	Description string
		//	Gender      *constants.Gender
		//	Season      *constants.Season
		//}{
		//	{
		//		Brand:       "Nike",
		//		Name:        "Air Force",
		//		Category:    constants.Shoes,
		//		Price:       123.45,
		//		Description: "Default white shoes",
		//		Gender:      &constants.Unisex,
		//		Season:      &constants.AllSeason,
		//	},
		//	{
		//		Brand:       "Adidas",
		//		Name:        "Ultraboost",
		//		Category:    constants.Shoes,
		//		Price:       189.99,
		//		Description: "Comfortable running shoes with boost technology",
		//		Gender:      &constants.Unisex,
		//		Season:      &constants.AllSeason,
		//	},
		//	{
		//		Brand:       "Zara",
		//		Name:        "Classic Blazer",
		//		Category:    constants.Clothes,
		//		Price:       89.95,
		//		Description: "Elegant office blazer for business meetings",
		//		Gender:      &constants.Female,
		//		Season:      &constants.Autumn,
		//	},
		//	{
		//		Brand:       "The North Face",
		//		Name:        "Nuptse Jacket",
		//		Category:    constants.Clothes,
		//		Price:       299.00,
		//		Description: "Warm down jacket for extreme cold weather",
		//		Gender:      &constants.Unisex,
		//		Season:      &constants.Autumn,
		//	},
		//	{
		//		Brand:       "Levi's",
		//		Name:        "501 Original",
		//		Category:    constants.Clothes,
		//		Price:       79.50,
		//		Description: "Classic straight fit jeans",
		//		Gender:      &constants.Male,
		//		Season:      &constants.AllSeason,
		//	},
		//	{
		//		Brand:       "Ray-Ban",
		//		Name:        "Aviator",
		//		Category:    constants.Watches,
		//		Price:       159.00,
		//		Description: "Classic aviator style sunglasses",
		//		Gender:      &constants.Unisex,
		//		Season:      nil,
		//	},
		//}

		value, err := e.client.ListProducts(r.Context(), &emptypb.Empty{})
		if err != nil {
			response.NotFound(w, "products not found")

			return
		}

		response.Ok(w, value)
	}
}
