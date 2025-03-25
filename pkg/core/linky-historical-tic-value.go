package core

import (
	"strconv"
	"strings"
)

// Internal linky values object to each metrics
type HistoricalTicValue struct {
	Adco     string // Adresse du compteur
	Optarif  string // Option tarifaire choisie
	Isousc   uint8  // Intensité souscrite en A
	Base     int32  // Index option Base
	Hchc     int32  // Index option Heures creuses : Heures Creuses en Wh
	Hchp     int32  // Index option Heures pleines : Heures Pleines en Wh
	Ejphn    int32  // Index option EJP : Heures Normales en Wh
	Ejphpn   int32  // Index option EJP : Heures de Pointe Mobile en Wh
	Bbrhcjb  int32  // Index option Tempo : Heures Creuses Jours Bleus en Wh
	Bbrhpjb  int32  // Index option Tempo : Heures Pleines Jours Bleus en Wh
	Bbrhcjw  int32  // Index option Tempo : Heures Creuses Jours Blancs en Wh
	Bbrhpjw  int32  // Index option Tempo : Heures Pleines Jours Blancs en Wh
	Bbrhcjr  int32  // Index option Tempo : Heures Creuses Jours Rouges en Wh
	Bbrhpjr  int32  // Index option Tempo : Heures Pleines Jours Rouges en Wh
	Pejp     int8   // Préavis Début EJP (30 min) en minutes
	Ptec     string // Période Tarifaire en cours
	Demain   string // Couleur du lendemain
	Iinst    int16  // Intensité instantanée en A : Courant efficace (en A)
	Iinst1   int16  // Intensité Instantanée phase 1 en A
	Iinst2   int16  // Intensité Instantanée phase 2 en A
	Iinst3   int16  // Intensité Instantanée phase 3 en A
	Adps     int16  // Avertissement de Dépassement De Puissance Souscrite en A : Courant efficace, si Ilnst > IR
	Imax     int16  // Intensité maximale appelée en A
	Imax1    int16  // Intensité maximale appelée phase 1 en A
	Imax2    int16  // Intensité maximale appelée phase 2 en A
	Imax3    int16  // Intensité maximale appelée phase 3 en A
	Pmax     int32  // Puissance maximale triphasée atteinte en W
	Papp     int32  // Puissance Apparente en VA
	Hhphc    string // Horaire Heures Pleines Heures Creuses
	Motdetat string // Mot d'état du compteur
	Ppot     string // potentiels is here
}

// Parse parameter with name and value
func (tic *HistoricalTicValue) ParseParam(name string, values []string) {
	if len(values) == 0 {
		return
	}

	switch strings.ToLower(name) {
	case "adco":
		tic.Adco = values[0]
	case "optarif":
		tic.Optarif = values[0]
	case "isousc":
		val, _ := strconv.ParseUint(values[0], 10, 8)
		tic.Isousc = uint8(val)
	case "base":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Base = int32(val)
	case "hchc":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Hchc = int32(val)
	case "hchp":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Hchp = int32(val)
	case "ejphn":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Ejphn = int32(val)
	case "ejphpn":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Ejphpn = int32(val)
	case "bbrhcjb":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Bbrhcjb = int32(val)
	case "bbrhpjb":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Bbrhpjb = int32(val)
	case "bbrhcjw":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Bbrhcjw = int32(val)
	case "bbrhpjw":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Bbrhpjw = int32(val)
	case "bbrhcjr":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Bbrhcjr = int32(val)
	case "bbrhpjr":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Bbrhpjr = int32(val)
	case "pejp":
		val, _ := strconv.ParseInt(values[0], 10, 8)
		tic.Pejp = int8(val)
	case "ptec":
		tic.Ptec = values[0]
	case "demain":
		tic.Demain = values[0]
	case "iinst":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Iinst = int16(val)
	case "iinst1":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Iinst1 = int16(val)
	case "iinst2":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Iinst2 = int16(val)
	case "iinst3":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Iinst3 = int16(val)
	case "adps":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Adps = int16(val)
	case "imax":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Imax = int16(val)
	case "imax1":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Imax1 = int16(val)
	case "imax2":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Imax2 = int16(val)
	case "imax3":
		val, _ := strconv.ParseInt(values[0], 10, 16)
		tic.Imax3 = int16(val)
	case "pmax":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Pmax = int32(val)
	case "papp":
		val, _ := strconv.ParseInt(values[0], 10, 32)
		tic.Papp = int32(val)
	case "hhphc":
		tic.Hhphc = values[0]
	case "motdetat":
		tic.Motdetat = strings.Join(values[:len(values)-1], " ")
	case "ppot":
		tic.Ppot = values[0]
	}
}
