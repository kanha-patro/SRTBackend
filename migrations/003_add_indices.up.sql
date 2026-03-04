CREATE INDEX idx_org_code ON organizations (org_code);
CREATE INDEX idx_route_code ON routes (route_code);
CREATE INDEX idx_trip_driver_id ON trips (driver_id);
CREATE INDEX idx_trip_status ON trips (status);
CREATE INDEX idx_location_trip_id ON locations (trip_id);
CREATE INDEX idx_audit_action ON audit_logs (action);