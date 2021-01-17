package partners


type Partner struct {
	Id 									int64
	Storage_partner_name 				string
	Storage_partner_company_name 		string
	Storage_partner_company_gst 		string
	Provides_goods_transport_service 	bool
	Provides_goods_packaging_service 	bool
	Provides_goods_insurance_service	bool
	Listing_active 						bool
	Email_id 							string
	Phone_numbers 						string

	Verified							bool
	Password							string

	Date_created 						string
}

type PartnerLoginRequest struct {
	Email_id 							string
	Password							string
}