package messages

import (
	"fmt"
	"github.com/astlaure/httpmirror/internal/core"
)

func InitTable() {
	core.DB.MustExec(createRequestTableQuery)
	core.DB.MustExec(createMessageTableQuery)
}

func CreateProxyRequest(request ProxyRequest, active ProxyMessage, preview ProxyMessage) {
	tx, err := core.DB.Beginx()

	if err != nil {
		fmt.Println("Cannot start a transaction")
	}

	// insert request
	result, err := tx.NamedExec(insertRequestQuery, &request)

	if err != nil {
		fmt.Println("Cannot insert a request")
		tx.Rollback()
	}

	// get request id
	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Cannot get last id")
		tx.Rollback()
	}

	// insert active
	active.RequestID = uint(id)
	result, err = tx.NamedExec(insertMessageQuery, &active)

	if err != nil {
		fmt.Println("Cannot insert a message")
		tx.Rollback()
	}

	// insert preview
	preview.RequestID = uint(id)
	result, err = tx.NamedExec(insertMessageQuery, &preview)

	if err != nil {
		fmt.Println("Cannot insert a message")
		tx.Rollback()
	}

	// Finish and run
	tx.Commit()
}

func RetrieveRequests() []ProxyRequest {
	var requests []ProxyRequest
	core.DB.Select(&requests, selectAllRequestQuery)
	return requests
}

func RetrieveMessagesByRequestID(requestID uint) []ProxyMessage {
	var proxyMessages []ProxyMessage
	core.DB.Select(&proxyMessages, selectMessagesByRequestIDQuery, requestID)
	return proxyMessages
}
