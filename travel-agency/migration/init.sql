CREATE TABLE "user"
(
  id SERIAL primary key,
  name VARCHAR not null,
  phonenumber VARCHAR not null,
  email VARCHAR not null,
  password VARCHAR not null,
);
CREATE TABLE "tour"
(
  id SERIAL primary key,
  start_date DATE not null,
  end_date DATE not null,
  amount_of_people VARCHAR not null,
  user_id INT REFERENCES "user" (id) ON DELETE CASCADE NOT NULL,
);
CREATE TABLE "reservation"
(
  id SERIAL primary key,
  reservation_date DATE not null,
  payment FLOAT not null,
  tour_id INT REFERENCES "tour" (id) ON DELETE CASCADE NOT NULL,
);
CREATE TABLE "tour_base"
(
  id SERIAL primary key,
  city_of_flight VARCHAR not null,
  arrival_country VARCHAR not null,
  duration INT not null,
  date_of_tour DATE not null,
  tour_cost FLOAT not null,
  tour_id INT REFERENCES "tour" (id) ON DELETE CASCADE NOT NULL,
);
CREATE TABLE "review"
(
  id SERIAL primary key,
  score INT not null,
  review_text VARCHAR not null,
  publish_date DATE not null,
  user_id INT REFERENCES "user" (id) ON DELETE CASCADE NOT NULL,
  tour_id INT REFERENCES "tour" (id) ON DELETE CASCADE NOT NULL,
);