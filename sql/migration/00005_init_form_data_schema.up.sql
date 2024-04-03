BEGIN;

-- Form data stores the actual form submissions in JSON format
CREATE TABLE "form_data" (
    id UUID PRIMARY KEY,
    form_id TEXT NOT NULL,
    status VARCHAR(255) NOT NULL,
    data JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    FOREIGN KEY (form_id) REFERENCES form (id)
);

END;