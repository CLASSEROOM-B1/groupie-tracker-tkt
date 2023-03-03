package groupie

func Filtre(Detail []Detail, Vide *Vide, DetailFull []Detail, DetailTest []Detail, member int, date int, album int, europe string, amerique string, afrique string, asie string, oceanie string) []Detail {

	// cont := 0

	// if member > 0 {
	// 	if date > 1901 {
	// 		if album > 1901 {
	// 			if europe != "" {
	// 				if asie != "" {
	// 					if afrique != "" {
	// 						if amerique != "" {
	// 							if oceanie != "" {
	// 								for i := 0; i < len(Vide.Index); i++ {
	// 									for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 										if Vide.Index[i].Continent[y] == 1 {
	// 											fmt.Println((Detail)[i].Name)
	// 											DetailTest = append(DetailTest, (Detail)[i])
	// 											cont = 1
	// 											break
	// 										}
	// 										if cont == 1 {
	// 											cont = 0
	// 											break
	// 										}
	// 									}
	// 								}
	// 								DetailFull = DetailTest
	// 							}
	// 						} else {

	// 							for i := 0; i < len(Vide.Index); i++ {
	// 								for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 									if Vide.Index[i].Continent[y] == 3 {
	// 										fmt.Println((Detail)[i].Name)
	// 										DetailTest = append(DetailTest, (Detail)[i])
	// 										cont = 1
	// 										break
	// 									}
	// 									if cont == 1 {
	// 										cont = 0
	// 										break
	// 									}
	// 								}
	// 							}
	// 							DetailFull = DetailTest
	// 						}
	// 					} else if oceanie != "" {
	// 						for i := 0; i < len(Vide.Index); i++ {
	// 							for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 								if Vide.Index[i].Continent[y] == 5 {
	// 									fmt.Println((Detail)[i].Name)
	// 									DetailTest = append(DetailTest, (Detail)[i])
	// 									cont = 1
	// 									break
	// 								}
	// 								if cont == 1 {
	// 									cont = 0
	// 									break
	// 								}
	// 							}
	// 						}
	// 					} else {
	// 						for i := 0; i < len(Vide.Index); i++ {
	// 							for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 								if Vide.Index[i].Continent[y] == 4 {
	// 									fmt.Println((Detail)[i].Name)
	// 									DetailTest = append(DetailTest, (Detail)[i])
	// 									cont = 1
	// 									break
	// 								}
	// 								if cont == 1 {
	// 									cont = 0
	// 									break
	// 								}
	// 							}
	// 						}
	// 						DetailFull = DetailTest
	// 					}
	// 				} else if amerique != "" {
	// 					if oceanie != "" {
	// 						for i := 0; i < len(Vide.Index); i++ {
	// 							for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 								if Vide.Index[i].Continent[y] == 1 {
	// 									fmt.Println((Detail)[i].Name)
	// 									DetailTest = append(DetailTest, (Detail)[i])
	// 									cont = 1
	// 									break
	// 								}
	// 								if cont == 1 {
	// 									cont = 0
	// 									break
	// 								}
	// 							}
	// 						}
	// 						DetailFull = DetailTest
	// 					} else {

	// 					for i := 0; i < len(Vide.Index); i++ {
	// 						for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 							if Vide.Index[i].Continent[y] == 3 {
	// 								fmt.Println((Detail)[i].Name)
	// 								DetailTest = append(DetailTest, (Detail)[i])
	// 								cont = 1
	// 								break
	// 							}
	// 							if cont == 1 {
	// 								cont = 0
	// 								break
	// 							}
	// 						}
	// 					}
	// 					DetailFull = DetailTest
	// 					}

	// 				} else {
	// 					for i := 0; i < len(Vide.Index); i++ {
	// 						for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 							if Vide.Index[i].Continent[y] == 2 {
	// 								fmt.Println((Detail)[i].Name)
	// 								DetailTest = append(DetailTest, (Detail)[i])
	// 								cont = 1
	// 								break
	// 							}
	// 							if cont == 1 {
	// 								cont = 0
	// 								break
	// 							}
	// 						}
	// 					}
	// 					DetailFull = DetailTest
	// 				}
	// 			}  else if asie != "" {
	// 				if afrique != "" {
	// 					if amerique != "" {
	// 						if oceanie != "" {
	// 							for i := 0; i < len(Vide.Index); i++ {
	// 								for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 									if Vide.Index[i].Continent[y] == 1 {
	// 										fmt.Println((Detail)[i].Name)
	// 										DetailTest = append(DetailTest, (Detail)[i])
	// 										cont = 1
	// 										break
	// 									}
	// 									if cont == 1 {
	// 										cont = 0
	// 										break
	// 									}
	// 								}
	// 							}
	// 							DetailFull = DetailTest
	// 						}
	// 					} else {

	// 						for i := 0; i < len(Vide.Index); i++ {
	// 							for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 								if Vide.Index[i].Continent[y] == 3 {
	// 									fmt.Println((Detail)[i].Name)
	// 									DetailTest = append(DetailTest, (Detail)[i])
	// 									cont = 1
	// 									break
	// 								}
	// 								if cont == 1 {
	// 									cont = 0
	// 									break
	// 								}
	// 							}
	// 						}
	// 						DetailFull = DetailTest
	// 					}
	// 				} else if oceanie != "" {
	// 					for i := 0; i < len(Vide.Index); i++ {
	// 						for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 							if Vide.Index[i].Continent[y] == 5 {
	// 								fmt.Println((Detail)[i].Name)
	// 								DetailTest = append(DetailTest, (Detail)[i])
	// 								cont = 1
	// 								break
	// 							}
	// 							if cont == 1 {
	// 								cont = 0
	// 								break
	// 							}
	// 						}
	// 					}
	// 				} else {
	// 					for i := 0; i < len(Vide.Index); i++ {
	// 						for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 							if Vide.Index[i].Continent[y] == 4 {
	// 								fmt.Println((Detail)[i].Name)
	// 								DetailTest = append(DetailTest, (Detail)[i])
	// 								cont = 1
	// 								break
	// 							}
	// 							if cont == 1 {
	// 								cont = 0
	// 								break
	// 							}
	// 						}
	// 					}
	// 					DetailFull = DetailTest
	// 				}
	// 			} else if amerique != "" {
	// 				if oceanie != "" {
	// 					for i := 0; i < len(Vide.Index); i++ {
	// 						for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 							if Vide.Index[i].Continent[y] == 1 {
	// 								fmt.Println((Detail)[i].Name)
	// 								DetailTest = append(DetailTest, (Detail)[i])
	// 								cont = 1
	// 								break
	// 							}
	// 							if cont == 1 {
	// 								cont = 0
	// 								break
	// 							}
	// 						}
	// 					}
	// 					DetailFull = DetailTest
	// 				} else {

	// 				for i := 0; i < len(Vide.Index); i++ {
	// 					for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 						if Vide.Index[i].Continent[y] == 3 {
	// 							fmt.Println((Detail)[i].Name)
	// 							DetailTest = append(DetailTest, (Detail)[i])
	// 							cont = 1
	// 							break
	// 						}
	// 						if cont == 1 {
	// 							cont = 0
	// 							break
	// 						}
	// 					}
	// 				}
	// 				DetailFull = DetailTest
	// 			}

	// 			} else {
	// 				for i := 0; i < len(Vide.Index); i++ {
	// 					for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 						if Vide.Index[i].Continent[y] == 2 {
	// 							fmt.Println((Detail)[i].Name)
	// 							DetailTest = append(DetailTest, (Detail)[i])
	// 							cont = 1
	// 							break
	// 						}
	// 						if cont == 1 {
	// 							cont = 0
	// 							break
	// 						}
	// 					}
	// 				}
	// 				DetailFull = DetailTest
	// 			}
	// 			} else {
	// 				for i := 0; i < len(Vide.Index); i++ {
	// 					for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 						if Vide.Index[i].Continent[y] == 1 {
	// 							fmt.Println((Detail)[i].Name)
	// 							DetailTest = append(DetailTest, (Detail)[i])
	// 							cont = 1
	// 							break
	// 						}
	// 						if cont == 1 {
	// 							cont = 0
	// 							break
	// 						}
	// 					}
	// 				}
	// 				DetailFull = DetailTest
	// 			}
	// 			}

	// 			for i := 0; i < len(Detail); i++ {
	// 				stringAlbumDate := Detail[i].FirstAlbum[6:]
	// 				albumDate, _ := strconv.Atoi(stringAlbumDate)
	// 				for y := album; y < album+10; y++ {
	// 					for z := date; z < date+10; z++ {
	// 						if albumDate == y && Detail[i].CreationDate == z && len((Detail)[i].Members) == member {
	// 							DetailTest = append(DetailTest, (Detail)[i])
	// 						}
	// 					}
	// 				}
	// 			}
	// 			album = 0
	// 			date = 1900
	// 			DetailFull = DetailTest

	// 		} else {

	// 			for i := 0; i < len(Detail); i++ {
	// 				for y := date; y < date+10; y++ {
	// 					if Detail[i].CreationDate == y && len((Detail)[i].Members) == member {
	// 						DetailTest = append(DetailTest, (Detail)[i])
	// 					}
	// 				}
	// 			}
	// 			date = 1900

	// 			DetailFull = DetailTest
	// 		}
	// 	} else if album > 1901 {

	// 		for i := 0; i < len(Detail); i++ {
	// 			for y := album; y < album+10; y++ {
	// 				stringAlbumDate := Detail[i].FirstAlbum[6:]
	// 				albumDate, _ := strconv.Atoi(stringAlbumDate)
	// 				if albumDate == y && len((Detail)[i].Members) == member {
	// 					DetailTest = append(DetailTest, (Detail)[i])
	// 				}
	// 			}
	// 		}
	// 		album = 0
	// 		DetailFull = DetailTest
	// 	} else {

	// 		for i := 0; i < len(Detail); i++ {
	// 			if len((Detail)[i].Members) == member {
	// 				DetailTest = append(DetailTest, (Detail)[i])

	// 			}
	// 		}
	// 		DetailFull = DetailTest
	// 	}
	// }

	// if date > 1901 {
	// 	if album > 1901 {

	// 		for i := 0; i < len(Detail); i++ {
	// 			stringAlbumDate := Detail[i].FirstAlbum[6:]
	// 			albumDate, _ := strconv.Atoi(stringAlbumDate)
	// 			for y := album; y < album+10; y++ {
	// 				for z := date; z < date+10; z++ {
	// 					if albumDate == y && Detail[i].CreationDate == z {
	// 						DetailTest = append(DetailTest, (Detail)[i])
	// 					}
	// 				}
	// 			}
	// 		}
	// 		album = 0
	// 		DetailFull = DetailTest
	// 	} else {

	// 		for i := 0; i < len(Detail); i++ {
	// 			for y := date; y < date+10; y++ {
	// 				if Detail[i].CreationDate == y {
	// 					DetailTest = append(DetailTest, (Detail)[i])
	// 				}
	// 			}
	// 		}
	// 		DetailFull = DetailTest
	// 	}
	// }

	// if album > 1901 {

	// 	for i := 0; i < len(Detail); i++ {
	// 		for y := album; y < album+10; y++ {
	// 			stringAlbumDate := Detail[i].FirstAlbum[6:]
	// 			albumDate, _ := strconv.Atoi(stringAlbumDate)
	// 			if albumDate == y {
	// 				DetailTest = append(DetailTest, (Detail)[i])
	// 			}
	// 		}
	// 	}
	// 	DetailFull = DetailTest
	// }

	// if europe != "" {
	// 	for i := 0; i < len(Vide.Index); i++ {
	// 		for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 			if Vide.Index[i].Continent[y] == 1 {
	// 				fmt.Println((Detail)[i].Name)
	// 				DetailTest = append(DetailTest, (Detail)[i])
	// 			}
	// 		}
	// 	}
	// 	DetailFull = DetailTest
	// }

	// xd := 0

	// if asie != "" {
	// 	for i := 0; i < len(Vide.Index); i++ {
	// 		for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 			if Vide.Index[i].Continent[y] == 2 {
	// 				fmt.Println((Detail)[i].Name)
	// 				DetailTest = append(DetailTest, (Detail)[i])
	// 				xd = 1
	// 				break
	// 			}
	// 			if xd == 1 {
	// 				break
	// 			}
	// 		}
	// 	}
	// 	DetailFull = DetailTest
	// }

	// if afrique != "" {
	// 	for i := 0; i < len(Vide.Index); i++ {
	// 		for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 			if Vide.Index[i].Continent[y] == 3 {
	// 				fmt.Println((Detail)[i].Name)
	// 				DetailTest = append(DetailTest, (Detail)[i])

	// 			}
	// 		}
	// 	}
	// 	DetailFull = DetailTest
	// }

	// if amerique != "" {
	// 	for i := 0; i < len(Vide.Index); i++ {
	// 		for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 			if Vide.Index[i].Continent[y] == 4 {
	// 				fmt.Println((Detail)[i].Name)
	// 				DetailTest = append(DetailTest, (Detail)[i])
	// 			}
	// 		}
	// 	}
	// 	DetailFull = DetailTest
	// }
	// if oceanie != "" {
	// 	for i := 0; i < len(Vide.Index); i++ {
	// 		for y := 0; y < len(Vide.Index[i].Continent); y++ {
	// 			if Vide.Index[i].Continent[y] == 6 {
	// 				fmt.Println((Detail)[i].Name)
	// 				DetailTest = append(DetailTest, (Detail)[i])
	// 			}
	// 		}
	// 	}
	// 	DetailFull = DetailTest
	// }

	return DetailFull
}
