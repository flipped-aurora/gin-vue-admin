package credential

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

// TestGetTicketFromServer .
func TestGetTicketFromServer(t *testing.T) {
	defer gock.Off()
	gock.New(getTicketURL).Reply(200).JSON(&ResTicket{Ticket: "mock-ticket", ExpiresIn: 10})
	ticket, err := GetTicketFromServer("arg-ak")
	assert.Nil(t, err)
	assert.Equal(t, int64(0), ticket.ErrCode)
	assert.Equal(t, "mock-ticket", ticket.Ticket, "they should be equal")
	assert.Equal(t, int64(10), ticket.ExpiresIn, "they should be equal")
}
