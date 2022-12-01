package main


import (
	"fmt"
	"math/rand"
	"math"
)

var MAX_RANDOM int = 300

type Share struct {
	x float64
	y float64
}

func getShareY(polynomial []float64, x float64) float64 {
	var y float64
	for i := 0; i < len(polynomial); i++ {
		y += polynomial[i] * math.Pow(x, float64(i))
	}
	return y
}

func generatePolynomial(minPersons int, secret int) []float64 {
	polynomial := make([]float64, minPersons)
	polynomial[0] = float64(secret)
	for i := 1; i < minPersons; i++ {
		polynomial[i] = rand.Float64() * float64(MAX_RANDOM)
	}
	return polynomial
}

func recoverSecret(shares []Share) float64 {
	var result float64

	for i, share := range shares {
		var numerator float64 = 1
		var denominator float64 = 1
		for j, share2 := range shares {
			if i != j {
				numerator *= -share2.x
				denominator *= share.x - share2.x
			}
		}
		result = result + (share.y * numerator / denominator)
	}

	return math.Round(result)
}


func main() {
	nbPersons := 5
	nbPersonsMin := 2
	secret := 1914
	fmt.Printf("Secret: %+v\n", secret)

	polynomial := generatePolynomial(nbPersonsMin, secret)

	// Store of shares of the key
	var shares []Share;

	// Generate n shares of the secret
	for i := 0; i < nbPersons; i++ {
		x := rand.Float64()
		share := Share{x, getShareY(polynomial, x)}
		shares = append(shares, share)
		fmt.Printf("Share %d: %+v\n",i,share)

	}

	// Recovered secret from the shares nbPersonsMin shares
	recoveredSecret := recoverSecret(shares[0:nbPersonsMin])
	fmt.Printf("Recovered secret: %+v\n", recoveredSecret)
}