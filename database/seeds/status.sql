INSERT INTO status (name, slug) VALUES
('Active', 'active'),
('Disabled', 'disabled'),
('Pending', 'pending'),
('Ex-member', 'ex')
ON CONFLICT (slug) DO NOTHING;
