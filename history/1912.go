package main

import (
	"github.com/emirpasic/gods/sets/treeset"
)

type MovieInfo struct {
	Price, Shop int
}

type PriceInfo struct {
	Shop, Movie int
}

type Movie struct {
	Price, Shop, Movie int
}

func MovieInfoComparer(a, b any) int {
	am, bm := a.(MovieInfo), b.(MovieInfo)
	if am.Price == bm.Price {
		return am.Shop - bm.Shop
	}
	return am.Price - bm.Price
}

func MovieComparer(a, b any) int {
	am, bm := a.(Movie), b.(Movie)
	if am.Price == bm.Price {
		if am.Shop == bm.Shop {
			return am.Movie - bm.Movie
		}
		return am.Shop - bm.Shop
	}
	return am.Price - bm.Price
}

type MovieRentingSystem struct {
	Unrented map[int]*treeset.Set // movie to *[]MovieInfo
	Rented   treeset.Set          // []Movie
	PriceMap map[PriceInfo]int
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	mrs := MovieRentingSystem{
		Unrented: make(map[int]*treeset.Set),
		Rented:   *treeset.NewWith(MovieComparer),
		PriceMap: map[PriceInfo]int{},
	}

	for _, entry := range entries {
		shop, movie, price := entry[0], entry[1], entry[2]
		if mrs.Unrented[movie] == nil {
			mrs.Unrented[movie] = treeset.NewWith(MovieInfoComparer)
		}
		mrs.Unrented[movie].Add(MovieInfo{Price: price, Shop: shop})
		mrs.PriceMap[PriceInfo{Shop: shop, Movie: movie}] = price
	}
	return mrs
}

func (mrs *MovieRentingSystem) Search(movie int) []int {
	if mrs.Unrented[movie] == nil {
		return []int{}
	}

	iter := mrs.Unrented[movie].Iterator()
	result := []int{}
	for count := 0; count < 5; count++ {
		if !iter.Next() {
			break
		}
		result = append(result, iter.Value().(MovieInfo).Shop)
	}
	return result
}

func (mrs *MovieRentingSystem) Rent(shop int, movie int) {
	price := mrs.PriceMap[PriceInfo{Movie: movie, Shop: shop}]
	mrs.Unrented[movie].Remove(MovieInfo{Shop: shop, Price: price})
	mrs.Rented.Add(Movie{Price: price, Shop: shop, Movie: movie})
}

func (mrs *MovieRentingSystem) Drop(shop int, movie int) {
	price := mrs.PriceMap[PriceInfo{Movie: movie, Shop: shop}]
	mrs.Rented.Remove(Movie{Price: price, Shop: shop, Movie: movie})
	mrs.Unrented[movie].Add(MovieInfo{Shop: shop, Price: price})
}

func (mrs *MovieRentingSystem) Report() [][]int {
	iter := mrs.Rented.Iterator()
	result := [][]int{}
	for count := 0; count < 5; count++ {
		if !iter.Next() {
			break
		}
		info := iter.Value().(Movie)
		result = append(result, []int{info.Shop, info.Movie})
	}
	return result
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */
