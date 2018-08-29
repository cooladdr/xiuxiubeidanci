package datamodels

type Word struct{
	ID			uint64		`json:"id" form:"id"`
	Spelling	string		`json:"spelling" form:"spelling"`
	InUSA		string		`json:"in_usa" form:"in_usa"`
	InUK		string		`json:"in_uk" form:"in_uk"`
	WType		uint		`json:"w_type" form:"w_type"`
}

type WordRelationship struct{
	ID 				uint64		`json:"id" form:"id"`
	Major 			string		`json:"major" form:"major"`
	Minor			string		`json:"minor" form:"minor"`
	Relationship 	uint		`json:"relationship" form:"relationship"`
}	

type WordShowName struct{
	WordId 			uint64		`json:"word_id" form:"word_id"`
	Spelling 		string		`json:"spelling" form:"spelling"`
}	


