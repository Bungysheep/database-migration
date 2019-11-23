CREATE TABLE communication_method_label (
	contact_system_code varchar(16) NOT NULL,
	communication_method_code varchar(16) NOT NULL,
	communication_method_label_code varchar(8) NOT NULL,
	caption varchar(16) NOT NULL,
	CONSTRAINT communication_method_label_contact_system_fk FOREIGN KEY (contact_system_code) REFERENCES contact_system(contact_system_code) ON DELETE CASCADE,
	CONSTRAINT communication_method_label_communication_method_fk FOREIGN KEY (communication_method_code) REFERENCES communication_method(communication_method_code) ON DELETE CASCADE
);
CREATE UNIQUE INDEX communication_method_label_idx1 ON communication_method_label (contact_system_code, communication_method_code, communication_method_label_code);

COMMENT ON TABLE communication_method_label IS 'Communication Method Label';
COMMENT ON COLUMN communication_method_label.contact_system_code IS 'Contact System Code';
COMMENT ON COLUMN communication_method_label.communication_method_code IS 'Communication Method Code';
COMMENT ON COLUMN communication_method_label.communication_method_label_code IS 'Label Code';
COMMENT ON COLUMN communication_method_label.caption IS 'Caption';

INSERT INTO communication_method_label (contact_system_code, communication_method_code, communication_method_label_code, caption) VALUES 
('$DEFAULT', 'PHONE', 'HOME', 'Home'),
('$DEFAULT', 'PHONE', 'WORK', 'Work'),
('$DEFAULT', 'MOBILE', 'PERSONAL', 'Personal'),
('$DEFAULT', 'MOBILE', 'WORK', 'Work'),
('$DEFAULT', 'EMAIL', 'PERSONAL', 'Personal'),
('$DEFAULT', 'EMAIL', 'WORK', 'Work');