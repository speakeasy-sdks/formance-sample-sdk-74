// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DrinkInput struct {
	// The name of the drink.
	Name string `json:"name"`
	// The price of one unit of the drink in US cents.
	Price float64 `json:"price"`
	// The product code of the drink, only available when authenticated.
	ProductCode *string `json:"productCode,omitempty"`
	// The type of drink.
	Type *DrinkType `json:"type,omitempty"`
}

func (o *DrinkInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DrinkInput) GetPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.Price
}

func (o *DrinkInput) GetProductCode() *string {
	if o == nil {
		return nil
	}
	return o.ProductCode
}

func (o *DrinkInput) GetType() *DrinkType {
	if o == nil {
		return nil
	}
	return o.Type
}
