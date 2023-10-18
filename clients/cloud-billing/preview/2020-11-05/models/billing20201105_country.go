// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Billing20201105Country Country is a two-letter country code as defined by the ISO 3166-1 standard.
//
// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
//
// swagger:model billing_20201105Country
type Billing20201105Country string

func NewBilling20201105Country(value Billing20201105Country) *Billing20201105Country {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Billing20201105Country.
func (m Billing20201105Country) Pointer() *Billing20201105Country {
	return &m
}

const (

	// Billing20201105CountryUNSET captures enum value "UNSET"
	Billing20201105CountryUNSET Billing20201105Country = "UNSET"

	// Billing20201105CountryAF captures enum value "AF"
	Billing20201105CountryAF Billing20201105Country = "AF"

	// Billing20201105CountryAL captures enum value "AL"
	Billing20201105CountryAL Billing20201105Country = "AL"

	// Billing20201105CountryAQ captures enum value "AQ"
	Billing20201105CountryAQ Billing20201105Country = "AQ"

	// Billing20201105CountryDZ captures enum value "DZ"
	Billing20201105CountryDZ Billing20201105Country = "DZ"

	// Billing20201105CountryAS captures enum value "AS"
	Billing20201105CountryAS Billing20201105Country = "AS"

	// Billing20201105CountryAD captures enum value "AD"
	Billing20201105CountryAD Billing20201105Country = "AD"

	// Billing20201105CountryAO captures enum value "AO"
	Billing20201105CountryAO Billing20201105Country = "AO"

	// Billing20201105CountryAG captures enum value "AG"
	Billing20201105CountryAG Billing20201105Country = "AG"

	// Billing20201105CountryAZ captures enum value "AZ"
	Billing20201105CountryAZ Billing20201105Country = "AZ"

	// Billing20201105CountryAR captures enum value "AR"
	Billing20201105CountryAR Billing20201105Country = "AR"

	// Billing20201105CountryAU captures enum value "AU"
	Billing20201105CountryAU Billing20201105Country = "AU"

	// Billing20201105CountryAT captures enum value "AT"
	Billing20201105CountryAT Billing20201105Country = "AT"

	// Billing20201105CountryBS captures enum value "BS"
	Billing20201105CountryBS Billing20201105Country = "BS"

	// Billing20201105CountryBH captures enum value "BH"
	Billing20201105CountryBH Billing20201105Country = "BH"

	// Billing20201105CountryBD captures enum value "BD"
	Billing20201105CountryBD Billing20201105Country = "BD"

	// Billing20201105CountryAM captures enum value "AM"
	Billing20201105CountryAM Billing20201105Country = "AM"

	// Billing20201105CountryBB captures enum value "BB"
	Billing20201105CountryBB Billing20201105Country = "BB"

	// Billing20201105CountryBE captures enum value "BE"
	Billing20201105CountryBE Billing20201105Country = "BE"

	// Billing20201105CountryBM captures enum value "BM"
	Billing20201105CountryBM Billing20201105Country = "BM"

	// Billing20201105CountryBT captures enum value "BT"
	Billing20201105CountryBT Billing20201105Country = "BT"

	// Billing20201105CountryBO captures enum value "BO"
	Billing20201105CountryBO Billing20201105Country = "BO"

	// Billing20201105CountryBA captures enum value "BA"
	Billing20201105CountryBA Billing20201105Country = "BA"

	// Billing20201105CountryBW captures enum value "BW"
	Billing20201105CountryBW Billing20201105Country = "BW"

	// Billing20201105CountryBV captures enum value "BV"
	Billing20201105CountryBV Billing20201105Country = "BV"

	// Billing20201105CountryBR captures enum value "BR"
	Billing20201105CountryBR Billing20201105Country = "BR"

	// Billing20201105CountryBZ captures enum value "BZ"
	Billing20201105CountryBZ Billing20201105Country = "BZ"

	// Billing20201105CountryIO captures enum value "IO"
	Billing20201105CountryIO Billing20201105Country = "IO"

	// Billing20201105CountrySB captures enum value "SB"
	Billing20201105CountrySB Billing20201105Country = "SB"

	// Billing20201105CountryVG captures enum value "VG"
	Billing20201105CountryVG Billing20201105Country = "VG"

	// Billing20201105CountryBN captures enum value "BN"
	Billing20201105CountryBN Billing20201105Country = "BN"

	// Billing20201105CountryBG captures enum value "BG"
	Billing20201105CountryBG Billing20201105Country = "BG"

	// Billing20201105CountryMM captures enum value "MM"
	Billing20201105CountryMM Billing20201105Country = "MM"

	// Billing20201105CountryBI captures enum value "BI"
	Billing20201105CountryBI Billing20201105Country = "BI"

	// Billing20201105CountryBY captures enum value "BY"
	Billing20201105CountryBY Billing20201105Country = "BY"

	// Billing20201105CountryKH captures enum value "KH"
	Billing20201105CountryKH Billing20201105Country = "KH"

	// Billing20201105CountryCM captures enum value "CM"
	Billing20201105CountryCM Billing20201105Country = "CM"

	// Billing20201105CountryCA captures enum value "CA"
	Billing20201105CountryCA Billing20201105Country = "CA"

	// Billing20201105CountryCV captures enum value "CV"
	Billing20201105CountryCV Billing20201105Country = "CV"

	// Billing20201105CountryKY captures enum value "KY"
	Billing20201105CountryKY Billing20201105Country = "KY"

	// Billing20201105CountryCF captures enum value "CF"
	Billing20201105CountryCF Billing20201105Country = "CF"

	// Billing20201105CountryLK captures enum value "LK"
	Billing20201105CountryLK Billing20201105Country = "LK"

	// Billing20201105CountryTD captures enum value "TD"
	Billing20201105CountryTD Billing20201105Country = "TD"

	// Billing20201105CountryCL captures enum value "CL"
	Billing20201105CountryCL Billing20201105Country = "CL"

	// Billing20201105CountryCN captures enum value "CN"
	Billing20201105CountryCN Billing20201105Country = "CN"

	// Billing20201105CountryTW captures enum value "TW"
	Billing20201105CountryTW Billing20201105Country = "TW"

	// Billing20201105CountryCX captures enum value "CX"
	Billing20201105CountryCX Billing20201105Country = "CX"

	// Billing20201105CountryCC captures enum value "CC"
	Billing20201105CountryCC Billing20201105Country = "CC"

	// Billing20201105CountryCO captures enum value "CO"
	Billing20201105CountryCO Billing20201105Country = "CO"

	// Billing20201105CountryKM captures enum value "KM"
	Billing20201105CountryKM Billing20201105Country = "KM"

	// Billing20201105CountryYT captures enum value "YT"
	Billing20201105CountryYT Billing20201105Country = "YT"

	// Billing20201105CountryCG captures enum value "CG"
	Billing20201105CountryCG Billing20201105Country = "CG"

	// Billing20201105CountryCD captures enum value "CD"
	Billing20201105CountryCD Billing20201105Country = "CD"

	// Billing20201105CountryCK captures enum value "CK"
	Billing20201105CountryCK Billing20201105Country = "CK"

	// Billing20201105CountryCR captures enum value "CR"
	Billing20201105CountryCR Billing20201105Country = "CR"

	// Billing20201105CountryHR captures enum value "HR"
	Billing20201105CountryHR Billing20201105Country = "HR"

	// Billing20201105CountryCU captures enum value "CU"
	Billing20201105CountryCU Billing20201105Country = "CU"

	// Billing20201105CountryCY captures enum value "CY"
	Billing20201105CountryCY Billing20201105Country = "CY"

	// Billing20201105CountryCZ captures enum value "CZ"
	Billing20201105CountryCZ Billing20201105Country = "CZ"

	// Billing20201105CountryBJ captures enum value "BJ"
	Billing20201105CountryBJ Billing20201105Country = "BJ"

	// Billing20201105CountryDK captures enum value "DK"
	Billing20201105CountryDK Billing20201105Country = "DK"

	// Billing20201105CountryDM captures enum value "DM"
	Billing20201105CountryDM Billing20201105Country = "DM"

	// Billing20201105CountryDO captures enum value "DO"
	Billing20201105CountryDO Billing20201105Country = "DO"

	// Billing20201105CountryEC captures enum value "EC"
	Billing20201105CountryEC Billing20201105Country = "EC"

	// Billing20201105CountrySV captures enum value "SV"
	Billing20201105CountrySV Billing20201105Country = "SV"

	// Billing20201105CountryGQ captures enum value "GQ"
	Billing20201105CountryGQ Billing20201105Country = "GQ"

	// Billing20201105CountryET captures enum value "ET"
	Billing20201105CountryET Billing20201105Country = "ET"

	// Billing20201105CountryER captures enum value "ER"
	Billing20201105CountryER Billing20201105Country = "ER"

	// Billing20201105CountryEE captures enum value "EE"
	Billing20201105CountryEE Billing20201105Country = "EE"

	// Billing20201105CountryFO captures enum value "FO"
	Billing20201105CountryFO Billing20201105Country = "FO"

	// Billing20201105CountryFK captures enum value "FK"
	Billing20201105CountryFK Billing20201105Country = "FK"

	// Billing20201105CountryGS captures enum value "GS"
	Billing20201105CountryGS Billing20201105Country = "GS"

	// Billing20201105CountryFJ captures enum value "FJ"
	Billing20201105CountryFJ Billing20201105Country = "FJ"

	// Billing20201105CountryFI captures enum value "FI"
	Billing20201105CountryFI Billing20201105Country = "FI"

	// Billing20201105CountryAX captures enum value "AX"
	Billing20201105CountryAX Billing20201105Country = "AX"

	// Billing20201105CountryFR captures enum value "FR"
	Billing20201105CountryFR Billing20201105Country = "FR"

	// Billing20201105CountryGF captures enum value "GF"
	Billing20201105CountryGF Billing20201105Country = "GF"

	// Billing20201105CountryPF captures enum value "PF"
	Billing20201105CountryPF Billing20201105Country = "PF"

	// Billing20201105CountryTF captures enum value "TF"
	Billing20201105CountryTF Billing20201105Country = "TF"

	// Billing20201105CountryDJ captures enum value "DJ"
	Billing20201105CountryDJ Billing20201105Country = "DJ"

	// Billing20201105CountryGA captures enum value "GA"
	Billing20201105CountryGA Billing20201105Country = "GA"

	// Billing20201105CountryGE captures enum value "GE"
	Billing20201105CountryGE Billing20201105Country = "GE"

	// Billing20201105CountryGM captures enum value "GM"
	Billing20201105CountryGM Billing20201105Country = "GM"

	// Billing20201105CountryPS captures enum value "PS"
	Billing20201105CountryPS Billing20201105Country = "PS"

	// Billing20201105CountryDE captures enum value "DE"
	Billing20201105CountryDE Billing20201105Country = "DE"

	// Billing20201105CountryGH captures enum value "GH"
	Billing20201105CountryGH Billing20201105Country = "GH"

	// Billing20201105CountryGI captures enum value "GI"
	Billing20201105CountryGI Billing20201105Country = "GI"

	// Billing20201105CountryKI captures enum value "KI"
	Billing20201105CountryKI Billing20201105Country = "KI"

	// Billing20201105CountryGR captures enum value "GR"
	Billing20201105CountryGR Billing20201105Country = "GR"

	// Billing20201105CountryGL captures enum value "GL"
	Billing20201105CountryGL Billing20201105Country = "GL"

	// Billing20201105CountryGD captures enum value "GD"
	Billing20201105CountryGD Billing20201105Country = "GD"

	// Billing20201105CountryGP captures enum value "GP"
	Billing20201105CountryGP Billing20201105Country = "GP"

	// Billing20201105CountryGU captures enum value "GU"
	Billing20201105CountryGU Billing20201105Country = "GU"

	// Billing20201105CountryGT captures enum value "GT"
	Billing20201105CountryGT Billing20201105Country = "GT"

	// Billing20201105CountryGN captures enum value "GN"
	Billing20201105CountryGN Billing20201105Country = "GN"

	// Billing20201105CountryGY captures enum value "GY"
	Billing20201105CountryGY Billing20201105Country = "GY"

	// Billing20201105CountryHT captures enum value "HT"
	Billing20201105CountryHT Billing20201105Country = "HT"

	// Billing20201105CountryHM captures enum value "HM"
	Billing20201105CountryHM Billing20201105Country = "HM"

	// Billing20201105CountryVA captures enum value "VA"
	Billing20201105CountryVA Billing20201105Country = "VA"

	// Billing20201105CountryHN captures enum value "HN"
	Billing20201105CountryHN Billing20201105Country = "HN"

	// Billing20201105CountryHK captures enum value "HK"
	Billing20201105CountryHK Billing20201105Country = "HK"

	// Billing20201105CountryHU captures enum value "HU"
	Billing20201105CountryHU Billing20201105Country = "HU"

	// Billing20201105CountryIS captures enum value "IS"
	Billing20201105CountryIS Billing20201105Country = "IS"

	// Billing20201105CountryIN captures enum value "IN"
	Billing20201105CountryIN Billing20201105Country = "IN"

	// Billing20201105CountryID captures enum value "ID"
	Billing20201105CountryID Billing20201105Country = "ID"

	// Billing20201105CountryIR captures enum value "IR"
	Billing20201105CountryIR Billing20201105Country = "IR"

	// Billing20201105CountryIQ captures enum value "IQ"
	Billing20201105CountryIQ Billing20201105Country = "IQ"

	// Billing20201105CountryIE captures enum value "IE"
	Billing20201105CountryIE Billing20201105Country = "IE"

	// Billing20201105CountryIL captures enum value "IL"
	Billing20201105CountryIL Billing20201105Country = "IL"

	// Billing20201105CountryIT captures enum value "IT"
	Billing20201105CountryIT Billing20201105Country = "IT"

	// Billing20201105CountryCI captures enum value "CI"
	Billing20201105CountryCI Billing20201105Country = "CI"

	// Billing20201105CountryJM captures enum value "JM"
	Billing20201105CountryJM Billing20201105Country = "JM"

	// Billing20201105CountryJP captures enum value "JP"
	Billing20201105CountryJP Billing20201105Country = "JP"

	// Billing20201105CountryKZ captures enum value "KZ"
	Billing20201105CountryKZ Billing20201105Country = "KZ"

	// Billing20201105CountryJO captures enum value "JO"
	Billing20201105CountryJO Billing20201105Country = "JO"

	// Billing20201105CountryKE captures enum value "KE"
	Billing20201105CountryKE Billing20201105Country = "KE"

	// Billing20201105CountryKP captures enum value "KP"
	Billing20201105CountryKP Billing20201105Country = "KP"

	// Billing20201105CountryKR captures enum value "KR"
	Billing20201105CountryKR Billing20201105Country = "KR"

	// Billing20201105CountryKW captures enum value "KW"
	Billing20201105CountryKW Billing20201105Country = "KW"

	// Billing20201105CountryKG captures enum value "KG"
	Billing20201105CountryKG Billing20201105Country = "KG"

	// Billing20201105CountryLA captures enum value "LA"
	Billing20201105CountryLA Billing20201105Country = "LA"

	// Billing20201105CountryLB captures enum value "LB"
	Billing20201105CountryLB Billing20201105Country = "LB"

	// Billing20201105CountryLS captures enum value "LS"
	Billing20201105CountryLS Billing20201105Country = "LS"

	// Billing20201105CountryLV captures enum value "LV"
	Billing20201105CountryLV Billing20201105Country = "LV"

	// Billing20201105CountryLR captures enum value "LR"
	Billing20201105CountryLR Billing20201105Country = "LR"

	// Billing20201105CountryLY captures enum value "LY"
	Billing20201105CountryLY Billing20201105Country = "LY"

	// Billing20201105CountryLI captures enum value "LI"
	Billing20201105CountryLI Billing20201105Country = "LI"

	// Billing20201105CountryLT captures enum value "LT"
	Billing20201105CountryLT Billing20201105Country = "LT"

	// Billing20201105CountryLU captures enum value "LU"
	Billing20201105CountryLU Billing20201105Country = "LU"

	// Billing20201105CountryMO captures enum value "MO"
	Billing20201105CountryMO Billing20201105Country = "MO"

	// Billing20201105CountryMG captures enum value "MG"
	Billing20201105CountryMG Billing20201105Country = "MG"

	// Billing20201105CountryMW captures enum value "MW"
	Billing20201105CountryMW Billing20201105Country = "MW"

	// Billing20201105CountryMY captures enum value "MY"
	Billing20201105CountryMY Billing20201105Country = "MY"

	// Billing20201105CountryMV captures enum value "MV"
	Billing20201105CountryMV Billing20201105Country = "MV"

	// Billing20201105CountryML captures enum value "ML"
	Billing20201105CountryML Billing20201105Country = "ML"

	// Billing20201105CountryMT captures enum value "MT"
	Billing20201105CountryMT Billing20201105Country = "MT"

	// Billing20201105CountryMQ captures enum value "MQ"
	Billing20201105CountryMQ Billing20201105Country = "MQ"

	// Billing20201105CountryMR captures enum value "MR"
	Billing20201105CountryMR Billing20201105Country = "MR"

	// Billing20201105CountryMU captures enum value "MU"
	Billing20201105CountryMU Billing20201105Country = "MU"

	// Billing20201105CountryMX captures enum value "MX"
	Billing20201105CountryMX Billing20201105Country = "MX"

	// Billing20201105CountryMC captures enum value "MC"
	Billing20201105CountryMC Billing20201105Country = "MC"

	// Billing20201105CountryMN captures enum value "MN"
	Billing20201105CountryMN Billing20201105Country = "MN"

	// Billing20201105CountryMD captures enum value "MD"
	Billing20201105CountryMD Billing20201105Country = "MD"

	// Billing20201105CountryME captures enum value "ME"
	Billing20201105CountryME Billing20201105Country = "ME"

	// Billing20201105CountryMS captures enum value "MS"
	Billing20201105CountryMS Billing20201105Country = "MS"

	// Billing20201105CountryMA captures enum value "MA"
	Billing20201105CountryMA Billing20201105Country = "MA"

	// Billing20201105CountryMZ captures enum value "MZ"
	Billing20201105CountryMZ Billing20201105Country = "MZ"

	// Billing20201105CountryOM captures enum value "OM"
	Billing20201105CountryOM Billing20201105Country = "OM"

	// Billing20201105CountryNA captures enum value "NA"
	Billing20201105CountryNA Billing20201105Country = "NA"

	// Billing20201105CountryNR captures enum value "NR"
	Billing20201105CountryNR Billing20201105Country = "NR"

	// Billing20201105CountryNP captures enum value "NP"
	Billing20201105CountryNP Billing20201105Country = "NP"

	// Billing20201105CountryNL captures enum value "NL"
	Billing20201105CountryNL Billing20201105Country = "NL"

	// Billing20201105CountryCW captures enum value "CW"
	Billing20201105CountryCW Billing20201105Country = "CW"

	// Billing20201105CountryAW captures enum value "AW"
	Billing20201105CountryAW Billing20201105Country = "AW"

	// Billing20201105CountrySX captures enum value "SX"
	Billing20201105CountrySX Billing20201105Country = "SX"

	// Billing20201105CountryBQ captures enum value "BQ"
	Billing20201105CountryBQ Billing20201105Country = "BQ"

	// Billing20201105CountryNC captures enum value "NC"
	Billing20201105CountryNC Billing20201105Country = "NC"

	// Billing20201105CountryVU captures enum value "VU"
	Billing20201105CountryVU Billing20201105Country = "VU"

	// Billing20201105CountryNZ captures enum value "NZ"
	Billing20201105CountryNZ Billing20201105Country = "NZ"

	// Billing20201105CountryNI captures enum value "NI"
	Billing20201105CountryNI Billing20201105Country = "NI"

	// Billing20201105CountryNE captures enum value "NE"
	Billing20201105CountryNE Billing20201105Country = "NE"

	// Billing20201105CountryNG captures enum value "NG"
	Billing20201105CountryNG Billing20201105Country = "NG"

	// Billing20201105CountryNU captures enum value "NU"
	Billing20201105CountryNU Billing20201105Country = "NU"

	// Billing20201105CountryNF captures enum value "NF"
	Billing20201105CountryNF Billing20201105Country = "NF"

	// Billing20201105CountryNO captures enum value "NO"
	Billing20201105CountryNO Billing20201105Country = "NO"

	// Billing20201105CountryMP captures enum value "MP"
	Billing20201105CountryMP Billing20201105Country = "MP"

	// Billing20201105CountryUM captures enum value "UM"
	Billing20201105CountryUM Billing20201105Country = "UM"

	// Billing20201105CountryFM captures enum value "FM"
	Billing20201105CountryFM Billing20201105Country = "FM"

	// Billing20201105CountryMH captures enum value "MH"
	Billing20201105CountryMH Billing20201105Country = "MH"

	// Billing20201105CountryPW captures enum value "PW"
	Billing20201105CountryPW Billing20201105Country = "PW"

	// Billing20201105CountryPK captures enum value "PK"
	Billing20201105CountryPK Billing20201105Country = "PK"

	// Billing20201105CountryPA captures enum value "PA"
	Billing20201105CountryPA Billing20201105Country = "PA"

	// Billing20201105CountryPG captures enum value "PG"
	Billing20201105CountryPG Billing20201105Country = "PG"

	// Billing20201105CountryPY captures enum value "PY"
	Billing20201105CountryPY Billing20201105Country = "PY"

	// Billing20201105CountryPE captures enum value "PE"
	Billing20201105CountryPE Billing20201105Country = "PE"

	// Billing20201105CountryPH captures enum value "PH"
	Billing20201105CountryPH Billing20201105Country = "PH"

	// Billing20201105CountryPN captures enum value "PN"
	Billing20201105CountryPN Billing20201105Country = "PN"

	// Billing20201105CountryPL captures enum value "PL"
	Billing20201105CountryPL Billing20201105Country = "PL"

	// Billing20201105CountryPT captures enum value "PT"
	Billing20201105CountryPT Billing20201105Country = "PT"

	// Billing20201105CountryGW captures enum value "GW"
	Billing20201105CountryGW Billing20201105Country = "GW"

	// Billing20201105CountryTL captures enum value "TL"
	Billing20201105CountryTL Billing20201105Country = "TL"

	// Billing20201105CountryPR captures enum value "PR"
	Billing20201105CountryPR Billing20201105Country = "PR"

	// Billing20201105CountryQA captures enum value "QA"
	Billing20201105CountryQA Billing20201105Country = "QA"

	// Billing20201105CountryRE captures enum value "RE"
	Billing20201105CountryRE Billing20201105Country = "RE"

	// Billing20201105CountryRO captures enum value "RO"
	Billing20201105CountryRO Billing20201105Country = "RO"

	// Billing20201105CountryRU captures enum value "RU"
	Billing20201105CountryRU Billing20201105Country = "RU"

	// Billing20201105CountryRW captures enum value "RW"
	Billing20201105CountryRW Billing20201105Country = "RW"

	// Billing20201105CountryBL captures enum value "BL"
	Billing20201105CountryBL Billing20201105Country = "BL"

	// Billing20201105CountrySH captures enum value "SH"
	Billing20201105CountrySH Billing20201105Country = "SH"

	// Billing20201105CountryKN captures enum value "KN"
	Billing20201105CountryKN Billing20201105Country = "KN"

	// Billing20201105CountryAI captures enum value "AI"
	Billing20201105CountryAI Billing20201105Country = "AI"

	// Billing20201105CountryLC captures enum value "LC"
	Billing20201105CountryLC Billing20201105Country = "LC"

	// Billing20201105CountryMF captures enum value "MF"
	Billing20201105CountryMF Billing20201105Country = "MF"

	// Billing20201105CountryPM captures enum value "PM"
	Billing20201105CountryPM Billing20201105Country = "PM"

	// Billing20201105CountryVC captures enum value "VC"
	Billing20201105CountryVC Billing20201105Country = "VC"

	// Billing20201105CountrySM captures enum value "SM"
	Billing20201105CountrySM Billing20201105Country = "SM"

	// Billing20201105CountryST captures enum value "ST"
	Billing20201105CountryST Billing20201105Country = "ST"

	// Billing20201105CountrySA captures enum value "SA"
	Billing20201105CountrySA Billing20201105Country = "SA"

	// Billing20201105CountrySN captures enum value "SN"
	Billing20201105CountrySN Billing20201105Country = "SN"

	// Billing20201105CountryRS captures enum value "RS"
	Billing20201105CountryRS Billing20201105Country = "RS"

	// Billing20201105CountrySC captures enum value "SC"
	Billing20201105CountrySC Billing20201105Country = "SC"

	// Billing20201105CountrySL captures enum value "SL"
	Billing20201105CountrySL Billing20201105Country = "SL"

	// Billing20201105CountrySG captures enum value "SG"
	Billing20201105CountrySG Billing20201105Country = "SG"

	// Billing20201105CountrySK captures enum value "SK"
	Billing20201105CountrySK Billing20201105Country = "SK"

	// Billing20201105CountryVN captures enum value "VN"
	Billing20201105CountryVN Billing20201105Country = "VN"

	// Billing20201105CountrySI captures enum value "SI"
	Billing20201105CountrySI Billing20201105Country = "SI"

	// Billing20201105CountrySO captures enum value "SO"
	Billing20201105CountrySO Billing20201105Country = "SO"

	// Billing20201105CountryZA captures enum value "ZA"
	Billing20201105CountryZA Billing20201105Country = "ZA"

	// Billing20201105CountryZW captures enum value "ZW"
	Billing20201105CountryZW Billing20201105Country = "ZW"

	// Billing20201105CountryES captures enum value "ES"
	Billing20201105CountryES Billing20201105Country = "ES"

	// Billing20201105CountrySS captures enum value "SS"
	Billing20201105CountrySS Billing20201105Country = "SS"

	// Billing20201105CountrySD captures enum value "SD"
	Billing20201105CountrySD Billing20201105Country = "SD"

	// Billing20201105CountryEH captures enum value "EH"
	Billing20201105CountryEH Billing20201105Country = "EH"

	// Billing20201105CountrySR captures enum value "SR"
	Billing20201105CountrySR Billing20201105Country = "SR"

	// Billing20201105CountrySJ captures enum value "SJ"
	Billing20201105CountrySJ Billing20201105Country = "SJ"

	// Billing20201105CountrySZ captures enum value "SZ"
	Billing20201105CountrySZ Billing20201105Country = "SZ"

	// Billing20201105CountrySE captures enum value "SE"
	Billing20201105CountrySE Billing20201105Country = "SE"

	// Billing20201105CountryCH captures enum value "CH"
	Billing20201105CountryCH Billing20201105Country = "CH"

	// Billing20201105CountrySY captures enum value "SY"
	Billing20201105CountrySY Billing20201105Country = "SY"

	// Billing20201105CountryTJ captures enum value "TJ"
	Billing20201105CountryTJ Billing20201105Country = "TJ"

	// Billing20201105CountryTH captures enum value "TH"
	Billing20201105CountryTH Billing20201105Country = "TH"

	// Billing20201105CountryTG captures enum value "TG"
	Billing20201105CountryTG Billing20201105Country = "TG"

	// Billing20201105CountryTK captures enum value "TK"
	Billing20201105CountryTK Billing20201105Country = "TK"

	// Billing20201105CountryTO captures enum value "TO"
	Billing20201105CountryTO Billing20201105Country = "TO"

	// Billing20201105CountryTT captures enum value "TT"
	Billing20201105CountryTT Billing20201105Country = "TT"

	// Billing20201105CountryAE captures enum value "AE"
	Billing20201105CountryAE Billing20201105Country = "AE"

	// Billing20201105CountryTN captures enum value "TN"
	Billing20201105CountryTN Billing20201105Country = "TN"

	// Billing20201105CountryTR captures enum value "TR"
	Billing20201105CountryTR Billing20201105Country = "TR"

	// Billing20201105CountryTM captures enum value "TM"
	Billing20201105CountryTM Billing20201105Country = "TM"

	// Billing20201105CountryTC captures enum value "TC"
	Billing20201105CountryTC Billing20201105Country = "TC"

	// Billing20201105CountryTV captures enum value "TV"
	Billing20201105CountryTV Billing20201105Country = "TV"

	// Billing20201105CountryUG captures enum value "UG"
	Billing20201105CountryUG Billing20201105Country = "UG"

	// Billing20201105CountryUA captures enum value "UA"
	Billing20201105CountryUA Billing20201105Country = "UA"

	// Billing20201105CountryMK captures enum value "MK"
	Billing20201105CountryMK Billing20201105Country = "MK"

	// Billing20201105CountryEG captures enum value "EG"
	Billing20201105CountryEG Billing20201105Country = "EG"

	// Billing20201105CountryGB captures enum value "GB"
	Billing20201105CountryGB Billing20201105Country = "GB"

	// Billing20201105CountryGG captures enum value "GG"
	Billing20201105CountryGG Billing20201105Country = "GG"

	// Billing20201105CountryJE captures enum value "JE"
	Billing20201105CountryJE Billing20201105Country = "JE"

	// Billing20201105CountryIM captures enum value "IM"
	Billing20201105CountryIM Billing20201105Country = "IM"

	// Billing20201105CountryTZ captures enum value "TZ"
	Billing20201105CountryTZ Billing20201105Country = "TZ"

	// Billing20201105CountryUS captures enum value "US"
	Billing20201105CountryUS Billing20201105Country = "US"

	// Billing20201105CountryVI captures enum value "VI"
	Billing20201105CountryVI Billing20201105Country = "VI"

	// Billing20201105CountryBF captures enum value "BF"
	Billing20201105CountryBF Billing20201105Country = "BF"

	// Billing20201105CountryUY captures enum value "UY"
	Billing20201105CountryUY Billing20201105Country = "UY"

	// Billing20201105CountryUZ captures enum value "UZ"
	Billing20201105CountryUZ Billing20201105Country = "UZ"

	// Billing20201105CountryVE captures enum value "VE"
	Billing20201105CountryVE Billing20201105Country = "VE"

	// Billing20201105CountryWF captures enum value "WF"
	Billing20201105CountryWF Billing20201105Country = "WF"

	// Billing20201105CountryWS captures enum value "WS"
	Billing20201105CountryWS Billing20201105Country = "WS"

	// Billing20201105CountryYE captures enum value "YE"
	Billing20201105CountryYE Billing20201105Country = "YE"

	// Billing20201105CountryZM captures enum value "ZM"
	Billing20201105CountryZM Billing20201105Country = "ZM"
)

// for schema
var billing20201105CountryEnum []interface{}

func init() {
	var res []Billing20201105Country
	if err := json.Unmarshal([]byte(`["UNSET","AF","AL","AQ","DZ","AS","AD","AO","AG","AZ","AR","AU","AT","BS","BH","BD","AM","BB","BE","BM","BT","BO","BA","BW","BV","BR","BZ","IO","SB","VG","BN","BG","MM","BI","BY","KH","CM","CA","CV","KY","CF","LK","TD","CL","CN","TW","CX","CC","CO","KM","YT","CG","CD","CK","CR","HR","CU","CY","CZ","BJ","DK","DM","DO","EC","SV","GQ","ET","ER","EE","FO","FK","GS","FJ","FI","AX","FR","GF","PF","TF","DJ","GA","GE","GM","PS","DE","GH","GI","KI","GR","GL","GD","GP","GU","GT","GN","GY","HT","HM","VA","HN","HK","HU","IS","IN","ID","IR","IQ","IE","IL","IT","CI","JM","JP","KZ","JO","KE","KP","KR","KW","KG","LA","LB","LS","LV","LR","LY","LI","LT","LU","MO","MG","MW","MY","MV","ML","MT","MQ","MR","MU","MX","MC","MN","MD","ME","MS","MA","MZ","OM","NA","NR","NP","NL","CW","AW","SX","BQ","NC","VU","NZ","NI","NE","NG","NU","NF","NO","MP","UM","FM","MH","PW","PK","PA","PG","PY","PE","PH","PN","PL","PT","GW","TL","PR","QA","RE","RO","RU","RW","BL","SH","KN","AI","LC","MF","PM","VC","SM","ST","SA","SN","RS","SC","SL","SG","SK","VN","SI","SO","ZA","ZW","ES","SS","SD","EH","SR","SJ","SZ","SE","CH","SY","TJ","TH","TG","TK","TO","TT","AE","TN","TR","TM","TC","TV","UG","UA","MK","EG","GB","GG","JE","IM","TZ","US","VI","BF","UY","UZ","VE","WF","WS","YE","ZM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billing20201105CountryEnum = append(billing20201105CountryEnum, v)
	}
}

func (m Billing20201105Country) validateBilling20201105CountryEnum(path, location string, value Billing20201105Country) error {
	if err := validate.EnumCase(path, location, value, billing20201105CountryEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this billing 20201105 country
func (m Billing20201105Country) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBilling20201105CountryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this billing 20201105 country based on context it is used
func (m Billing20201105Country) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}