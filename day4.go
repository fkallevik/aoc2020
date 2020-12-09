package main

import (
	"fmt"
	"regexp"
	"strconv"
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

func emptyPassport() passport {
	return passport{
		byr: "",
		iyr: "",
		eyr: "",
		hgt: "",
		hcl: "",
		ecl: "",
		pid: "",
		cid: "",
	}
}

func getRegisteredPassports(input []string) []passport {

	isNewPassport := true

	p := emptyPassport()

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

			p = emptyPassport()
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

	// Append the last parsed passport and return.
	return append(passports, p)
}

func ppFieldsAreEmpty(p passport) bool {
	return p.byr == "" || p.iyr == "" || p.eyr == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == ""
}

func invalidYear(s string, min int, max int) bool {
	i, err := strconv.Atoi(s)

	return err != nil || len(s) != 4 || i < min || i > max
}

func invalidHeight(s string) bool {
	var min, max int

	re := regexp.MustCompile(`^([0-9]{2,3})(cm|in)$`)

	if !re.MatchString(s) {
		return true
	}

	substrings := re.FindStringSubmatch(s)

	if len(substrings) != 3 {
		return true
	}

	val := substrings[1]
	unit := substrings[2]

	if unit == "cm" {
		min = 150
		max = 193
	}

	if unit == "in" {
		min = 59
		max = 76
	}

	h, err := strconv.Atoi(val)

	return err != nil || h < min || h > max
}

func invalidHairColor(s string) bool {
	re := regexp.MustCompile(`#[0-9a-f]{6}$`)

	return !re.MatchString(s)
}

func invalidEyeColor(s string) bool {
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)

	return !re.MatchString(s)
}

func invalidPassportID(s string) bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)

	return !re.MatchString(s)
}

func validatePassports(passports []passport, strict bool) []passport {
	var validated []passport

	for _, p := range passports {
		if ppFieldsAreEmpty(p) {
			continue
		}

		if strict {
			if invalidYear(p.byr, 1920, 2002) {
				continue
			}
			if invalidYear(p.iyr, 2010, 2020) {
				continue
			}
			if invalidYear(p.eyr, 2020, 2030) {
				continue
			}
			if invalidHeight(p.hgt) {
				continue
			}
			if invalidHairColor(p.hcl) {
				continue
			}
			if invalidEyeColor(p.ecl) {
				continue
			}
			if invalidPassportID((p.pid)) {
				continue
			}
		}

		validated = append(validated, p)
	}

	return validated
}

func Day4() {
	input := strings.Split(ReadFile("input/day4"), "\n")

	registeredPassports := getRegisteredPassports(input)

	validPassportsPartOne := validatePassports(registeredPassports, false)
	validPassportsPartTwo := validatePassports(registeredPassports, true)

	fmt.Printf("\nPart One - Valid Passports: %v\n", len(validPassportsPartOne))
	fmt.Printf("\nPart Two - Valid Passports: %v\n", len(validPassportsPartTwo))
}
