CREATE TABLE recipes (
  id int PRIMARY KEY AUTO_INCREMENT,
  -- name of recipe
  title varchar(100)  NOT NULL,
  -- time required to cook/bake the recipe
  making_time varchar(100) NOT NULL,
  -- number of people the recipe will feed
  serves varchar(100) NOT NULL,
  -- food items necessary to prepare the recipe
  ingredients varchar(300) NOT NULL,
  -- price of recipe
  cost integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `recipes` (
  id,
  title,
  making_time,
  serves,
  ingredients,
  cost,
  created_at,
  updated_at
)
VALUES (
  1,
  'チキンカレー',
  '45分',
  '4人',
  '玉ねぎ,肉,スパイス',
  1000,
  '2016-01-10 12:10:12',
  '2016-01-10 12:10:12'
);

INSERT INTO `recipes` (
  id,
  title,
  making_time,
  serves,
  ingredients,
  cost,
  created_at,
  updated_at
)
VALUES (
  2,
  'オムライス',
  '30分',
  '2人',
  '玉ねぎ,卵,スパイス,醤油',
  700,
  '2016-01-11 13:10:12',
  '2016-01-11 13:10:12'
);
