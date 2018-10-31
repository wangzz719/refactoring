package main

import (
	"fmt"
)

const (
	MovieRegular = iota
	MovieChildrens
	MovieNewRelease
)

type Pricer interface {
	getCharge(daysRented int) float64
	getFrequentRenterPoints(daysRented int) int
}

type Price struct {
	code int
}

func (p *Price) getFrequentRenterPoints(daysRented int) int {
	return 1
}

type RegularPrice struct {
	Price
}

func NewRegularPrice() *RegularPrice {
	return &RegularPrice{
		Price: Price{code: MovieRegular},
	}
}

func (p *RegularPrice) getCharge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}

type ChildrensPrice struct {
	Price
}

func NewChildrensPrice() *ChildrensPrice {
	return &ChildrensPrice{
		Price: Price{code: MovieChildrens},
	}
}

func (p *ChildrensPrice) getCharge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

type NewReleasePrice struct {
	Price
}

func NewNewReleasePrice() *NewReleasePrice {
	return &NewReleasePrice{
		Price: Price{code: MovieNewRelease},
	}
}

func (p *NewReleasePrice) getCharge(daysRented int) float64 {
	return float64(daysRented * 3)
}

func (p *NewReleasePrice) getFrequentRenterPoints(daysRented int) int {
	if daysRented > 1 {
		return 2
	}
	return 1
}

type Movie struct {
	title string
	Pricer
}

func NewMovie(title string, price Pricer) *Movie {
	return &Movie{
		title:  title,
		Pricer: price,
	}
}

type Rental struct {
	daysRented int
	movie      *Movie
}

func NewRental(daysRented int, movie *Movie) *Rental {
	return &Rental{
		daysRented: daysRented,
		movie:      movie,
	}
}

func (r *Rental) getCharge() float64 {
	return r.movie.getCharge(r.daysRented)
}

func (r *Rental) getFrequentRenterPoints() int {
	return r.movie.getFrequentRenterPoints(r.daysRented)
}

type Customer struct {
	name    string
	rentals []*Rental
}

func NewCustomer(name string, rentals []*Rental) *Customer {
	return &Customer{
		name:    name,
		rentals: rentals,
	}
}

func (c *Customer) getTotalCharge() float64 {
	result := 0.0
	for _, rental := range c.rentals {
		result += rental.getCharge()
	}
	return result
}

func (c *Customer) getTotalFrequentRenterPoints() int {
	result := 0
	for _, rental := range c.rentals {
		result += rental.getFrequentRenterPoints()
	}
	return result
}

func (c *Customer) statement() string {
	result := fmt.Sprintf("Rental Record for %s \n ", c.name)
	for _, rental := range c.rentals {
		result = fmt.Sprintf("%s \t%s \t%f \n", result, rental.movie.title, rental.getCharge())
	}

	result = fmt.Sprintf("%sAmount owed is %f\nYou earned %d frequent renter points", result, c.getTotalCharge(), c.getTotalFrequentRenterPoints())
	return result
}

func main() {
	regularPrice := NewRegularPrice()
	regularMovie := NewMovie("regular movie", regularPrice)
	regularMovieRental := NewRental(2, regularMovie)

	childrenPrice := NewChildrensPrice()
	childrenMovie := NewMovie("children movie", childrenPrice)
	childrenMovieRental := NewRental(5, childrenMovie)

	newReleasePrice := NewNewReleasePrice()
	newReleaseMovie := NewMovie("new release movie", newReleasePrice)
	newReleaseMovieRental := NewRental(6, newReleaseMovie)

	customer := NewCustomer("wangzz", []*Rental{regularMovieRental, childrenMovieRental, newReleaseMovieRental})
	fmt.Println(customer.statement())
}
