package datamodels


type WordDefinition struct{
	ID 				uint64		`json:"id" form:"id"`
	WordId 			uint64		`json:"word_id" form:"word_id"`
	PartOfSpeach	uint		`json:"part_of_speach" form:"part_of_speach"`
	DefData			string		`json:"def_data" form:"def_data"`
	RefFrom 		string		`json:"ref_from" form:"ref_from"`
	Sort 			int32		`json:"sort" form:"sort"`
	Cati 			int32		`json:"cati" form:"cati"`
}	

type DefTranslation struct{
	ID    			uint64		`json:"id" form:"id"`
	DefId 			uint64		`json:"def_id" form:"def_id"`
	TrsData			string		`json:"trs_data" form:"trs_data"`
	Lang 			string		`json:"lang" form:"lang"`
}

type DefSentence struct{
	ID    			uint64		`json:"id" form:"id"`
	DefId 			uint64		`json:"def_id" form:"def_id"`
	StcData 		string		`json:"stc_data" form:"stc_data"`
	Sort 			uint32		`json:"sort" form:"sort"`
}

type StcTranslation struct{
	ID    			uint64		`json:"id" form:"id"`
	DefId 			uint64		`json:"def_id" form:"def_id"`
	TrsData 		string		`json:"trs_data" form:"trs_data"`
	Lang 			string		`json:"lang" form:"lang"`
}