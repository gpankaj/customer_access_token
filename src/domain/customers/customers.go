package customers


type Customer struct {
	Customer_id 						int64
	Customer_name		 				string
	Customer_phone_number		 		string
	Customer_address					string
	Customer_city						string
	Customer_state						string
	Customer_phone_verified			 	bool
	Customer_phone_verification_code	string
	Customer_comments					string
	Customer_email_id 					string
	Customer_active 					bool
	Customer_date_created 				string
	Customer_password					string
	Customer_verified					bool
}

type CustomerLoginRequest struct {
	Customer_email_id 							string
	Customer_password							string
}
