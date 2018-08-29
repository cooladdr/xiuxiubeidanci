package datamodels

type BookClass struct{
	ID			uint64	`json:"id" form:"id"`
	Name		string	`json:"name" form:"name"`
}	

type Book struct{
	ID			uint64	`json:"id" form:"id"`
	Name		string	`json:"name" form:"name"`
	ClassId		uint64	`json:"class_id" form:"class_id"`
}

type BookWords struct{
	ID 			uint64	`json:"id" form:"id"`
	BookId		uint64	`json:"book_id" form:"book_id"`
	WordId		uint64	`json:"word_id" form:"word_id"`
}
