CREATE TABLE "roles" (
  "id" integer PRIMARY KEY,
  "role" varchar NOT NULL
);

CREATE TABLE "roles_permissions" (
  "role_id" integer NOT NULL,
  "permission_id" integer NOT NULL
);

CREATE TABLE "permissions" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "auths" (
  "id" integer PRIMARY KEY,
  "user_id" integer UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "store_id" integer NOT NULL,
  "role_id" integer NOT NULL,
  "email" varchar,
  "phone_number" varchar NOT NULL,
  "name" varchar NOT NULL,
  "is_active" boolean DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stores" (
  "id" integer PRIMARY KEY,
  "store_acces_id" integer NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stores_access" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "menus" (
  "id" integer PRIMARY KEY,
  "store_id" integer NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "items" (
  "id" integer PRIMARY KEY,
  "menu_id" integer NOT NULL,
  "name" varchar NOT NULL,
  "price" integer NOT NULL,
  "category" varchar,
  "is_active" boolean DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "store_tables" (
  "id" integer PRIMARY KEY,
  "store_id" integer NOT NULL,
  "table_number" varchar NOT NULL,
  "qr_code" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "carts" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "store_id" integer NOT NULL,
  "table_id" integer NOT NULL,
  "total_cost" integer,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "carts_items" (
  "id" integer PRIMARY KEY,
  "cart_id" integer NOT NULL,
  "item_id" integer NOT NULL,
  "quantity" integer NOT NULL,
  "price" integer NOT NULL
);

CREATE TABLE "payments" (
  "id" integer PRIMARY KEY,
  "cart_id" integer NOT NULL,
  "method" varchar NOT NULL,
  "midtrans_id" varchar,
  "gross_amount" numeric(12,2) NOT NULL,
  "payment_method" varchar,
  "bank" varchar,
  "payment_receipt" varchar,
  "status" varchar,
  "invoice" varchar,
  "midtrans_data" jsonb,
  "paid_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "payment_id" integer UNIQUE NOT NULL,
  "store_id" integer NOT NULL,
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" integer PRIMARY KEY,
  "store_id" integer NOT NULL,
  "order_id" integer NOT NULL,
  "income" integer NOT NULL,
  "detail" JSON,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "audit_logs" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "role_id" integer NOT NULL,
  "permission_id" integer NOT NULL,
  "store_id" integer NOT NULL,
  "old_data" JSON,
  "new_data" JSON,
  "title" varchar,
  "description" varchar,
  "action" varchar,
  "ip_address" VARCHAR,
  "user_access_from" VARCHAR,
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "roles_permissions" ("role_id", "permission_id");

COMMENT ON COLUMN "roles"."role" IS 'e.g Super Admin , owner, cashier, customer';

COMMENT ON COLUMN "permissions"."name" IS '(e.g. "create_order", "manage_users", "view_reports", "change_table_number")';

COMMENT ON COLUMN "users"."role_id" IS '// 0 / 1 / 2 / 3 /  4';

COMMENT ON COLUMN "users"."name" IS 'nama orang';

COMMENT ON COLUMN "users"."is_active" IS 'account status';

COMMENT ON COLUMN "stores"."store_acces_id" IS '0 no acces  , 1 menu qr  ';

COMMENT ON COLUMN "stores"."name" IS 'nama toko';

COMMENT ON COLUMN "stores_access"."name" IS 'meja_qr , kasir management ';

COMMENT ON COLUMN "items"."category" IS 'snack ,  makanan , minuman';

COMMENT ON COLUMN "items"."is_active" IS 'stock ada = true , stock kosong = false';

COMMENT ON COLUMN "carts"."id" IS 'orderan yang sudah dipesan namun belum di bayar';

COMMENT ON COLUMN "carts"."store_id" IS 'required kalo customer';

COMMENT ON COLUMN "carts"."table_id" IS 'required kalo customer';

COMMENT ON COLUMN "carts"."total_cost" IS 'jumlah quantity x price';

COMMENT ON COLUMN "carts_items"."price" IS 'snapshot harga saat order';

COMMENT ON COLUMN "payments"."method" IS 'bayar dengan qr_midtrans atau di cashier';

COMMENT ON COLUMN "payments"."midtrans_id" IS 'order id (midtrans)';

COMMENT ON COLUMN "payments"."gross_amount" IS 'gross amount (midtrans) ambil dari total cost di carts';

COMMENT ON COLUMN "payments"."payment_method" IS 'payment_type (midtrans)';

COMMENT ON COLUMN "payments"."bank" IS 'jenis bank ';

COMMENT ON COLUMN "payments"."payment_receipt" IS 'untuk struk ';

COMMENT ON COLUMN "payments"."status" IS 'pending, success, failed (status pembayaran)';

COMMENT ON COLUMN "payments"."invoice" IS 'untuk struk';

COMMENT ON COLUMN "orders"."id" IS 'orderan yang sudah di bayar';

COMMENT ON COLUMN "orders"."status" IS 'in_procces (lagi dibuat),canceled (makanan di batalkan), done (sudah di antar ke meja),  (status pesanan)';

COMMENT ON COLUMN "transactions"."id" IS 'catatan untuk database toko berapa pemasukan yang berhasil di proses , kalau status order di cancel maka hapus dari transactions';

COMMENT ON COLUMN "transactions"."detail" IS 'ambil dari order id --> cart id --> cart_items apa aja';

COMMENT ON COLUMN "audit_logs"."old_data" IS '-- data sebelum diubah';

COMMENT ON COLUMN "audit_logs"."new_data" IS '-- data sesudah diubah';

COMMENT ON COLUMN "audit_logs"."description" IS 'bahasa manusia';

COMMENT ON COLUMN "audit_logs"."action" IS 'LOGIN, LOGOUT, CREATE, UPDATE, DELETE, PAYMENT, EXPORT.';

COMMENT ON COLUMN "audit_logs"."ip_address" IS '-- IPv4/IPv6';

COMMENT ON COLUMN "audit_logs"."user_access_from" IS '-- device/browser';

COMMENT ON COLUMN "audit_logs"."status" IS 'SUCCESS, FAILED.';

ALTER TABLE "roles_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "roles_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "auths" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "stores" ADD FOREIGN KEY ("store_acces_id") REFERENCES "stores_access" ("id");

ALTER TABLE "menus" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");

ALTER TABLE "store_tables" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "carts" ADD FOREIGN KEY ("table_id") REFERENCES "store_tables" ("id");

ALTER TABLE "carts_items" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "carts_items" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
