package core

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type StandardTicValue struct {
	Adsc                               string    // Adresse Secondaire du Compteur
	Vtic                               string    // Version de la TIC
	Date                               time.Time // Date et heure courante
	Ngtf                               string    // Nom du calendrier tarifaire fournisseur
	Ltarf                              string    // Libellé tarif fournisseur en cours
	East                               int32     // Energie active soutirée totale
	Easf01                             int32     // Energie active soutirée Fournisseur, index 01
	Easf02                             int32     // Energie active soutirée Fournisseur, index 02
	Easf03                             int32     // Energie active soutirée Fournisseur, index 03
	Easf04                             int32     // Energie active soutirée Fournisseur, index 04
	Easf05                             int32     // Energie active soutirée Fournisseur, index 05
	Easf06                             int32     // Energie active soutirée Fournisseur, index 06
	Easf07                             int32     // Energie active soutirée Fournisseur, index 07
	Easf08                             int32     // Energie active soutirée Fournisseur, index 08
	Easf09                             int32     // Energie active soutirée Fournisseur, index 09
	Easf10                             int32     // Energie active soutirée Fournisseur, index 10
	Easd01                             int32     // Energie active soutirée Distributeur, index 01
	Easd02                             int32     // Energie active soutirée Distributeur, index 02
	Easd03                             int32     // Energie active soutirée Distributeur, index 03
	Easd04                             int32     // Energie active soutirée Distributeur, index 04
	Eait                               int32     // Energie active injectée totale
	Erq1                               int32     // Energie réactive Q1 totale
	Erq2                               int32     // Energie réactive Q2 totale
	Erq3                               int32     // Energie réactive Q3 totale
	Erq4                               int32     // Energie réactive Q4 totale
	Irms1                              int16     // Courant efficace, phase 1
	Irms2                              int16     // Courant efficace, phase 2
	Irms3                              int16     // Courant efficace, phase 3
	Urms1                              int16     // Tension efficace, phase 1
	Urms2                              int16     // Tension efficace, phase 2
	Urms3                              int16     // Tension efficace, phase 3
	Pref                               int8      // Puissance app. de référence (PREF)
	Pcoup                              int8      // Puissance app. de coupure (PCOUP)
	Sinsts                             int32     // Puissance app. Instantanée soutirée
	Sinsts1                            int32     // Puissance app. Instantanée soutirée phase 1
	Sinsts2                            int32     // Puissance app. instantanée soutirée phase 2
	Sinsts3                            int32     // Puissance app. instantanée soutirée phase 3
	Smaxsn                             int32     // Puissance app. max. soutirée n
	Smaxsn1                            int32     // Puissance app. max. soutirée n phase 1
	Smaxsn2                            int32     // Puissance app. max. soutirée n phase 2
	Smaxsn3                            int32     // Puissance app. max. soutirée n phase 3
	Smaxsnly                           int32     // Puissance app max. soutirée n-1
	Smaxsn1ly                          int32     // Puissance app max. soutirée n-1 phase 1
	Smaxsn2ly                          int32     // Puissance app max. soutirée n-1 phase 2
	Smaxsn3ly                          int32     // Puissance app max. soutirée n-1 phase 3
	Sinsti                             int32     // Puissance app. Instantanée injectée
	Smaxin                             int32     // Puissance app. max. injectée n
	Smaxinly                           int32     // Puissance app max. injectée n-1
	Ccasn                              int32     // Point n de la courbe de charge active soutirée
	Ccasnly                            int32     // Point n-1 de la courbe de charge active soutirée
	Ccain                              int32     // Point n de la courbe de charge active injectée
	Ccainly                            int32     // Point n-1 de la courbe de charge active injectée
	Umoy1                              int16     // Tension moy. ph. 1
	Umoy2                              int16     // Tension moy. ph. 2
	Umoy3                              int16     // Tension moy. ph. 3
	DryContactStatus                   uint8     // Status Contact sec
	CutOffDeviceStatus                 uint8     // Status Organe de coupure
	LinkyTerminalShieldStatus          uint8     // Status État du cache-bornes distributeur
	SurgeStatus                        uint8     // Status Surtension sur une des phases
	ReferencePowerExceededStatus       uint8     // Status Dépassement de la puissance de référence
	ConsumptionStatus                  uint8     // Status Fonctionnement producteur/consommateur
	EnergyDirectionStatus              uint8     // Status Sens de l’énergie active
	ContractTypePriceStatus            uint8     // Status Tarif en cours sur le contrat fourniture
	ContractTypePriceDistributorStatus uint8     // Status Tarif en cours sur le contrat distributeur
	ClockStatus                        uint8     // Status Mode dégradée de l’horloge (perte de l’horodate de l’horloge interne)
	TicStatus                          uint8     // Status État de la sortie télé-information
	EuridisLinkStatus                  uint8     // Status État de la sortie communication Euridis
	CPLStatus                          uint8     // Statut du CPL
	CPLSyncStatus                      uint8     // Status Synchronisation CPL
	TempoContractColorStatus           uint8     // Status Couleur du jour pour le contrat historique tempo
	TempoContractNextDayColorStatus    uint8     // Status Couleur du lendemain pour le contrat historique tempo
	MovingPeakNoticeStatus             uint8     // Status Préavis pointersrs mobiles
	MovingPeakStatus                   uint8     // Status pointers mobile (PM)
	Dpm1                               int8      // Début pointers Mobile 1
	Fpm1                               int8      // Fin pointers Mobile 1
	Dpm2                               int8      // Début pointers Mobile 2
	Fpm2                               int8      // Fin pointers Mobile 2
	Dpm3                               int8      // Début pointers Mobile 3
	Fpm3                               int8      // Fin pointers Mobile 3
	Msg1                               string    // Message court
	Msg2                               string    // Message Ultra court
	Prm                                string    // PRM
	Relai1                             int8      // Relai 1 (Réel)
	Relai2                             int8      // Relai 2
	Relai3                             int8      // Relai 3
	Relai4                             int8      // Relai 4
	Relai5                             int8      // Relai 5
	Relai6                             int8      // Relai 6
	Relai7                             int8      // Relai 7
	Relai8                             int8      // Relai 8
	Ntarf                              int8      // Numéro de l’index tarifaire en cours
	Njourf                             int8      // Numéro du jour en cours calendrier fournisseur
	Njourfnd                           int8      // Numéro du prochain jour calendrier fournisseur
	Pjourfnd                           string    // Profil du prochain jour calendrier fournisseur
	Ppointe                            string    // Profil du prochain jour de pointe
}

// safeUint64ToInt64 converts uint64 to int64 with overflow check
func safeUint64ToInt64(val uint64) int64 {
	if val > math.MaxInt64 {
		return math.MaxInt64
	}
	return safeUint64ToInt64(val)
}

// safeUint64ToInt32 converts uint64 to int32 with overflow check
func safeUint64ToInt32(val uint64) int32 {
	if val > math.MaxInt32 {
		return math.MaxInt32
	}
	return safeUint64ToInt32(val)
}

// safeUint64ToInt16 converts uint64 to int16 with overflow check
func safeUint64ToInt16(val uint64) int16 {
	if val > math.MaxInt16 {
		return math.MaxInt16
	}
	return safeUint64ToInt16(val)
}

// safeUint64ToInt8 converts uint64 to int8 with overflow check
func safeUint64ToInt8(val uint64) int8 {
	if val > math.MaxInt8 {
		return math.MaxInt8
	}
	return safeUint64ToInt8(val)
}

// Parse parameter with name and value
func (tic *StandardTicValue) ParseParam(name string, values []string) {
	if len(values) == 0 {
		return
	}

	switch strings.ToLower(name) {
	case "adsc":
		tic.Adsc = values[0]
	case "vtic":
		tic.Vtic = values[0]
	case "date":
		tic.parseDate(values[1])
	case "ngtf":
		tic.Ngtf = values[0]
	case "ltarf":
		tic.Ltarf = values[0]
	case "east":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.East = safeUint64ToInt32(val)
	case "easf01":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf01 = safeUint64ToInt32(val)
	case "easf02":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf02 = safeUint64ToInt32(val)
	case "easf03":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf03 = safeUint64ToInt32(val)
	case "easf04":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf04 = safeUint64ToInt32(val)
	case "easf05":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf05 = safeUint64ToInt32(val)
	case "easf06":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf06 = safeUint64ToInt32(val)
	case "easf07":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf07 = safeUint64ToInt32(val)
	case "easf08":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf08 = safeUint64ToInt32(val)
	case "easf09":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf09 = safeUint64ToInt32(val)
	case "easf10":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easf10 = safeUint64ToInt32(val)
	case "easd01":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easd01 = safeUint64ToInt32(val)
	case "easd02":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easd02 = safeUint64ToInt32(val)
	case "easd03":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easd03 = safeUint64ToInt32(val)
	case "easd04":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Easd04 = safeUint64ToInt32(val)
	case "eait":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Eait = safeUint64ToInt32(val)
	case "erq1":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Erq1 = safeUint64ToInt32(val)
	case "erq2":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Erq2 = safeUint64ToInt32(val)
	case "erq3":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Erq3 = safeUint64ToInt32(val)
	case "erq4":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Erq4 = safeUint64ToInt32(val)

	case "irms1":
		val, _ := strconv.ParseUint(values[0], 10, 16)
		tic.Irms1 = safeUint64ToInt16(val)

	case "irms2":
		val, _ := strconv.ParseUint(values[0], 10, 16)
		tic.Irms2 = safeUint64ToInt16(val)

	case "irms3":
		val, _ := strconv.ParseUint(values[0], 10, 16)
		tic.Irms3 = safeUint64ToInt16(val)

	case "urms1":
		val, _ := strconv.ParseUint(values[0], 10, 16)
		tic.Urms1 = safeUint64ToInt16(val)

	case "urms2":
		val, _ := strconv.ParseUint(values[0], 10, 16)
		tic.Urms2 = safeUint64ToInt16(val)

	case "urms3":
		val, _ := strconv.ParseUint(values[0], 10, 16)
		tic.Urms3 = safeUint64ToInt16(val)

	case "pref":
		val, _ := strconv.ParseUint(values[0], 10, 8)
		tic.Pref = safeUint64ToInt8(val)

	case "pcoup":
		val, _ := strconv.ParseUint(values[0], 10, 8)
		tic.Pcoup = safeUint64ToInt8(val)

	case "sinsts":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Sinsts = safeUint64ToInt32(val)

	case "sinsts1":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Sinsts1 = safeUint64ToInt32(val)

	case "sinsts2":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Sinsts2 = safeUint64ToInt32(val)

	case "sinsts3":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Sinsts3 = safeUint64ToInt32(val)

	case "smaxsn":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsn = safeUint64ToInt32(val)

	case "smaxsn1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsn1 = safeUint64ToInt32(val)

	case "smaxsn2":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsn2 = safeUint64ToInt32(val)

	case "smaxsn3":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsn3 = safeUint64ToInt32(val)

	case "smaxsn-1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsnly = safeUint64ToInt32(val)

	case "smaxsn1-1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsn1ly = safeUint64ToInt32(val)

	case "smaxsn2-1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsn2ly = safeUint64ToInt32(val)

	case "smaxsn3-1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxsn3ly = safeUint64ToInt32(val)

	case "sinsti":
		val, _ := strconv.ParseUint(values[0], 10, 32)
		tic.Sinsti = safeUint64ToInt32(val)

	case "smaxin":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxin = safeUint64ToInt32(val)

	case "smaxin-1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Smaxinly = safeUint64ToInt32(val)

	case "ccasn":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Ccasn = safeUint64ToInt32(val)

	case "ccasn-1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Ccasnly = safeUint64ToInt32(val)

	case "ccain":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Ccain = safeUint64ToInt32(val)

	case "ccain-1":
		val, _ := strconv.ParseUint(values[1], 10, 32)
		tic.Ccainly = safeUint64ToInt32(val)

	case "umoy1":
		val, _ := strconv.ParseUint(values[1], 10, 16)
		tic.Umoy1 = safeUint64ToInt16(val)

	case "umoy2":
		val, _ := strconv.ParseUint(values[1], 10, 16)
		tic.Umoy2 = safeUint64ToInt16(val)

	case "umoy3":
		val, _ := strconv.ParseUint(values[1], 10, 16)
		tic.Umoy3 = safeUint64ToInt16(val)

	case "status":
		val, _ := strconv.ParseInt(values[0], 10, 64)
		tic.parseStatus(val)

	case "dpm1":
		val, _ := strconv.ParseUint(values[1], 10, 8)
		tic.Dpm1 = safeUint64ToInt8(val)

	case "fpm1":
		val, _ := strconv.ParseUint(values[1], 10, 8)
		tic.Fpm1 = safeUint64ToInt8(val)

	case "dpm2":
		val, _ := strconv.ParseUint(values[1], 10, 8)
		tic.Dpm2 = safeUint64ToInt8(val)

	case "fpm2":
		val, _ := strconv.ParseUint(values[1], 10, 8)
		tic.Fpm2 = safeUint64ToInt8(val)

	case "dpm3":
		val, _ := strconv.ParseUint(values[1], 10, 8)
		tic.Dpm3 = safeUint64ToInt8(val)

	case "fpm3":
		val, _ := strconv.ParseUint(values[1], 10, 8)
		tic.Fpm3 = safeUint64ToInt8(val)

	case "msg1":
		tic.Msg1 = strings.Join(values[:len(values)-1], " ")

	case "msg2":
		tic.Msg2 = strings.Join(values[:len(values)-1], " ")

	case "prm":
		tic.Prm = values[0]

	case "relais":
		val, _ := strconv.ParseUint(values[0], 10, 64)
		tic.parseRelais(safeUint64ToInt64(val))

	case "ntarf":
		val, _ := strconv.ParseUint(values[0], 10, 8)
		tic.Ntarf = safeUint64ToInt8(val)

	case "njourf":
		val, _ := strconv.ParseUint(values[0], 10, 8)
		tic.Njourf = safeUint64ToInt8(val)

	case "njourf+1":
		val, _ := strconv.ParseUint(values[0], 10, 8)
		tic.Njourfnd = safeUint64ToInt8(val)

	case "pjourf+1":
		tic.Pjourfnd = values[0]

	case "ppointe":
		tic.Ppointe = values[0]
	}
}

// Parse date from Tic value
func (values *StandardTicValue) parseDate(value string) {
	season := strings.ToLower(value[0:1])
	if season == "h" {
		value += "+01"
	} else {
		value += "+02"
	}

	val, _ := time.Parse("060102150405-07", value[1:])
	values.Date = val
}

const (
	StatusBitCount = 32
)

// Parse TIC Status information into real status representation
func (values *StandardTicValue) parseStatus(value int64) {
	binaries := addZerosPrefix(strconv.FormatInt(value, 2), StatusBitCount)

	// Bit 0
	values.DryContactStatus = convertStatusToUint(string(binaries[31]))
	// Bit 1 to 3
	values.CutOffDeviceStatus = convertStatusToUint(binaries[28:30])
	// Bit 4
	values.LinkyTerminalShieldStatus = convertStatusToUint(string(binaries[27]))
	// Bit 5 unused
	// Bit 6
	values.SurgeStatus = convertStatusToUint(string(binaries[25]))
	// Bit 7
	values.ReferencePowerExceededStatus = convertStatusToUint(string(binaries[24]))
	// Bit 8
	values.ConsumptionStatus = convertStatusToUint(string(binaries[23]))
	// Bit 9
	values.EnergyDirectionStatus = convertStatusToUint(string(binaries[22]))
	// Bit 10 to 13
	values.ContractTypePriceStatus = convertStatusToUint(binaries[18:21])
	// Bit 14 to 15
	values.ContractTypePriceDistributorStatus = convertStatusToUint(binaries[16:17])
	// Bit 16
	values.ClockStatus = convertStatusToUint(string(binaries[15]))
	// Bit 17
	values.TicStatus = convertStatusToUint(string(binaries[14]))
	// Bit 18 unused
	// Bit 19 to 20
	values.EuridisLinkStatus = convertStatusToUint(binaries[11:12])
	// Bit 21 to 22
	values.CPLStatus = convertStatusToUint(binaries[9:10])
	// Bit 23
	values.CPLSyncStatus = convertStatusToUint(string(binaries[8]))
	// Bit 24 to 25
	values.TempoContractColorStatus = convertStatusToUint(binaries[6:7])
	// Bit 26 to 27
	values.TempoContractNextDayColorStatus = convertStatusToUint(binaries[4:5])
	// Bit 28 to 29
	values.MovingPeakNoticeStatus = convertStatusToUint(binaries[2:3])
	// Bit 30 to 31
	values.MovingPeakStatus = convertStatusToUint(binaries[0:1])
}

const (
	RelayBitCount = 8
)

// Parse TIC Relais information into real representation
func (values *StandardTicValue) parseRelais(value int64) {
	binaries := addZerosPrefix(strconv.FormatInt(value, 2), RelayBitCount)

	// Bit 0
	values.Relai1 = convertRelayValue(binaries[7])
	// Bit 1
	values.Relai2 = convertRelayValue(binaries[6])
	// Bit 2
	values.Relai3 = convertRelayValue(binaries[5])
	// Bit 3
	values.Relai4 = convertRelayValue(binaries[4])
	// Bit 4
	values.Relai5 = convertRelayValue(binaries[3])
	// Bit 5
	values.Relai6 = convertRelayValue(binaries[2])
	// Bit 6
	values.Relai7 = convertRelayValue(binaries[1])
	// Bit 7
	values.Relai8 = convertRelayValue(binaries[0])
}

// Add zeros prefix to represent a full binary value of specified length
func addZerosPrefix(value string, count int) string {
	if len(value) >= count {
		return value
	}
	return strings.Repeat("0", count-len(value)) + value
}

// Convert one status string value to uint8
func convertStatusToUint(status string) uint8 {
	intValue, _ := strconv.ParseUint(status, 2, 8)
	return uint8(intValue)
}

// Convert one relay value to byte
func convertRelayValue(relay byte) int8 {
	if relay == '0' {
		return 0
	}
	return 1
}
