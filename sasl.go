// Library for Simple Authentication and Security Layer (SASL).
package sasl

import (
	"github.com/emersion/imap/backend"
)

// Note:
//   Most of this code was copied, with some modifications, from net/smtp. It
//   would be better if Go provided a standard package (e.g. crypto/sasl) that
//   could be shared by SMTP, IMAP, and other packages.

// Client interface to perform challenge-response authentication.
type Client interface {
	// Begins SASL authentication with the server. It returns the
	// authentication mechanism name and "initial response" data (if required by
	// the selected mechanism). A non-nil error causes the client to abort the
	// authentication attempt.
	//
	// A nil ir value is different from a zero-length value. The nil value
	// indicates that the selected mechanism does not use an initial response,
	// while a zero-length value indicates an empty initial response, which must
	// be sent to the server.
	Start() (mech string, ir []byte, err error)

	// Continues challenge-response authentication. A non-nil error causes
	// the client to abort the authentication attempt.
	Next(challenge []byte) (response []byte, err error)
}

// Server interface to perform challenge-response authentication.
type Server interface {
	// Begins SASL authentication with the client.
	Start() (ir []byte, err error)

	// Continues challenge-response authentication.
	Next(challenge []byte) (response []byte, err error)

	// Get the authenticated user. Returns nil if no user has been authenticated.
	User() backend.User
}

// A function that creates SASL servers.
type ServerFactory func() Server
