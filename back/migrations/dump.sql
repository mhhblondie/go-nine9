CREATE TABLE "users"(
    "id" UUID NOT NULL,
    "lastname" VARCHAR(255) NOT NULL,
    "firstname" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "roles" VARCHAR(255) NOT NULL,
    "salon_id" UUID NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);


CREATE TABLE "reservations"(
    "id" UUID NOT NULL,
    "slot_id" UUID NOT NULL,
    "customer_id" UUID NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);

CREATE TABLE "salons"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "description" TEXT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);

CREATE TABLE "services"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NULL,
    "salon_id" UUID NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);

CREATE TABLE "prestations"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "price" DOUBLE PRECISION NOT NULL,
    "service_id" UUID NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);


CREATE TABLE "slots"(
    "id" UUID NOT NULL,
    "slot_time" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "hairdressing_staff_id" UUID NULL,
    "salon_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);

CREATE TABLE "hours"(
    "id" UUID NOT NULL,
    "salon_id" UUID NOT NULL,
    "day_of_week" VARCHAR(255) NOT NULL,
    "opening_time" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "closing_time" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);


ALTER TABLE
    "reservations" ADD PRIMARY KEY("id");
ALTER TABLE
    "salons" ADD PRIMARY KEY("id");
ALTER TABLE
    "services" ADD PRIMARY KEY("id");
ALTER TABLE
    "prestations" ADD PRIMARY KEY("id");
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
ALTER TABLE
    "hours" ADD PRIMARY KEY("id");
ALTER TABLE
    "slots" ADD PRIMARY KEY("id");

ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
ALTER TABLE
    "hours" ADD CONSTRAINT "hours_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id") ON DELETE CASCADE;
ALTER TABLE
    "reservations" ADD CONSTRAINT "reservations_slot_id_foreign" FOREIGN KEY("slot_id") REFERENCES "slots"("id") ON DELETE CASCADE;
ALTER TABLE
    "slots" ADD CONSTRAINT "slots_hairdressing_staff_id_foreign" FOREIGN KEY("hairdressing_staff_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "slots" ADD CONSTRAINT "slots_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id") ON DELETE CASCADE;
ALTER TABLE
    "reservations" ADD CONSTRAINT "reservations_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "services" ADD CONSTRAINT "services_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id") ON DELETE CASCADE;
ALTER TABLE
    "users" ADD CONSTRAINT "users_salon_id_foreign" FOREIGN KEY("salon_id") REFERENCES "salons"("id") ON DELETE CASCADE;
ALTER TABLE
    "prestations" ADD CONSTRAINT "prestations_service_id_foreign" FOREIGN KEY("service_id") REFERENCES "services"("id") ON DELETE CASCADE;


INSERT INTO "salons" ("id", "name", "address", "phone","description")
    VALUES ('ce801349-752c-4cee-a660-9f7cceaf7132', 'Hairstyle', '18 rue Paris,75016', '010230','Plongez dans un monde de luxe et de détente à "Harmonie des Sens", où la coiffure et le bien-être se rencontrent pour créer une expérience inoubliable. Notre salon de coiffure et spa combine l''art de la coiffure avec des soins spa indulgents pour vous offrir une escapade revitalisante loin du stress quotidien. Laissez-vous dorloter par nos coiffeurs experts qui maîtrisent les dernières tendances en matière de coupe, de couleur et de coiffage pour vous aider à révéler votre beauté naturelle. Après votre séance de coiffure, offrez-vous une expérience relaxante dans notre spa où vous pourrez profiter de massages apaisants, de soins du visage rajeunissants et de traitements capillaires luxueux. Avec une attention méticuleuse aux détails et un service personnalisé, notre équipe dévouée à "Harmonie des Sens" s''engage à vous offrir une expérience sensorielle complète qui éveillera vos sens et restaurera votre équilibre intérieur.');

INSERT INTO "salons" ("id", "name", "address", "phone","description")
    VALUES ('ce801349-752c-4cee-a660-9f7cceaf7139', 'Hairstyle2', '18 rue Paris,75016', '010232','Bienvenue à "Élégance Urbaine", votre refuge de beauté et de style au cœur de la ville. Notre salon de coiffure allie l''expertise professionnelle à une atmosphère chaleureuse et accueillante pour vous offrir une expérience de coiffure exceptionnelle. Nos coiffeurs talentueux sont passionnés par leur métier et sont dévoués à vous aider à réaliser la coiffure parfaite qui correspond à votre style et à votre personnalité. Que vous souhaitiez une coupe moderne et tendance, une coloration vibrante ou un traitement capillaire revitalisant, notre équipe est là pour vous offrir des services de haute qualité qui vous laisseront sentir confiant et magnifique. Détendez-vous et laissez-nous prendre soin de vous dans notre oasis de beauté où le bien-être de nos clients est notre priorité absolue.');

INSERT INTO "users" ("id", "lastname", "firstname", "email", "password", "roles", "salon_id")
    VALUES ('9f7ceeaf-7132-4cee-a660-ce801349-852c', 'Lastname', 'Firstname', 'email@example.com', 'password', 'manager', 'ce801349-752c-4cee-a660-9f7cceaf7132');

INSERT INTO "users" ("id", "lastname", "firstname", "email", "password", "roles", "salon_id")
    VALUES ('9f7ceeaf-7132-4cee-a660-ce801349-853c', 'admin', 'admin', 'admin', 'admin', 'admin', 'ce801349-752c-4cee-a660-9f7cceaf7139');

INSERT INTO "slots" ("id", "slot_time", "hairdressing_staff_id", "salon_id","created_at","updated_at","deleted_at") VALUES
    ('9f7ceeaf-7132-4cee-a660-ce801349854c',	'2024-02-12 10:00:00',	'9f7ceeaf-7132-4cee-a660-ce801349852c',	'ce801349-752c-4cee-a660-9f7cceaf7132',NULL,NULL,NULL),
    ('9f7ceeaf-7132-4cee-a660-ce801349858c',	'2024-02-12 11:00:00',	'9f7ceeaf-7132-4cee-a660-ce801349852c',	'ce801349-752c-4cee-a660-9f7cceaf7132',NULL,NULL,NULL),
    ('9f7ceeaf-7132-4cee-a660-ce801349857c',	'2024-02-12 12:00:00',	'9f7ceeaf-7132-4cee-a660-ce801349852c',	'ce801349-752c-4cee-a660-9f7cceaf7132',NULL,NULL,NULL),
    ('9f7ceeaf-7132-4cee-a660-ce801349856c',	'2024-02-13 10:00:00',	'9f7ceeaf-7132-4cee-a660-ce801349852c',	'ce801349-752c-4cee-a660-9f7cceaf7132',NULL,NULL,NULL),
    ('9f7ceeaf-7132-4cee-a660-ce801349855c',	'2024-02-13 12:00:00',	'9f7ceeaf-7132-4cee-a660-ce801349852c',	'ce801349-752c-4cee-a660-9f7cceaf7132',NULL,NULL,NULL);

INSERT INTO "hours" ("id","salon_id", "day_of_week", "opening_time", "closing_time")
VALUES
    ('ce801349-752c-4cee-a660-9f7cceaf7125','ce801349-752c-4cee-a660-9f7cceaf7132', 'Lundi','2024-01-21 09:00:00', '2024-01-21 18:00:00'),
    ('ce801349-752c-4cee-a660-9f7cceaf7126','ce801349-752c-4cee-a660-9f7cceaf7132', 'Mardi','2024-01-21 09:00:00', '2024-01-21 18:00:00'),
    ('ce801349-752c-4cee-a660-9f7cceaf7123','ce801349-752c-4cee-a660-9f7cceaf7132', 'Mercredi','2024-01-21 09:00:00', '2024-01-21 18:00:00'),
    ('ce801349-752c-4cee-a660-9f7cceaf7124','ce801349-752c-4cee-a660-9f7cceaf7132', 'Jeudi','2024-01-21 09:00:00', '2024-01-21 20:00:00'),
    ('ce801349-752c-4cee-a660-9f7cceaf7122','ce801349-752c-4cee-a660-9f7cceaf7132', 'Vendredi','2024-01-21 09:00:00', '2024-01-21 20:00:00'),
    ('ce801349-752c-4cee-a660-9f7cceaf7121','ce801349-752c-4cee-a660-9f7cceaf7132', 'Samedi','2024-01-21 09:00:00', '2024-01-21 17:00:00'),
    ('ce801349-752c-4cee-a660-9f7cceaf7120','ce801349-752c-4cee-a660-9f7cceaf7132', 'Dimanche', NULL, NULL);


INSERT INTO "services" ("id", "name", "description", "salon_id")
    VALUES ('ce801349-752c-4cee-a660-9f7cceaf7133', 'Coupe - Homme ', 'Chaque prestation comprend un diagnostic où nous prenons le temps d''échanger sur vos envies et attentes, suivi d''un shampoing avec le traditionnel massage crânien, ensuite nous passons à la coupe et procédons pour finir au coiffage.', 'ce801349-752c-4cee-a660-9f7cceaf7132');

INSERT INTO "prestations" ("id", "name", "price", "service_id")
VALUES 
    ('ce801349-752c-4cee-a660-9f7cceaf7134', 'Styliste (Etudiant ou -20ans) - Shampoing + Coupe personnalisée + Coiffage', 40, 'ce801349-752c-4cee-a660-9f7cceaf7133'),
    ('ce801349-752c-4cee-a660-9f7cceaf7135', 'Styliste (+ 20 ans) - Shampoing + Coupe personnalisée + Coiffage', 45, 'ce801349-752c-4cee-a660-9f7cceaf7133'),
    ('ce801349-752c-4cee-a660-9f7cceaf7136', 'Master styliste (Etudiant ou -20ans) - Shampoing + Coupe personnalisée + Coiffage',50, 'ce801349-752c-4cee-a660-9f7cceaf7133'),
    ('ce801349-752c-4cee-a660-9f7cceaf7138', 'Master styliste (+20 ans) - Shampoing + Coupe personnalisée + Coiffage 1h', 65, 'ce801349-752c-4cee-a660-9f7cceaf7133');