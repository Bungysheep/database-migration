CREATE TABLE contact_system (
	contact_system_code varchar(16) NOT NULL,
	description varchar(32) NOT NULL,
	details varchar(255) NULL,
	status varchar(1) NOT NULL DEFAULT 'A',
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	vers integer NULL,
	CONSTRAINT contact_system_pk PRIMARY KEY (contact_system_code)
);
CREATE UNIQUE INDEX contact_system_idx1 ON contact_system (contact_system_code);

COMMENT ON TABLE contact_system IS 'Contact System';
COMMENT ON COLUMN contact_system.contact_system_code IS 'Contact System Code';
COMMENT ON COLUMN contact_system.description IS 'Description';
COMMENT ON COLUMN contact_system.details IS 'Details';
COMMENT ON COLUMN contact_system.status IS 'Status';
COMMENT ON COLUMN contact_system.created_at IS 'Created At';
COMMENT ON COLUMN contact_system.modified_at IS 'Updated At';
COMMENT ON COLUMN contact_system.vers IS 'Vers';

INSERT INTO contact_system (contact_system_code, description, details, status, vers) VALUES 
('$DEFAULT', 'Default Contact System', 'Default Contact System', 'A', 1);