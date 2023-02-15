package groupie

import (
	"fmt"
	"strconv"
)

func Filtre(Detail []Detail, DetailFull []Detail, DetailTest []Detail, member int, date int, album int) []Detail {

	if member > 0 {
		if date > 1900 {
			if album > 0 {

				for i := 0; i < len(Detail); i++ {
					stringAlbumDate := Detail[i].FirstAlbum[6:]
					albumDate, _ := strconv.Atoi(stringAlbumDate)
					for y := album; y < album+10; y++ {
						for z := date; z < date+10; z++ {
							if albumDate == y && Detail[i].CreationDate == z && len((Detail)[i].Members) == member {
								fmt.Println((Detail)[i].Name)
								DetailTest = append(DetailTest, (Detail)[i])
							}
						}
					}
				}
				album = 0
				date = 1900
				//fmt.Println(Detail)
				fmt.Println(len(DetailTest))
				DetailFull = DetailTest

			} else {

				for i := 0; i < len(Detail); i++ {
					for y := date; y < date+10; y++ {
						if Detail[i].CreationDate == y && len((Detail)[i].Members) == member {
							fmt.Println((Detail)[i].Name)
							DetailTest = append(DetailTest, (Detail)[i])
						}
					}
				}
				date = 1900
				// fmt.Println(Detail)
				fmt.Println(len(DetailTest))
				DetailFull = DetailTest
			}
		} else if album > 0 {

			for i := 0; i < len(Detail); i++ {
				for y := album; y < album+10; y++ {
					stringAlbumDate := Detail[i].FirstAlbum[6:]
					albumDate, _ := strconv.Atoi(stringAlbumDate)
					if albumDate == y && len((Detail)[i].Members) == member {
						fmt.Println((Detail)[i].Name)
						DetailTest = append(DetailTest, (Detail)[i])
					}
				}
			}
			album = 0
			//fmt.Println(Detail)
			fmt.Println(len(DetailTest))
			DetailFull = DetailTest
		} else {

			for i := 0; i < len(Detail); i++ {
				if len((Detail)[i].Members) == member {
					fmt.Println((Detail)[i].Name)
					DetailTest = append(DetailTest, (Detail)[i])

				}
			}
			//fmt.Println(Detail)
			fmt.Println(len(DetailTest))
			DetailFull = DetailTest
		}
	}

	if date > 1900 {
		if album > 0 {

			for i := 0; i < len(Detail); i++ {
				stringAlbumDate := Detail[i].FirstAlbum[6:]
				albumDate, _ := strconv.Atoi(stringAlbumDate)
				for y := album; y < album+10; y++ {
					for z := date; z < date+10; z++ {
						if albumDate == y && Detail[i].CreationDate == z {
							fmt.Println((Detail)[i].Name)
							DetailTest = append(DetailTest, (Detail)[i])
						}
					}
				}
			}
			album = 0
			//fmt.Println(Detail)
			fmt.Println(len(DetailTest))
			DetailFull = DetailTest
		} else {

			for i := 0; i < len(Detail); i++ {
				for y := date; y < date+10; y++ {
					if Detail[i].CreationDate == y {
						fmt.Println((Detail)[i].Name)
						DetailTest = append(DetailTest, (Detail)[i])
						fmt.Println("test")
					}
				}
			}
			//fmt.Println(Detail)
			fmt.Println(len(DetailTest))
			DetailFull = DetailTest
		}
	}

	if album > 0 {

		for i := 0; i < len(Detail); i++ {
			for y := album; y < album+10; y++ {
				stringAlbumDate := Detail[i].FirstAlbum[6:]
				albumDate, _ := strconv.Atoi(stringAlbumDate)
				if albumDate == y {
					fmt.Println((Detail)[i].Name)
					DetailTest = append(DetailTest, (Detail)[i])
				}
			}
		}
		//fmt.Println(Detail)
		fmt.Println(len(DetailTest))
		DetailFull = DetailTest
	}
	return DetailFull
}
