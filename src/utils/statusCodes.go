package utils

type HttpStatus int

const (
	Ok                  = 200
	Created             = 201
	NoContent           = 204
	NotModified         = 304
	BadRequest          = 400
	NotFound            = 404
	Conflict            = 409
	InternalServerError = 500
)
