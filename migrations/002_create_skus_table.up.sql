-- Create skus table
CREATE TABLE IF NOT EXISTS skus (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(100),
    brand VARCHAR(100),
    weight DECIMAL(10,2),
    dimensions VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    tenant_id VARCHAR(100) NOT NULL,
    seller_id VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_skus_tenant_id ON skus(tenant_id);
CREATE INDEX IF NOT EXISTS idx_skus_seller_id ON skus(seller_id);
CREATE INDEX IF NOT EXISTS idx_skus_code ON skus(code);
CREATE INDEX IF NOT EXISTS idx_skus_deleted_at ON skus(deleted_at);

-- Create updated_at trigger
CREATE TRIGGER update_skus_updated_at 
    BEFORE UPDATE ON skus 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column(); 