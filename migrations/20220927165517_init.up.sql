CREATE TABLE "delivery" (
  "id" uuid PRIMARY KEY,
  "template_id" uuid NOT NULL,
  "send_at" timestamp NOT NULL,
  "sended" bool NOT NULL,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL
);

CREATE TABLE "template" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL,
  "path" varchar(300) NOT NULL
);

CREATE TABLE "subscribers" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "surname" varchar(100) NOT NULL,
  "email" varchar(100) NOT NULL,
  "age" int NOT NULL,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL
);

CREATE TABLE "subscriber_in_delivery" (
  "id" uuid PRIMARY KEY,
  "delivery_id" uuid NOT NULL,
  "subscriber_id" uuid NOT NULL,
  "opened" bool NOT NULL
);

ALTER TABLE "delivery" ADD FOREIGN KEY ("template_id") REFERENCES "template" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "subscriber_in_delivery" ADD FOREIGN KEY ("delivery_id") REFERENCES "delivery" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "subscriber_in_delivery" ADD FOREIGN KEY ("subscriber_id") REFERENCES "subscribers" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;