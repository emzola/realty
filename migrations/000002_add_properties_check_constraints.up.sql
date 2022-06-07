ALTER TABLE properties ADD CONSTRAINT properties_price_check CHECK (price >= 0);
ALTER TABLE properties ADD CONSTRAINT type_length_check CHECK (array_length(type, 1) BETWEEN 1 AND 1);
ALTER TABLE properties ADD CONSTRAINT category_length_check CHECK (array_length(category, 1) BETWEEN 1 AND 1);
ALTER TABLE properties ADD CONSTRAINT currency_length_check CHECK (array_length(currency, 1) BETWEEN 1 AND 1);