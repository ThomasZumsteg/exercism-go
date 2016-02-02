package ledger

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

//TestVersion is the version of the unit tests that this will pass
const TestVersion = 2

//Entry a line in a ledger
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

//byDateDescriptionChange how to sort entries
type byDateDescriptionChange []Entry

func (d byDateDescriptionChange) Len() int      { return len(d) }
func (d byDateDescriptionChange) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d byDateDescriptionChange) Less(i, j int) bool {
	if d[i].Date != d[j].Date {
		return d[i].Date < d[j].Date
	} else if d[i].Description != d[j].Description {
		return d[i].Description < d[j].Description
	} else {
		return d[i].Change < d[j].Change
	}
}

/*FormatLedger prints a ledger of entries given some localization.*/
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	var currencySymbol string
	switch currency {
	case "USD":
		currencySymbol = "$"
	case "EUR":
		currencySymbol = "€"
	default:
		return "", fmt.Errorf("Unknown currency: %s", currency)
	}

	dateFmt := "2006-02-01"
	headerFmt := "%-10s | %-25s | %s\n"
	var localDateFmt, ledger string
	var currencyFormatter func(int, string) string
	switch locale {
	case "nl-NL":
		localDateFmt = "01-02-2006"
		ledger = fmt.Sprintf(headerFmt, "Datum", "Omschrijving", "Verandering")
		currencyFormatter = formatDutchCurrency
	case "en-US":
		localDateFmt = "02/01/2006"
		ledger = fmt.Sprintf(headerFmt, "Date", "Description", "Change")
		currencyFormatter = formatUSCurrency
	default:
		return "", fmt.Errorf("Unknown locale: %s", locale)
	}

	sort.Sort(byDateDescriptionChange(entriesCopy))

	for _, entry := range entriesCopy {
		if len(entry.Description) > 25 {
			entry.Description = entry.Description[:22] + "..."
		}

		date, err := time.Parse(dateFmt, entry.Date)
		if err != nil {
			return "", err
		}
		entry.Date = date.Format(localDateFmt)

		ledger += fmt.Sprintf("%10s | %-25s | %13s\n",
			entry.Date, entry.Description,
			currencyFormatter(entry.Change, currencySymbol))
	}
	return ledger, nil
}

/*formatDutchCurrency formats currency the Dutch way.*/
func formatDutchCurrency(cents int, currencySymbol string) string {
	if cents < 0 {
		return fmt.Sprintf("%s %s-", currencySymbol, humanReadable(-cents, ".", ","))
	}
	return fmt.Sprintf("%s %s ", currencySymbol, humanReadable(cents, ".", ","))
}

/*formatUSCurrency formats currency the US way.*/
func formatUSCurrency(cents int, currencySymbol string) string {
	if cents < 0 {
		return fmt.Sprintf("(%s%s)", currencySymbol, humanReadable(-cents, ",", "."))
	}
	return fmt.Sprintf("%s%s ", currencySymbol, humanReadable(cents, ",", "."))
}

/*humanReadable makes a number with powers seperated.*/
func humanReadable(cents int, seperator, decimal string) string {
	dollars := strconv.Itoa(cents / 100)
	for i := len(dollars) - 3; 0 < i; i -= 3 {
		dollars = dollars[:i] + seperator + dollars[i:]
	}
	return fmt.Sprintf("%s%s%02d", dollars, decimal, cents%100)
}
