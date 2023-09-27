package gowoocommerce

type Order struct {
	ID                 int           `json:"id"`
	ParentID           int           `json:"parent_id"`
	Number             string        `json:"number"`
	OrderKey           string        `json:"order_key"`
	CreatedVia         string        `json:"created_via"`
	Version            string        `json:"version"`
	Status             string        `json:"status"`
	Currency           string        `json:"currency"`
	DateCreated        string        `json:"date_created"`
	DateCreatedGmt     string        `json:"date_created_gmt"`
	DateModified       string        `json:"date_modified"`
	DateModifiedGmt    string        `json:"date_modified_gmt"`
	DiscountTotal      string        `json:"discount_total"`
	DiscountTax        string        `json:"discount_tax"`
	ShippingTotal      string        `json:"shipping_total"`
	ShippingTax        string        `json:"shipping_tax"`
	CartTax            string        `json:"cart_tax"`
	Total              string        `json:"total"`
	TotalTax           string        `json:"total_tax"`
	PricesIncludeTax   bool          `json:"prices_include_tax"`
	CustomerID         int           `json:"customer_id"`
	CustomerIPAddress  string        `json:"customer_ip_address"`
	CustomerUserAgent  string        `json:"customer_user_agent"`
	CustomerNote       string        `json:"customer_note"`
	Billing            Billing       `json:"billing"`
	Shipping           Shipping      `json:"shipping"`
	PaymentMethod      string        `json:"payment_method"`
	PaymentMethodTitle string        `json:"payment_method_title"`
	TransactionID      string        `json:"transaction_id"`
	DatePaid           string        `json:"date_paid"`
	DatePaidGmt        string        `json:"date_paid_gmt"`
	DateCompleted      string        `json:"date_completed"`
	DateCompletedGmt   string        `json:"date_completed_gmt"`
	CartHash           string        `json:"cart_hash"`
	MetaData           []MetaData    `json:"meta_data"`
	LineItems          []LineItems   `json:"line_items"`
	TaxLines           []interface{} `json:"tax_lines"`
	ShippingLines      []interface{} `json:"shipping_lines"`
	FeeLines           []interface{} `json:"fee_lines"`
	CouponLines        []interface{} `json:"coupon_lines"`
	Refunds            []interface{} `json:"refunds"`
	Links              Links         `json:"_links"`
}
type Billing struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
type Shipping struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
}
type MetaData struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
type LineItems struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	ProductID   int           `json:"product_id"`
	VariationID int           `json:"variation_id"`
	Quantity    int           `json:"quantity"`
	TaxClass    string        `json:"tax_class"`
	Subtotal    string        `json:"subtotal"`
	SubtotalTax string        `json:"subtotal_tax"`
	Total       string        `json:"total"`
	TotalTax    string        `json:"total_tax"`
	Taxes       []interface{} `json:"taxes"`
	MetaData    []interface{} `json:"meta_data"`
	Sku         string        `json:"sku"`
	Price       int           `json:"price"`
}
type Self struct {
	Href string `json:"href"`
}
type Collection struct {
	Href string `json:"href"`
}
type Customer struct {
	Href string `json:"href"`
}
type Links struct {
	Self       []Self       `json:"self"`
	Collection []Collection `json:"collection"`
	Customer   []Customer   `json:"customer"`
}
