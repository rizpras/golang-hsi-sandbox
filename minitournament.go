package main

import (
	"fmt"
)

func main() {

	const numPlay = 5

	var totalGoal int
	var numDraw int
	var venueArray [numPlay]string
	var historySlice []string

	pointsMap := make(map[string]int)

	//input venue stadion
	fmt.Println("===================================")
	fmt.Println("     INPUT DATA VENUE STADION")
	fmt.Println("===================================")

	for i := 1; i <= numPlay; i++ {
		fmt.Printf("Masukkan nama stadion untuk laga ke-%d : \n", i)
		fmt.Scan(&venueArray[i-1])
	}

	fmt.Println("===================================")
	fmt.Println("     SISTEM TURNAMEN MINI")
	fmt.Println("===================================")

	for i := 1; i <= numPlay; i++ {

		var homeTeam string
		var awayTeam string
		var homeGoal int
		var awayGoal int

		fmt.Println()
		fmt.Printf("=== Pertandingan Ke-%d di Stadion %s===\n", i, venueArray[i-1])

		fmt.Print("Nama Tim Home : ")
		fmt.Scan(&homeTeam)

		fmt.Print("Nama Tim Away : ")
		fmt.Scan(&awayTeam)

		// Validasi gol home
		for {
			fmt.Print("Gol ", homeTeam, " : ")
			fmt.Scan(&homeGoal)

			if homeGoal >= 0 {
				break
			}

			fmt.Println("ERROR: Gol tidak boleh negatif!")
		}

		// Validasi gol away
		for {
			fmt.Print("Gol ", awayTeam, " : ")
			fmt.Scan(&awayGoal)

			if awayGoal >= 0 {
				break
			}

			fmt.Println("ERROR: Gol tidak boleh negatif!")
		}

		if homeGoal > awayGoal {

			//Tambahkan poin 3 untuk home
			pointsMap[homeTeam] += 3

		} else if awayGoal > homeGoal {

			//Tambahkan poin 3 untuk away
			pointsMap[awayTeam] += 3

		} else {

			pointsMap[homeTeam]++
			pointsMap[awayTeam]++
			numDraw++
		}

		historySlice = append(historySlice, fmt.Sprintf("[%s] %s %d - %d %s --> Pemenang: %s", venueArray[i-1], homeTeam, homeGoal, awayTeam, awayGoal, homeTeam))
		totalGoal += homeGoal + awayGoal
	}

	averageGoal := float64(totalGoal) / float64(numPlay)

	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("     RINGKASAN & KLASEMEN AKHIR")
	fmt.Println("===================================")
	fmt.Println("Daftar Venue Stadion")

	for i := 1; i <= numPlay; i++ {
		fmt.Printf("- Laga %d bertempat di %s\n", i, venueArray[i-1])
	}

	fmt.Println("Riwayat Pertandingan:")
	for i := 0; i < len(historySlice); i++ {
		fmt.Printf("%d. %s", i+1, historySlice[i])
	}

	fmt.Println("Klasemen Poin Akhir:")
	for index, value := range pointsMap {
		fmt.Printf("- %s          : %d Poin\n", index, value)
	}

	fmt.Println("--------------------------------")
	fmt.Println("Jumlah Pertandingan :", numPlay)
	fmt.Println("Total Gol           :", totalGoal)
	fmt.Println("Jumlah Seri         :", numDraw)
	fmt.Printf("Rata-rata Gol       : %.2f\n", averageGoal)
}
