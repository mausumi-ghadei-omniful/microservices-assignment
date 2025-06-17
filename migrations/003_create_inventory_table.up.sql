-- Create inventory table
CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL PRIMARY KEY,
    hub_id INTEGER NOT NULL,
    sku_id INTEGER NOT NULL,
    quantity INTEGER DEFAULT 0,
    reserved INTEGER DEFAULT 0,
    available INTEGER DEFAULT 0,
    min_threshold INTEGER DEFAULT 0,
    max_threshold INTEGER,
    tenant_id VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    -- Foreign key constraints
    CONSTRAINT fk_inventory_hub FOREIGN KEY (hub_id) REFERENCES hubs(id) ON DELETE CASCADE,
    CONSTRAINT fk_inventory_sku FOREIGN KEY (sku_id) REFERENCES skus(id) ON DELETE CASCADE,
    
    -- Unique constraint for hub-sku combination
    CONSTRAINT uk_inventory_hub_sku UNIQUE (hub_id, sku_id)
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_inventory_tenant_id ON inventory(tenant_id);
CREATE INDEX IF NOT EXISTS idx_inventory_hub_id ON inventory(hub_id);
CREATE INDEX IF NOT EXISTS idx_inventory_sku_id ON inventory(sku_id);
CREATE INDEX IF NOT EXISTS idx_inventory_deleted_at ON inventory(deleted_at);

-- Create updated_at trigger
CREATE TRIGGER update_inventory_updated_at 
    BEFORE UPDATE ON inventory 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Create trigger to calculate available quantity
CREATE OR REPLACE FUNCTION calculate_available_quantity()
RETURNS TRIGGER AS $$
BEGIN
    NEW.available = NEW.quantity - NEW.reserved;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_inventory_available 
    BEFORE INSERT OR UPDATE ON inventory 
    FOR EACH ROW 
    EXECUTE FUNCTION calculate_available_quantity(); 