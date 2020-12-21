package gluasocket_socketheaders

import (
	"github.com/yuin/gopher-lua"
)

// ----------------------------------------------------------------------------

func Loader(l *lua.LState) int {
	if err := l.DoString(headersDotLua); err != nil {
		l.RaiseError("Error loading headers.lua: %v", err)
		return 0
	}
	return 1
}

const headersDotLua = `-----------------------------------------------------------------------------
-- Canonic header field capitalization
-- LuaSocket toolkit.
-- Author: Diego Nehab
-----------------------------------------------------------------------------
local socket = require("socket")
socket.headers = {}
local _M = socket.headers

_M.canonic = {
    ["accept"] = "Accept",
    ["accept-charset"] = "Accept-Charset",
    ["accept-encoding"] = "Accept-Encoding",
    ["accept-language"] = "Accept-Language",
    ["accept-ranges"] = "Accept-Ranges",
    ["action"] = "Action",
    ["alternate-recipient"] = "Alternate-Recipient",
    ["age"] = "Age",
    ["allow"] = "Allow",
    ["arrival-date"] = "Arrival-Date",
    ["authorization"] = "Authorization",
    ["bcc"] = "Bcc",
    ["cache-control"] = "Cache-Control",
    ["cc"] = "Cc",
    ["comments"] = "Comments",
    ["connection"] = "Connection",
    ["content-description"] = "Content-Description",
    ["content-disposition"] = "Content-Disposition",
    ["content-encoding"] = "Content-Encoding",
    ["content-id"] = "Content-ID",
    ["content-language"] = "Content-Language",
    ["content-length"] = "Content-Length",
    ["content-location"] = "Content-Location",
    ["content-md5"] = "Content-MD5",
    ["content-range"] = "Content-Range",
    ["content-transfer-encoding"] = "Content-Transfer-Encoding",
    ["content-type"] = "Content-Type",
    ["cookie"] = "Cookie",
    ["date"] = "Date",
    ["diagnostic-code"] = "Diagnostic-Code",
    ["dsn-gateway"] = "DSN-Gateway",
    ["etag"] = "ETag",
    ["expect"] = "Expect",
    ["expires"] = "Expires",
    ["final-log-id"] = "Final-Log-ID",
    ["final-recipient"] = "Final-Recipient",
    ["from"] = "From",
    ["host"] = "Host",
    ["if-match"] = "If-Match",
    ["if-modified-since"] = "If-Modified-Since",
    ["if-none-match"] = "If-None-Match",
    ["if-range"] = "If-Range",
    ["if-unmodified-since"] = "If-Unmodified-Since",
    ["in-reply-to"] = "In-Reply-To",
    ["keywords"] = "Keywords",
    ["last-attempt-date"] = "Last-Attempt-Date",
    ["last-modified"] = "Last-Modified",
    ["location"] = "Location",
    ["max-forwards"] = "Max-Forwards",
    ["message-id"] = "Message-ID",
    ["mime-version"] = "MIME-Version",
    ["original-envelope-id"] = "Original-Envelope-ID",
    ["original-recipient"] = "Original-Recipient",
    ["pragma"] = "Pragma",
    ["proxy-authenticate"] = "Proxy-Authenticate",
    ["proxy-authorization"] = "Proxy-Authorization",
    ["range"] = "Range",
    ["received"] = "Received",
    ["received-from-mta"] = "Received-From-MTA",
    ["references"] = "References",
    ["referer"] = "Referer",
    ["remote-mta"] = "Remote-MTA",
    ["reply-to"] = "Reply-To",
    ["reporting-mta"] = "Reporting-MTA",
    ["resent-bcc"] = "Resent-Bcc",
    ["resent-cc"] = "Resent-Cc",
    ["resent-date"] = "Resent-Date",
    ["resent-from"] = "Resent-From",
    ["resent-message-id"] = "Resent-Message-ID",
    ["resent-reply-to"] = "Resent-Reply-To",
    ["resent-sender"] = "Resent-Sender",
    ["resent-to"] = "Resent-To",
    ["retry-after"] = "Retry-After",
    ["return-path"] = "Return-Path",
    ["sender"] = "Sender",
    ["server"] = "Server",
    ["smtp-remote-recipient"] = "SMTP-Remote-Recipient",
    ["status"] = "Status",
    ["subject"] = "Subject",
    ["te"] = "TE",
    ["to"] = "To",
    ["trailer"] = "Trailer",
    ["transfer-encoding"] = "Transfer-Encoding",
    ["upgrade"] = "Upgrade",
    ["user-agent"] = "User-Agent",
    ["vary"] = "Vary",
    ["via"] = "Via",
    ["warning"] = "Warning",
    ["will-retry-until"] = "Will-Retry-Until",
    ["www-authenticate"] = "WWW-Authenticate",
    ["x-mailer"] = "X-Mailer",
}

return _M
`
