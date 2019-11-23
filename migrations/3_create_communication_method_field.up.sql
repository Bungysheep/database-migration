CREATE TABLE communication_method_field (
	contact_system_code varchar(16) NOT NULL,
	communication_method_code varchar(16) NOT NULL,
	field_code varchar(16) NOT NULL,
	caption varchar(32) NOT NULL,
	sequence integer NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	modified_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	vers integer NULL,
	CONSTRAINT communication_method_field_pk PRIMARY KEY (field_code),
	CONSTRAINT communication_method_field_contact_system_fk FOREIGN KEY (contact_system_code) REFERENCES contact_system(contact_system_code) ON DELETE CASCADE,
	CONSTRAINT communication_method_field_communication_method_fk FOREIGN KEY (communication_method_code) REFERENCES communication_method(communication_method_code) ON DELETE CASCADE
);
CREATE UNIQUE INDEX communication_method_field_idx1 ON communication_method_field (contact_system_code, communication_method_code, field_code);

COMMENT ON TABLE communication_method_field IS 'Communication Method Field';
COMMENT ON COLUMN communication_method_field.contact_system_code IS 'Contact System Code';
COMMENT ON COLUMN communication_method_field.communication_method_code IS 'Communication Method Code';
COMMENT ON COLUMN communication_method_field.field_code IS 'Field Code';
COMMENT ON COLUMN communication_method_field.caption IS 'Caption';
COMMENT ON COLUMN communication_method_field.sequence IS 'Sequence';
COMMENT ON COLUMN communication_method_field.created_at IS 'Created At';
COMMENT ON COLUMN communication_method_field.modified_at IS 'Updated At';
COMMENT ON COLUMN communication_method_field.vers IS 'Vers';

INSERT INTO communication_method_field (contact_system_code, communication_method_code, field_code, caption, sequence, vers) VALUES 
('$DEFAULT', 'PHONE', 'PHONE_NO', 'Phone', 1, 1),
('$DEFAULT', 'MOBILE', 'MOBILE_NO', 'Mobile', 1, 1),
('$DEFAULT', 'EMAIL', 'EMAIL_ADDRESS', 'Email', 1, 1);