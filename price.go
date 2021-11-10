package goprice

import (
	"encoding/json"
	"fmt"
	"github.com/leekchan/accounting"
	"math"
	"strconv"
	"strings"
)

const (
	DefaultRoundOnFloat     = 0.6
	DefaultRoundPlacesFloat = 2
)

type Price float64

func NewPrice(value float64) *Price {
	price := Price(value)
	return &price
}

func (a Price) MarshalJSON() ([]byte, error) {
	str := a.RoundPriceStringFormatAtCheckout(DefaultRoundOnFloat, DefaultRoundPlacesFloat)
	return json.Marshal(str)
}

func (a *Price) UnmarshalJSON(data []byte) error {
	var stringValue string
	var floatValue float64
	err := json.Unmarshal(data, &stringValue)
	if err != nil {
		return err
	}

	// remove ',' from string value
	stringValue = strings.Replace(stringValue, ",", "", -1)
	floatValue, err = strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return err
	}

	*a = Price(floatValue)
	return nil
}

func (a *Price) Add(value *Price)  {
	newPrice := *a + *value
	*a = newPrice
}

func (a *Price) Multiply(value *Price) {
	newPrice := *a * *value
	*a = newPrice
}

func (a *Price) Minus(value *Price) {
	newPrice := *a - *value
	*a = newPrice
}

func (a *Price) Copy() *Price {
	newPrice := *a
	return &newPrice
}

func (a Price) ToStringWithCurrencySymbol(currencySymbol string) string {
	str := a.RoundPriceStringFormatAtCheckout(DefaultRoundOnFloat, 0)
	return fmt.Sprintf("%s%s", currencySymbol, str)
}

func (a Price) ToString() string {
	return a.RoundPriceStringFormatAtCheckout(DefaultRoundOnFloat, DefaultRoundPlacesFloat)
}

func (a Price) ToFloat64() float64 {
	return float64(a)
}

func (a Price) ToAdyenString(adyenInt *int) string {
	if adyenInt == nil {
		temp := a.ToAdyenInt()
		adyenInt = &temp
	}
	return fmt.Sprintf("%d", adyenInt)
}

func (a Price) ToAdyenInt() int {
	amountMulti100 := float64(a) * 100
	amount := a.round(amountMulti100, DefaultRoundOnFloat, DefaultRoundPlacesFloat)
	adyenAmount := int(amount)
	return adyenAmount
}

func (a Price) round(val, roundOn float64, precision int) float64 {
	var round float64
	pow := math.Pow(10, float64(precision))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

func (a Price) RoundPriceStringFormatAtCheckout(roundOn float64, precision int) string {
	return a.priceStringFormatAtCheckout(a.round(a.ToFloat64(), roundOn, precision), precision)
}

func (a Price) priceStringFormatAtCheckout(val float64, precision int) string {
	ac := accounting.Accounting{Precision: precision}
	price := ac.FormatMoney(val)
	return a.beautyMoney(price)
}

// beautyMoney format 2.00 to 2
func (a Price) beautyMoney(price string) string {
	return strings.Replace(price, ".00", "", 1)
}

