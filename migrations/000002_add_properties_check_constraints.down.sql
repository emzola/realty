ALTER TABLE properties DROP CONSTRAINT IF EXISTS properties_price_check;
ALTER TABLE properties DROP CONSTRAINT IF EXISTS type_length_check;
ALTER TABLE properties DROP CONSTRAINT IF EXISTS category_length_check;
ALTER TABLE properties DROP CONSTRAINT IF EXISTS currency_length_check;
ALTER TABLE properties DROP CONSTRAINT IF EXISTS features_length_check;
ALTER TABLE properties DROP CONSTRAINT IF EXISTS nearby_length_check;