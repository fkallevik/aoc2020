package main

import (
	"fmt"
	"strings"
)

type passport struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

var passports []passport

func Day4() {
	input := strings.Split(ReadFile("input/day4"), "\n")

	isNewPassport := true

	p := passport{
		byr: "", iyr: "", eyr: "", hgt: "", hcl: "", ecl: "", pid: "", cid: "",
	}

	for i, line := range input {

		if line == "" {
			isNewPassport = true
			continue
		}

		if isNewPassport {
			isNewPassport = false

			if i > 0 {
				passports = append(passports, p)
			}

			p = passport{
				byr: "", iyr: "", eyr: "", hgt: "", hcl: "", ecl: "", pid: "", cid: "",
			}
		}

		props := strings.Split(strings.Trim(line, " "), " ")

		for _, prop := range props {
			splitProp := strings.Split(prop, ":")
			key := splitProp[0]
			val := splitProp[1]

			switch key {
			case "byr":
				p.byr = val
			case "iyr":
				p.iyr = val
			case "eyr":
				p.eyr = val
			case "hgt":
				p.hgt = val
			case "hcl":
				p.hcl = val
			case "ecl":
				p.ecl = val
			case "pid":
				p.pid = val
			case "cid":
				p.cid = val
			default:
			}
		}
	}

	// Append the last parsed passport.
	passports = append(passports, p)

	var validPassports []passport

	// Validate passports.
	for _, p := range passports {
		if p.byr == "" || p.iyr == "" || p.eyr == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "" {
			continue
		}

		validPassports = append(validPassports, p)
	}

	fmt.Printf("\nPart One - Valid Passports: %v\n", len(validPassports))
}
