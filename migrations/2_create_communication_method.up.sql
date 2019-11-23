CREATE TABLE communication_method (
	contact_system_code varchar(16) NOT NULL,
	communication_method_code varchar(16) NOT NULL,
	description varchar(32) NOT NULL,
	details varchar(255) NULL,
	status varchar(1) NOT NULL DEFAULT 'A',
	format_field varchar(1024) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	vers integer NULL,
	CONSTRAINT communication_method_pk PRIMARY KEY (communication_method_code),
	CONSTRAINT communication_method_fk FOREIGN KEY (contact_system_code) REFERENCES contact_system(contact_system_code) ON DELETE CASCADE
);
CREATE UNIQUE INDEX communication_method_idx1 ON communication_method (contact_system_code, communication_method_code);

COMMENT ON TABLE communication_method IS 'Communication Method';
COMMENT ON COLUMN communication_method.contact_system_code IS 'Contact System Code';
COMMENT ON COLUMN communication_method.communication_method_code IS 'Communication Method Code';
COMMENT ON COLUMN communication_method.description IS 'Description';
COMMENT ON COLUMN communication_method.details IS 'Details';
COMMENT ON COLUMN communication_method.status IS 'Status';
COMMENT ON COLUMN communication_method.format_field IS 'Format Field';
COMMENT ON COLUMN communication_method.created_at IS 'Created At';
COMMENT ON COLUMN communication_method.modified_at IS 'Updated At';
COMMENT ON COLUMN communication_method.vers IS 'Vers';

INSERT INTO communication_method (contact_system_code, communication_method_code, description, details, status, format_field, vers) VALUES 
('$DEFAULT', 'PHONE', 'Phone', 'Phone', 'A', '[PHONE_NO]', 1),
('$DEFAULT', 'MOBILE', 'Mobile', 'Mobile', 'A', '[MOBILE_NO]', 1),
('$DEFAULT', 'EMAIL', 'Email', 'Email', 'A', '[EMAIL_ADDRESS]', 1);