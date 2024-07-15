package trivec

import "time"

type Tickets []Ticket

type Ticket struct {
	Type          string  `json:"$type"`
	ActualDate    Date    `json:"ActualDate"`
	CenterKey     string  `json:"CenterKey"`
	CenterLeftNr  int     `json:"CenterLeftNr"`
	CenterName    string  `json:"CenterName"`
	CenterNr      string  `json:"CenterNr"`
	CenterRightNr int     `json:"CenterRightNr"`
	Date          Date    `json:"Date,omitempty"`
	Key           string  `json:"Key"`
	AccountKey    string  `json:"AccountKey"`
	AccountName   string  `json:"AccountName"`
	AccountNr     string  `json:"AccountNr"`
	Orders        Orders  `json:"Orders"`
	PcName        string  `json:"PcName"`
	PcNr          int     `json:"PcNr"`
	PrepStatus    string  `json:"PrepStatus"`
	TableName     string  `json:"TableName"`
	TicketNr      int     `json:"TicketNr,omitempty"`
	Time          string  `json:"Time"`
	TotalPrice    float64 `json:"TotalPrice"`
	TotalToPay    float64 `json:"TotalToPay"`
	UserID        string  `json:"UserId"`
	UserKey       string  `json:"UserKey"`
	UserName      string  `json:"UserName"`
	ZNumber       int     `json:"ZNumber,omitempty"`
	Covers        int     `json:"Covers,omitempty"`
	TableNr       int     `json:"TableNr,omitempty"`
}

type Orders []Order

type Order struct {
	Type       string     `json:"$type"`
	ActionID   int        `json:"ActionId"`
	ActualDate Date       `json:"ActualDate"`
	Date       Date       `json:"Date"`
	Key        string     `json:"Key"`
	Lines      OrderLines `json:"Lines"`
	Paymodes   Paymodes   `json:"Paymodes"`
	PcName     string     `json:"PcName"`
	PcNr       int        `json:"PcNr"`
	TicketKey  string     `json:"TicketKey"`
	Time       string     `json:"Time"`
	UserID     string     `json:"UserId"`
	UserKey    string     `json:"UserKey"`
	UserName   string     `json:"UserName"`
	TableNr    int        `json:"TableNr,omitempty"`
}

type OrderLines []OrderLine

type OrderLine struct {
	Type                  string        `json:"$type"`
	Addons                []interface{} `json:"Addons"`
	CourseNr              int           `json:"CourseNr"`
	CourseName            string        `json:"CourseName"`
	GroupKey              string        `json:"GroupKey"`
	GroupLeftNr           int           `json:"GroupLeftNr"`
	GroupName             string        `json:"GroupName"`
	GroupRightNr          int           `json:"GroupRightNr"`
	Key                   string        `json:"Key"`
	Memo                  string        `json:"Memo"`
	MenuID                string        `json:"MenuId"`
	Price                 float64       `json:"Price"`
	ProductKey            string        `json:"ProductKey"`
	ProductName           string        `json:"ProductName"`
	ProductNr             int           `json:"ProductNr"`
	ProductType           string        `json:"ProductType"`
	ProductTypeTranslated string        `json:"ProductTypeTranslated"`
	Qty                   float64       `json:"Qty"`
	TotalDisc             float64       `json:"TotalDisc"`
	TotalEx               float64       `json:"TotalEx"`
	TotalInc              float64       `json:"TotalInc"`
	VatNr                 int           `json:"VatNr"`
	VatPerc               float64       `json:"VatPerc"`
}

type Payments []Payment

type Payment struct {
	Name             string    `json:"name"`
	TotalPrice       float64   `json:"totalPrice"`
	TransactionID    string    `json:"transactionId"`
	TransactionError string    `json:"transactionError"`
	Quantity         float64   `json:"quantity"`
	Units            int       `json:"units"`
	Price            float64   `json:"price"`
	CreateDT         time.Time `json:"createDT"`
	Key              string    `json:"key"`
}

type DxProducts []DxProduct

type DxProduct struct {
	Type               string        `json:"$type"`
	AllowDiscount      bool          `json:"AllowDiscount,omitempty"`
	AskCourse          bool          `json:"AskCourse"`
	AutoAddons         []interface{} `json:"AutoAddons"`
	AutoWindows        []interface{} `json:"AutoWindows"`
	BarCode            string        `json:"BarCode"`
	CenterKeys         []interface{} `json:"CenterKeys"`
	Color              Color         `json:"Color"`
	ColorKey           ColorKey      `json:"ColorKey,omitempty"`
	Description        string        `json:"Description"`
	Fields             []interface{} `json:"Fields"`
	GroupKey           string        `json:"GroupKey"`
	Info               string        `json:"Info"`
	KDSColorKey        KDSColorKey   `json:"KDSColorKey"`
	Key                string        `json:"Key"`
	Name               string        `json:"Name"`
	Parts              []interface{} `json:"Parts"`
	PAXCode            string        `json:"PAXCode"`
	PieceGood          bool          `json:"PieceGood,omitempty"`
	PreparationInfo    string        `json:"PreparationInfo"`
	PrepGroupKey       string        `json:"PrepGroupKey,omitempty"`
	PrepName           string        `json:"PrepName"`
	PrepPrintAddons    bool          `json:"PrepPrintAddons"`
	Price              float64       `json:"Price,omitempty"`
	Prices             Prices        `json:"Prices"`
	PrintMemo          bool          `json:"PrintMemo"`
	PrintX             int           `json:"PrintX"`
	ProdNr             int           `json:"ProdNr"`
	ProductLinks       []interface{} `json:"ProductLinks"`
	ProductType        string        `json:"ProductType"`
	Promo              string        `json:"Promo"`
	PromoCondition     string        `json:"PromoCondition"`
	ShortName          string        `json:"ShortName"`
	ShowInSearch       bool          `json:"ShowInSearch"`
	Tags               []interface{} `json:"Tags"`
	TimeZoneKeys       []interface{} `json:"TimeZoneKeys"`
	Translations       []interface{} `json:"Translations"`
	UnitBase           int           `json:"UnitBase"`
	Vat                float64       `json:"Vat,omitempty"`
	Vat2               float64       `json:"Vat2,omitempty"`
	AskMemo            bool          `json:"AskMemo,omitempty"`
	AskPrice           bool          `json:"AskPrice,omitempty"`
	AllowZeroPrice     bool          `json:"AllowZeroPrice,omitempty"`
	ImageKey           string        `json:"ImageKey,omitempty"`
	PaymodeKey         string        `json:"PaymodeKey,omitempty"`
	VoucherServiceName string        `json:"VoucherServiceName,omitempty"`
	Paymode            DxPaymode     `json:"Paymode,omitempty"`
}

type Prices []Price

type Price struct {
	Type         string  `json:"$type"`
	SeqNr        int     `json:"seqNr"`
	Price        float64 `json:"price"`
	PriceCodeKey string  `json:"PriceCodeKey"`
	Key          string  `json:"key"`
}

type DxPaymodes []DxPaymode

type DxPaymode struct {
	Type             string      `json:"$type"`
	AskPrice         bool        `json:"AskPrice,omitempty"`
	Color            Color       `json:"Color"`
	ColorKey         ColorKey    `json:"ColorKey,omitempty"`
	ConnectionID     string      `json:"ConnectionId"`
	Data             string      `json:"Data"`
	DefaultCash      bool        `json:"DefaultCash"`
	Description      string      `json:"Description"`
	Footer           string      `json:"Footer"`
	GroupKey         string      `json:"GroupKey"`
	GroupLeftNr      int         `json:"GroupLeftNr"`
	GroupRightNr     int         `json:"GroupRightNr"`
	Hotel            string      `json:"Hotel"`
	Info             string      `json:"Info"`
	KDSColorKey      KDSColorKey `json:"KDSColorKey"`
	Key              string      `json:"Key"`
	Name             string      `json:"Name"`
	OpenDrawer       bool        `json:"OpenDrawer,omitempty"`
	PrintMemo        bool        `json:"PrintMemo"`
	PrintX           int         `json:"PrintX"`
	ProdNr           int         `json:"ProdNr,omitempty"`
	ShortName        string      `json:"ShortName"`
	ShowInSearch     bool        `json:"ShowInSearch"`
	TerminalName     string      `json:"TerminalName"`
	UnitBase         int         `json:"UnitBase"`
	NoRefundDiffKey  string      `json:"NoRefundDiffKey,omitempty"`
	PaymodeType      string      `json:"PaymodeType,omitempty"`
	PaymodeTypeID    int         `json:"PaymodeTypeId,omitempty"`
	AskQty           bool        `json:"AskQty,omitempty"`
	IsCash           bool        `json:"IsCash,omitempty"`
	Price            float64     `json:"Price,omitempty"`
	AskTip           bool        `json:"AskTip,omitempty"`
	Drawer           bool        `json:"Drawer,omitempty"`
	FastPayMode      bool        `json:"FastPayMode,omitempty"`
	AccountRequired  bool        `json:"AccountRequired,omitempty"`
	PrintSig         bool        `json:"PrintSig,omitempty"`
	AskMemo          bool        `json:"AskMemo,omitempty"`
	DefaultEpay      bool        `json:"DefaultEpay,omitempty"`
	ParentPaymodeKey string      `json:"ParentPaymodeKey,omitempty"`
	CurrencyKey      string      `json:"CurrencyKey,omitempty"`
	ImageKey         string      `json:"ImageKey,omitempty"`
	PrintAddons      bool        `json:"PrintAddons"`
	InvoiceType      string      `json:"InvoiceType,omitempty"`
	InvoiceTypeID    int         `json:"InvoiceTypeId,omitempty"`
}

type PaymodeGroups []PaymodeGroup

type PaymodeGroup struct {
	Type       string   `json:"$type"`
	Data       string   `json:"Data"`
	GroupLevel int      `json:"GroupLevel"`
	Key        string   `json:"Key"`
	LeftNr     int      `json:"LeftNr"`
	Name       string   `json:"Name"`
	RightNr    int      `json:"RightNr"`
	ColorKey   ColorKey `json:"ColorKey,omitempty"`
}

type ColorKey struct {
	Type  string `json:"$type"`
	PcKey int    `json:"PcKey"`
	RecNr int    `json:"RecNr"`
}

type Color struct {
	Type       string `json:"$type"`
	Background string `json:"Background"`
	Foreground string `json:"Foreground"`
	Key        string `json:"Key"`
}

type KDSColorKey struct {
	Type string `json:"$type"`
}

type Paymodes []Paymode

type Paymode struct {
	Type          string  `json:"$type"`
	GroupKey      string  `json:"GroupKey"`
	GroupLeftNr   int     `json:"GroupLeftNr"`
	GroupName     string  `json:"GroupName"`
	GroupNr       int     `json:"GroupNr"`
	GroupRightNr  int     `json:"GroupRightNr"`
	Key           string  `json:"Key"`
	Memo          string  `json:"Memo"`
	PaymodeKey    string  `json:"PaymodeKey"`
	PaymodeName   string  `json:"PaymodeName"`
	PaymodeNr     int     `json:"PaymodeNr"`
	PaymodeType   string  `json:"PaymodeType"`
	Price         float64 `json:"Price"`
	Qty           float64 `json:"Qty"`
	TerminalID    string  `json:"TerminalId"`
	Tip           float64 `json:"Tip"`
	Total         float64 `json:"Total"`
	TransactionID string  `json:"TransactionId"`
}

type DxProdGroups []DxProdGroup

type DxProdGroup struct {
	Type       string   `json:"$type"`
	ColorKey   ColorKey `json:"ColorKey,omitempty"`
	Data       string   `json:"Data"`
	GroupLevel int      `json:"GroupLevel"`
	Key        string   `json:"Key"`
	LeftNr     int      `json:"LeftNr"`
	Name       string   `json:"Name"`
	RightNr    int      `json:"RightNr"`
	VATCode    VATCode  `json:"VatCode,omitempty"`
	VatCodeKey string   `json:"VatCodeKey,omitempty"`
}

type VATCode struct {
	Type     string `json:"$type"`
	Code     int    `json:"Code"`
	Default  bool   `json:"Default"`
	Default2 bool   `json:"Default2"`
	Key      string `json:"Key"`
	Name     string `json:"Name"`
	Vats     DxVATs `json:"Vats"`
}

type DxVATs []DxVAT

type DxVAT struct {
	Type        string  `json:"$type"`
	CountryCode string  `json:"CountryCode"`
	FromDate    Date    `json:"FromDate"`
	Key         string  `json:"Key"`
	VAT         float64 `json:"VAT"`
}

type AccountInfo struct {
	Type        string  `json:"$type"`
	Address     string  `json:"Address"`
	City        string  `json:"City"`
	ClientNr    string  `json:"ClientNr"`
	CompanyName string  `json:"CompanyName"`
	Country     string  `json:"Country"`
	Data        string  `json:"Data"`
	EMail       string  `json:"EMail"`
	FirstName   string  `json:"FirstName"`
	Gender      string  `json:"Gender"`
	GenderID    int     `json:"GenderId"`
	Info        string  `json:"Info"`
	Key         string  `json:"Key"`
	Language    string  `json:"Language"`
	LastName    string  `json:"LastName"`
	MiddleName  string  `json:"MiddleName"`
	NickName    string  `json:"NickName"`
	Phone       string  `json:"Phone"`
	State       string  `json:"State"`
	VATNumber   string  `json:"VATNumber"`
	ZipCode     string  `json:"ZipCode"`
	BirthDay    float64 `json:"BirthDay"`
	CountryNr   int     `json:"CountryNr"`
}
