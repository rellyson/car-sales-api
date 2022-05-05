CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "sellers"(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name varchar(100) NOT NULL,
    email varchar(35) NOT NULL,
    password varchar NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);

CREATE TYPE "car_type" AS ENUM ('HATCH', 'SEDAN', 'SUV', 'TRUCK');

CREATE TABLE "cars"(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    brand varchar(35) NOT NULL,
    model varchar(50) NOT NULL,
    type car_type NOT NULL,
    manufacture_year integer NOT NULL,
    model_year integer NOT NULL,
    accessories text[], 
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);

CREATE TABLE "ads"(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    value float(2) NOT NULL,
    is_available boolean DEFAULT false,
    sold_in timestamp NOT NULL,
    car_id uuid,
    seller_id uuid,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    CONSTRAINT fk_car
      FOREIGN KEY(car_id) 
	    REFERENCES cars(id),
    CONSTRAINT fk_seller
      FOREIGN KEY(seller_id) 
	    REFERENCES sellers(id)
);
