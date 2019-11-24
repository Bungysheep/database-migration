CREATE TABLE contact_communication_method_field (
	contact_system_code varchar(16) NOT NULL,
	contact_id integer NOT NULL,
	contact_communication_method_id integer NOT NULL,
	field_code varchar(16) NOT NULL,
	field_value varchar(32) NULL,
	CONSTRAINT contact_communication_method_field_contact_system_fk FOREIGN KEY (contact_system_code) REFERENCES contact_system(contact_system_code) ON DELETE CASCADE,
	CONSTRAINT contact_communication_method_field_contact_fk FOREIGN KEY (contact_id) REFERENCES contact(contact_id) ON DELETE CASCADE,
	CONSTRAINT contact_communication_method_field_contact_communication_method_fk FOREIGN KEY (contact_communication_method_id) REFERENCES contact_communication_method(contact_communication_method_id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX contact_communication_method_field_idx1 ON contact_communication_method_field (contact_system_code, contact_id, contact_communication_method_id, field_code);

COMMENT ON TABLE contact_communication_method_field IS 'Contact Communication Method Field';
COMMENT ON COLUMN contact_communication_method_field.contact_system_code IS 'Contact System Code';
COMMENT ON COLUMN contact_communication_method_field.contact_id IS 'Contact ID';
COMMENT ON COLUMN contact_communication_method_field.contact_communication_method_id IS 'Contact Communication Method ID';
COMMENT ON COLUMN contact_communication_method_field.field_code IS 'Field Code';
COMMENT ON COLUMN contact_communication_method_field.field_value IS 'Field Value';