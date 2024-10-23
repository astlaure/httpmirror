package messages

var createRequestTableQuery = `
CREATE TABLE IF NOT EXISTS proxy_requests (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    service VARCHAR(255) NOT NULL,
    tracking VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

var createMessageTableQuery = `
CREATE TABLE IF NOT EXISTS proxy_messages (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    request_id INTEGER NOT NULL,
    status INTEGER NOT NULL,
    path VARCHAR(255) NOT NULL,
    protocol VARCHAR(255) NOT NULL,
    headers TEXT,
    body TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (request_id) REFERENCES proxy_requests(id)
);
`

var selectAllRequestQuery = `
SELECT * FROM proxy_requests LIMIT 12;
`

var insertRequestQuery = `
INSERT INTO proxy_requests (service, tracking)
VALUES (:service, :tracking);
`

var insertMessageQuery = `
INSERT INTO proxy_messages (request_id, status, path, protocol, headers, body)
VALUES (:request_id, :status, :path, :protocol, :headers, :body);
`

var selectMessagesByRequestIDQuery = `
SELECT * FROM proxy_messages WHERE request_id = ?;
`
